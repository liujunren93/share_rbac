package ctrl

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_rbac/middleware"
	"github.com/liujunren93/share_rbac/pb"
	"github.com/liujunren93/share_utils/auth"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/netHelper"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 11:44
 */
type Rbac struct {
	Auther auth.Auther
}

func (*Rbac) DomainList(ctx *gin.Context) {
	var req pb.DomainListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.MDomainList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}
func (*Rbac) DomainInfo(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))

	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}

	info, err := client.MDomainInfo(ctx, &req)

	netHelper.Response(ctx, info, err, nil)
}
func (*Rbac) DomainDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))

	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}

	info, err := client.MDomainDel(ctx, &req)
	fmt.Println(info, err)
	netHelper.Response(ctx, info, err, nil)

}
func (*Rbac) DomainCreate(ctx *gin.Context) {
	var req pb.DomainCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	create, err := client.MDomainCreate(ctx, &req)
	netHelper.Response(ctx, create, err, nil)
}
func (*Rbac) DomainUpdate(ctx *gin.Context) {
	var req pb.DomainUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	dr, err := client.MDomainUpdate(ctx, &req)
	netHelper.Response(ctx, dr, err, nil)

}

func (*Rbac) AdminList(ctx *gin.Context) {
	var req pb.AdminListReq
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}

	res, err := client.MAdminList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)

}

func (*Rbac) AdminCreate(ctx *gin.Context) {
	var req pb.AdminCreateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	create, err := client.MAdminCreate(ctx, &req)
	netHelper.Response(ctx, create, err, nil)
}

func (*Rbac) AdminUpdate(ctx *gin.Context) {
	var req pb.AdminUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}

	update, err := client.MAdminUpdate(ctx, &req)
	netHelper.Response(ctx, update, err, nil)

}

func (*Rbac) AdminDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	update, err := client.MAdminDel(ctx, &req)
	netHelper.Response(ctx, update, err, nil)
}

func (*Rbac) AdminInfo(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}

	update, err := client.MAdminInfo(ctx, &req)
	netHelper.Response(ctx, update, err, nil)
}

func (*Rbac) AdminRoleList(ctx *gin.Context) {
	var req pb.AdminRoleListReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	req.RoleID = int64(atoi)

	update, err := client.MAdminRoleList(ctx, &req)
	netHelper.Response(ctx, update, err, nil)
}

func (*Rbac) AdminRoleSet(ctx *gin.Context) {
	var req pb.AdminRoleSetReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.AdminID = int64(atoi)
	update, err := client.MAdminRoleSet(ctx, &req)
	netHelper.Response(ctx, update, err, nil)
}

func (*Rbac) PermissionList(ctx *gin.Context) {
	var req pb.PermissionListReq
	// req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	list, err := client.MPermissionList(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) PermissionInfo(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}

	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	list, err := client.MPermissionInfo(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) PermissionCreate(ctx *gin.Context) {
	var req pb.PermissionCreateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	list, err := client.MPermissionCreate(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) PermissionUpdate(ctx *gin.Context) {
	var req pb.PermissionUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	list, err := client.MPermissionUpdate(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) PermissionDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	list, err := client.MPermissionDel(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) PermissionPathList(ctx *gin.Context) {
	var req pb.PermissionPathListReq

	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	req.PermissionID = int64(id)
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	res, err := client.MPermissionPathList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) PermissionPathSet(ctx *gin.Context) {
	var req pb.PermissionPathSetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	req.PermissionID = int64(id)
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	res, err := client.MPermissionPathSet(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) RoleList(ctx *gin.Context) {
	var req pb.RoleListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	res, err := client.MRoleList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) RoleCreate(ctx *gin.Context) {
	var req pb.RoleCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	res, err := client.MRoleCreate(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) RoleUpdate(ctx *gin.Context) {
	var req pb.RoleUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	list, err := client.MRoleUpdate(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) RoleDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	list, err := client.MRoleDel(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) RoleInfo(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	list, err := client.MRoleInfo(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) RolePermissionList(ctx *gin.Context) {
	var req pb.RolePermissionListReq
	id, _ := strconv.Atoi(ctx.Param("id"))
	req.RoleID = int64(id)
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	res, err := client.MRolePermissionList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) RolePermissionSet(ctx *gin.Context) {
	var req pb.RolePermissionSetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	req.RoleID = int64(id)
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	res, err := client.MRolePermissionSet(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) PathList(ctx *gin.Context) {
	var req pb.PathListReq
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.MPathList(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) PathCreate(ctx *gin.Context) {
	var req pb.PathCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	if req.DomainID != -1 {
		req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	}

	res, err := client.MPathCreate(ctx, &req)
	netHelper.Response(ctx, res, err, nil)
}

func (*Rbac) PathUpdate(ctx *gin.Context) {
	var req pb.PathUpdateReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.ID = int64(atoi)
	if req.DomainID != -1 {
		req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	list, err := client.MPathUpdate(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) PathDel(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	list, err := client.MPathDel(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}
func (*Rbac) PathInfo(ctx *gin.Context) {
	var req pb.DefaultPkReq
	atoi, _ := strconv.Atoi(ctx.Param("id"))
	req.Pk = &pb.DefaultPkReq_ID{ID: int64(atoi)}
	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	list, err := client.MPathInfo(ctx, &req)
	netHelper.Response(ctx, list, err, nil)
}

func (*Rbac) UserMenu(ctx *gin.Context) {
	var req pb.AdminMenuTreeReq

	req.DomainID = ctx.GetInt64(middleware.DOMIAN_ID)
	value, _ := ctx.Get(middleware.ROLES)
	req.RoleIDs = value.([]int64)
	tree, err := client.AdminMenuTree(ctx, &req)
	netHelper.Response(ctx, tree, err, nil)
}

func (*Rbac) UserPermission(ctx *gin.Context) {
	var req pb.RolePermissionReq
	value, _ := ctx.Get(middleware.ROLES)
	req.RoleIDs = value.([]int64)
	dr, err := client.RolePermission(ctx, &req)
	netHelper.Response(ctx, dr, err, nil)
}

func (r *Rbac) Login(ctx *gin.Context) {
	var req pb.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		netHelper.Response(ctx, errors.StatusBadRequest, err, nil)
		return
	}
	res, err := client.Login(ctx, &req)
	if err != nil {
		netHelper.Response(ctx, res, err, nil)
		return
	}
	if res.Code != 200 {
		netHelper.Response(ctx, res, err, nil)
		return
	}
	r.Auther.SetData(middleware.UID, res.Data.UID)
	r.Auther.SetData(middleware.DOMIAN_ID, req.DomainID)
	r.Auther.SetData(middleware.ROLES, res.Data.RoleIDs)
	t, err := r.Auther.Token("")
	netHelper.Response(ctx, nil, err, map[string]interface{}{"token": t, "user_info": res.Data})
	return

}

func (*Rbac) UserInfo(ctx *gin.Context) {
	uid := ctx.GetInt64(middleware.UID)
	req := pb.DefaultPkReq{
		Pk:       &pb.DefaultPkReq_ID{ID: uid},
		DomainID: ctx.GetInt64(middleware.DOMIAN_ID),
	}
	lr, err := client.AdminInfo(ctx, &req)
	netHelper.Response(ctx, lr, err, nil)
	return
}
