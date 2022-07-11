package ctrl

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	pb "github.com/liujunren93/share_rbac/rbac_pb"
	"github.com/liujunren93/share_utils/common/auth"
	"github.com/liujunren93/share_utils/common/mq"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/netHelper"
	"google.golang.org/grpc"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 11:44
 */
type rbacCtrl struct {
	Auther         auth.Auther
	mq             mq.Mqer
	grpcClient     pb.RbacClient
	syncedEnforcer *casbin.SyncedEnforcer
	casOnce        *sync.Once
	prolicyMap     sync.Map
	ctx            context.Context
}

var (
	DOMIAN_ID = pb.SESSION_DOMAIN_ID.String() //int64
	UID       = pb.SESSION_UID.String()
	ROLES     = pb.SESSION_ROLE_IDS.String() //[]int64
)

var (
	RbacCtrl     *rbacCtrl
	rbacCtrlOnce = &sync.Once{}
)

func InitRbacCtrl(ctx context.Context, Auther auth.Auther, mq mq.Mqer, grpcClient grpc.ClientConnInterface) {
	rbacCtrlOnce.Do(func() {
		RbacCtrl = &rbacCtrl{
			Auther:     Auther,
			mq:         mq,
			grpcClient: pb.NewRbacClient(grpcClient),
			casOnce:    &sync.Once{},
			prolicyMap: sync.Map{},
			ctx:        ctx,
		}
	})

}

func (ctrl *rbacCtrl) DomainList(ctx *gin.Context) {
	var req pb.DomainListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := ctrl.grpcClient.MDomainList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}
func (ctrl *rbacCtrl) DomainInfo(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))

	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}

	info, err := ctrl.grpcClient.MDomainInfo(ctx, &req)

	netHelper.Response(ctx, info, err, nil)
}
func (ctrl *rbacCtrl) DomainDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))

	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}

	info, err := ctrl.grpcClient.MDomainDel(ctx, &req)
	fmt.Println(info, err)
	netHelper.Response(ctx, info, err, nil)

}
func (ctrl *rbacCtrl) DomainCreate(ctx *gin.Context) {
	var req pb.DomainCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	create, err := ctrl.grpcClient.MDomainCreate(ctx, &req)
	netHelper.Response(ctx, create, err, nil)
}
func (ctrl *rbacCtrl) DomainUpdate(ctx *gin.Context) {
	var req pb.DomainUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	dr, err := ctrl.grpcClient.MDomainUpdate(ctx, &req)
	netHelper.Response(ctx, dr, err, nil)

}

func (ctrl *rbacCtrl) AdminList(ctx *gin.Context) {
	var req pb.AdminListReq
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}

	res, err := ctrl.grpcClient.MAdminList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)

}

func (ctrl *rbacCtrl) AdminCreate(ctx *gin.Context) {
	var req pb.AdminCreateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	create, err := ctrl.grpcClient.MAdminCreate(ctx, &req)
	netHelper.Response(ctx, create, err, nil)
}

func (ctrl *rbacCtrl) AdminUpdate(ctx *gin.Context) {
	var req pb.AdminUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}

	update, err := ctrl.grpcClient.MAdminUpdate(ctx, &req)
	netHelper.Response(ctx, update, err, nil)

}

func (ctrl *rbacCtrl) AdminDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	update, err := ctrl.grpcClient.MAdminDel(ctx, &req)
	netHelper.Response(ctx, update, err, nil)
}

func (ctrl *rbacCtrl) AdminInfo(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}

	update, err := ctrl.grpcClient.MAdminInfo(ctx, &req)
	netHelper.Response(ctx, update, err, nil)
}

func (ctrl *rbacCtrl) AdminRoleList(ctx *gin.Context) {
	var req pb.AdminRoleListReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	req.UID = int64(atoi)

	update, err := ctrl.grpcClient.MAdminRoleList(ctx, &req)
	netHelper.Response(ctx, update, err, nil)
}

func (ctrl *rbacCtrl) AdminRoleSet(ctx *gin.Context) {
	var req pb.AdminRoleSetReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.DomainID = ctx.GetInt64(DOMIAN_ID)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.UID = int64(atoi)
	update, err := ctrl.grpcClient.MAdminRoleSet(ctx, &req)
	netHelper.Response(ctx, update, err, nil)
}

