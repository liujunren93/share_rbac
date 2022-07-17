package middleware

import (
	"github.com/liujunren93/share_rbac/log"
	"github.com/liujunren93/share_rbac/service/api/ctrl"

	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_utils/common/auth"
	"github.com/liujunren93/share_utils/helper"
)

const ISLOGIN = "is_login"

var thisAuth auth.Auther

func UpdateAuther(Auth auth.Auther) {
	thisAuth = Auth

}
func Session(auther auth.Auther) func(*gin.Context) {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		authData, tp, err := thisAuth.Inspect(token)
		if err != nil || tp != 1 {
			ctx.Next()
			return
		}
		if data, ok := authData.(map[string]interface{}); ok {
			if domain, ok := data[ctrl.DOMIAN_ID]; ok {
				ctx.Set(ctrl.DOMIAN_ID, int64(domain.(float64)))
			}
			if roles, ok := data[ctrl.ROLES]; ok {
				if roles != nil {
					da, err := helper.InterfaceSlice2NumberSlice[float64](roles.([]interface{}))
					if err != nil {
						log.Logger.Error(err)
						return
					}
					ctx.Set(ctrl.ROLES, helper.TransSliceType[float64, int64](da))
				}

			}
			if uid, ok := data[ctrl.UID]; ok {
				ctx.Set(ctrl.UID, int64(uid.(float64)))
				if uid != 0 {
					ctx.Set(ISLOGIN, true)
				}
			}
		}
		ctx.Next()
	}
}
