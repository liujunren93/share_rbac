package middleware

import (
	"github.com/gin-gonic/gin"
	pb "github.com/liujunren93/share_rbac/rbac_pb"
	"github.com/liujunren93/share_rbac/service/api/ctrl"
	"github.com/liujunren93/share_utils/common/gin/router"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/netHelper"
)

func Rbac(ctx *gin.Context) {
	if !ctx.GetBool(ISLOGIN) {
		ctx.Next()
		return
	}
	if router.InWhitelist(ctx, "rbac") {
		ctx.Next()
		return
	}

	uid := ctx.GetInt64(pb.SESSION_UID.String())
	domainId := ctx.GetInt64(pb.SESSION_DOMAIN_ID.String())
	var roleIds []int64
	if v, ok := ctx.Get(pb.SESSION_ROLE_IDS.String()); ok {
		roleIds = v.([]int64)
	}
	err := ctrl.RbacCtrl.CheckPermission(ctx, ctx.Request.URL.Path, ctx.Request.Method, roleIds, uid, domainId)
	if err != nil {
		netHelper.Response(ctx, errors.NewForbidden(""), err, nil)
		ctx.Abort()
		return

	}
	ctx.Next()
}

func monitorRabc(ctx *gin.Context) {
	// ctrl.RbacCtrl{}.CheckPermission()
}