func (ctrl *rbacCtrl) PermissionList(ctx *gin.Context) {
	var req pb.PermissionListReq
	// req.DomainID = ctx.GetInt64(DOMIAN_ID)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	list, err := ctrl.grpcClient.MPermissionList(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (ctrl *rbacCtrl) PermissionInfo(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}

	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	list, err := ctrl.grpcClient.MPermissionInfo(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (ctrl *rbacCtrl) PermissionCreate(ctx *gin.Context) {
	var req pb.PermissionCreateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	list, err := ctrl.grpcClient.MPermissionCreate(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (ctrl *rbacCtrl) PermissionUpdate(ctx *gin.Context) {
	var req pb.PermissionUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	list, err := ctrl.grpcClient.MPermissionUpdate(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (ctrl *rbacCtrl) PermissionDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	list, err := ctrl.grpcClient.MPermissionDel(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (ctrl *rbacCtrl) PermissionPathList(ctx *gin.Context) {
	var req pb.PermissionPathListReq

	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	req.PermissionID = int64(id)
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	res, err := ctrl.grpcClient.MPermissionPathList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (ctrl *rbacCtrl) PermissionPathSet(ctx *gin.Context) {
	var req pb.PermissionPathSetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	req.PermissionID = int64(id)
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	res, err := ctrl.grpcClient.MPermissionPathSet(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (ctrl *rbacCtrl) RoleList(ctx *gin.Context) {
	var req pb.RoleListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	res, err := ctrl.grpcClient.MRoleList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (ctrl *rbacCtrl) RoleCreate(ctx *gin.Context) {
	var req pb.RoleCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	res, err := ctrl.grpcClient.MRoleCreate(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (ctrl *rbacCtrl) RoleUpdate(ctx *gin.Context) {
	var req pb.RoleUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	list, err := ctrl.grpcClient.MRoleUpdate(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (ctrl *rbacCtrl) RoleDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	list, err := ctrl.grpcClient.MRoleDel(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (ctrl *rbacCtrl) RoleInfo(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	list, err := ctrl.grpcClient.MRoleInfo(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (ctrl *rbacCtrl) RolePermissionList(ctx *gin.Context) {
	var req pb.RolePermissionListReq
	id, _ := strconv.Atoi(ctx.Param("id"))
	req.RoleID = int64(id)
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	res, err := ctrl.grpcClient.MRolePermissionList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (ctrl *rbacCtrl) RolePermissionSet(ctx *gin.Context) {
	var req pb.RolePermissionSetReq
	id, _ := strconv.Atoi(ctx.Param("id"))
	req.RoleID = int64(id)
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}

	res, err := ctrl.grpcClient.MRolePermissionSet(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (ctrl *rbacCtrl) PathList(ctx *gin.Context) {
	var req pb.PathListReq
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := ctrl.grpcClient.MPathList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (ctrl *rbacCtrl) PathCreate(ctx *gin.Context) {
	var req pb.PathCreateReq
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}

	res, err := ctrl.grpcClient.MPathCreate(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (ctrl *rbacCtrl) PathUpdate(ctx *gin.Context) {
	var req pb.PathUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)
	if req.DomainID != -1 {
		req.DomainID = ctx.GetInt64(DOMIAN_ID)
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	list, err := ctrl.grpcClient.MPathUpdate(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (ctrl *rbacCtrl) PathDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	list, err := ctrl.grpcClient.MPathDel(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}
func (ctrl *rbacCtrl) PathInfo(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	list, err := ctrl.grpcClient.MPathInfo(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (ctrl *rbacCtrl) UserMenu(ctx *gin.Context) {
	var req pb.AdminMenuTreeReq

	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	value, _ := ctx.Get(ROLES)
	req.RoleIDs = value.([]int64)
	tree, err := ctrl.grpcClient.AdminMenuTree(ctx, &req)
	netHelper.Response(ctx, tree, err, nil)
}

func (ctrl *rbacCtrl) UserPermission(ctx *gin.Context) {
	var req pb.RolePermissionReq
	value, _ := ctx.Get(ROLES)
	req.RoleIDs = value.([]int64)
	dr, err := ctrl.grpcClient.RolePermission(ctx, &req)
	netHelper.Response(ctx, dr, err, nil)
}

func (ctrl *rbacCtrl) Login(ctx *gin.Context) {
	var req pb.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}

	res, err := ctrl.grpcClient.Login(ctx, &req)
	if err != nil {
		netHelper.Response(ctx, res, err, nil)
		return
	}
	fmt.Println(res)
	if res.Code != 200 {
		netHelper.Response(ctx, res, err, nil)
		return
	}
	ctrl.Auther.SetData(UID, res.Data.UID)
	ctrl.Auther.SetData(DOMIAN_ID, req.DomainID)
	ctrl.Auther.SetData(ROLES, res.Data.RoleIDs)
	t, err := ctrl.Auther.Token("")
	netHelper.Response(ctx, nil, err, map[string]interface{}{"token": t, "user_info": res.Data})

}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

func (ctrl *rbacCtrl) RefreshToken(ctx *gin.Context) {
	var req RefreshTokenReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	ntoken, err := ctrl.Auther.Token(req.RefreshToken)
	if err != nil {
		netHelper.Response(ctx, errors.NewUnauthorized(""), err, nil)
		return
	}
	netHelper.Response(ctx, nil, err, map[string]interface{}{"token": ntoken})
}

func (ctrl *rbacCtrl) UserInfo(ctx *gin.Context) {
	uid := ctx.GetInt64(UID)
	req := pb.DefaultPkReq{
		Pk:       &pb.DefaultPkReq_ID{ID: uid},
		DomainID: ctx.GetInt64(DOMIAN_ID),
	}
	lr, err := ctrl.grpcClient.AdminInfo(ctx, &req)
	netHelper.Response(ctx, lr, err, nil)
}

func (ctrl *rbacCtrl) AccountInfo(ctx *gin.Context) {
	uid := ctx.GetInt64(UID)
	req := pb.DefaultPkReq{
		Pk:       &pb.DefaultPkReq_ID{ID: uid},
		DomainID: ctx.GetInt64(DOMIAN_ID),
	}
	lr, err := ctrl.grpcClient.AccountInfo(ctx, &req)
	netHelper.Response(ctx, lr, err, nil)
}
func (ctrl *rbacCtrl) AccountEdit(ctx *gin.Context) {
	var req pb.AccountEditReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(DOMIAN_ID)
	req.UID = ctx.GetInt64(UID)
	res, err := ctrl.grpcClient.AccountEdit(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (ctrl *rbacCtrl) Permission(ctx *gin.Context) {
	roles, _ := ctx.Get(ROLES)

	req := pb.RolePermissionReq{
		DomainID: ctx.GetInt64(DOMIAN_ID),
		RoleIDs:  roles.([]int64),
	}
	lr, err := ctrl.grpcClient.RolePermission(ctx, &req)
	netHelper.Response(ctx, lr, err, nil)
}
