package ctrl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/liujunren93/share_rbac/intenal/entity"
	"github.com/liujunren93/share_rbac/log"
	"github.com/liujunren93/share_rbac/pb"
	"github.com/liujunren93/share_utils/common/casbin/adapter/simple"
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
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act
`

func init() {
	casBinMode, _ = model.NewModelFromString(modeStr)
	localStorage, _ = lru.NewLRU[[][]string](0, 86400)
}

func (*rbacCtrl) userPolicy(uid, domainId int64, roleIds []int64) ([][]string, error) {
	var prolicy = make([][]string, 0, len(roleIds))
	for _, v := range roleIds {
		prolicy = append(prolicy, []string{"g", strconv.Itoa(int(uid)), strconv.Itoa(int(v)), strconv.Itoa(int(domainId))})
	}
	return prolicy, nil
}
func (r *rbacCtrl) domainPolicy(ctx context.Context, domainId int64) ([][]string, error) {
	var prolicy [][]string
	if value, ok := localStorage.Get(domainId); !ok {

		req := pb.GetDomainPolicyReq{DomainID: domainId}
		dr, err := r.grpcClient.GetDomainPolicy(ctx, &req)
		if err != nil {
			log.Logger.Error(err)
			return nil, err
		}
		var data []entity.DomainPolicy
		err = json.Unmarshal([]byte(dr.Data), &data)
		if err != nil {
			log.Logger.Error(err)
			return nil, err
		}
		prolicy = make([][]string, 0, len(data))
		for _, v := range data {
			// //p, admin, domain1, data1, read
			prolicy = append(prolicy, []string{"p", strconv.Itoa(int(v.RoleID)), strconv.Itoa(int(domainId)), v.ApiPath, v.Method})
		}
		// localStorage.Add(domainId, prolicy)
	} else {
		prolicy = value
	}
	return prolicy, nil
}

func (ctrl *rbacCtrl) CheckPermission(ctx context.Context, reqPath, method string, roleIds []int64, uid, domainId int64) error {
	domainPolicy, err := ctrl.domainPolicy(ctx, domainId)
	if err != nil {
		log.Logger.Error("CheckPermission.domainPolicy", err)
		return err
	}
	userPolicy, err := ctrl.userPolicy(uid, domainId, roleIds)
	if err != nil {
		log.Logger.Error("CheckPermission.userPolicy", err)
		return err
	}
	domainPolicy = append(domainPolicy, userPolicy...)
	b, _ := json.Marshal(domainPolicy)
	fmt.Println(string(b))
	f, _ := os.Create("./1.csy")
	tmp := ""
	for _, datas := range domainPolicy {

		for _, v := range datas {
			tmp += v + ","
		}
		tmp = tmp[:len(tmp)-1]
		tmp += "\n"
	}

	io.WriteString(f, tmp)
	f.Close()
	e, err := casbin.NewEnforcer(casBinMode, simple.NewAdapter(domainPolicy))
	if err != nil {
		log.Logger.Error("CheckPermission.casbin.NewEnforcer", err)
		return err
	}
	// ok, err := e.Enforce("1", "1", reqPath, method)
	ok, err := e.Enforce("1", "1", "/rbac/path", "GET")
	if err != nil {
		log.Logger.Error("CheckPermission.casbin.Enforce", err)
		return err
	}
	if ok {
		return nil
	}

	return errors.New("Forbidden")
}
