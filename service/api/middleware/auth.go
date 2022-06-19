package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_utils/common/auth"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/netHelper"
)

func Auth(auther auth.Auther) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if ctx.GetBool(ISLOGIN) {
			ctx.Next()
		} else {
			// reflushToken := ctx.Request.Header.Get("ReflushToken")

			netHelper.Response(ctx, errors.NewUnauthorized(""), nil, nil)
			ctx.Abort()
			return
		}

	}

}
