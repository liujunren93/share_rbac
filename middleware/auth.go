package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_utils/auth"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/helper"
	"github.com/liujunren93/share_utils/netHelper"
)

const (
	DOMIAN_ID = "domian_id"
	UID       = "uid"
	ROLES     = "roles" //[]int64
)

func Auth(auther auth.Auther) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		authData, err := auther.Inspect(token)
		if err != nil {
			netHelper.Response(ctx, errors.NewUnauthorized(""), nil, nil)
			return
		}
		if data, ok := authData.(map[string]interface{}); ok {
			if domain, ok := data[DOMIAN_ID]; ok {
				ctx.Set(DOMIAN_ID, int64(domain.(float64)))
			}
			if roles, ok := data[ROLES]; ok {
				da, err := helper.InterfaceSlice2NumberSlice[float64](roles.([]interface{}))
				if err != nil {
					goto Unauthorized
				}
				if len(da) == 0 {
					goto Unauthorized
				}
				ctx.Set(ROLES, helper.TransSliceType[float64, int64](da))
			}
			if uid, ok := data[UID]; ok {
				ctx.Set(UID, int64(uid.(float64)))
				if uid == 0 {
					goto Unauthorized
				}
			}
			ctx.Next()
			return

		}
	Unauthorized:
		netHelper.Response(ctx, errors.NewUnauthorized(""), nil, nil)
		return
	}

}
