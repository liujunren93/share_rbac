package ctrl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/liujunren93/share_rbac/intenal/entity"
	"github.com/liujunren93/share_rbac/log"
	"github.com/liujunren93/share_rbac/pb"
	"github.com/liujunren93/share_utils/common/storage/lru"
)

var localStorage *lru.LRU[[][]string]
var casBinMode model.Model

const modeStr = `
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom &&  keyMatch2(r.obj, p.obj)  && regexMatch(r.act, p.act)
`

func init() {
	casBinMode, _ = model.NewModelFromString(modeStr)
	localStorage, _ = lru.NewLRU[[][]string](0, 86400)
}

func (ctrl *rbacCtrl) initCasPolicy() error {
	var err error
	if ctrl.syncedEnforcer == nil {
		ctrl.casOnce.Do(func() {
			ctrl.syncedEnforcer, err = casbin.NewSyncedEnforcer(casBinMode)
			if err != nil {
				log.Logger.Error("initCasPolicy.NewSyncedEnforcer", err)

			}

		})
	}
	return err

}

func (ctrl *rbacCtrl) domainPolicy(ctx context.Context, domainId int64) error {
	key := fmt.Sprintf("p_%d", domainId)

	if _, ok := ctrl.prolicyMap.Load(key); !ok {
		var prolicy [][]string
		if err := ctrl.initCasPolicy(); err != nil {
			return err
		}
		req := pb.GetDomainPolicyReq{DomainID: domainId}
		dr, err := ctrl.grpcClient.GetDomainPolicy(ctx, &req)
		if err != nil {
			log.Logger.Error("domainPolicy.GetDomainPolicy", err)
			return err
		}
		var data []entity.DomainPolicy
		err = json.Unmarshal([]byte(dr.Data), &data)
		if err != nil {
			log.Logger.Error("domainPolicy.Unmarshal", err)
			return err
		}
		prolicy = make([][]string, 0, len(data))
		for _, v := range data {
			// //p, admin, domain1, data1, read
			prolicy = append(prolicy, []string{strconv.Itoa(int(v.RoleID)), strconv.Itoa(int(domainId)), v.ApiPath, v.Method})
		}
		_, err = ctrl.syncedEnforcer.AddPolicies(prolicy)
		if err != nil {
			log.Logger.Error("domainPolicy.AddPolicy", err)
			return err
		}
		ctrl.prolicyMap.Store(key, struct{}{})
	}

	return nil

}

func (ctrl *rbacCtrl) userPolicy(uid, domainId int64, roleIds []int64) error {

	key := fmt.Sprintf("g_%d_%d", domainId, uid)
	if _, ok := ctrl.prolicyMap.Load(key); !ok {
		var prolicy = make([][]string, 0, len(roleIds))
		for _, v := range roleIds {
			prolicy = append(prolicy, []string{strconv.Itoa(int(uid)), strconv.Itoa(int(v)), strconv.Itoa(int(domainId))})
		}
		_, err := ctrl.syncedEnforcer.AddGroupingPolicies(prolicy)
		if err != nil {
			log.Logger.Error("userPolicy.AddPolicy", err)
			return err
		}
		ctrl.prolicyMap.Store(key, struct{}{})

	}
	return nil

}

func (ctrl *rbacCtrl) CheckPermission(ctx context.Context, reqPath, method string, roleIds []int64, uid, domainId int64) error {
	err := ctrl.domainPolicy(ctx, domainId)
	if err != nil {
		log.Logger.Debug("CheckPermission.domainPolicy", err)
		return err
	}
	err = ctrl.userPolicy(uid, domainId, roleIds)
	if err != nil {
		log.Logger.Debug("CheckPermission.userPolicy", err)
		return err
	}

	// ok, err := e.Enforce(uid, domain, source, method)
	ok, err := ctrl.syncedEnforcer.Enforce(strconv.Itoa(int(uid)), strconv.Itoa(int(domainId)), reqPath, method)
	if err != nil {
		log.Logger.Error("CheckPermission.casbin.Enforce", err)
		return err
	}
	if ok {
		return nil
	}
	return errors.New("Forbidden")
}
