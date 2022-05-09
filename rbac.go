package share_rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_rbac/intenal/dao"
	"github.com/liujunren93/share_rbac/router"
	"github.com/liujunren93/share_utils/auth"
	"gorm.io/gorm"
)

type rabc struct {
}

func RegisterRouter(group *gin.RouterGroup, auther auth.Auther) {
	router.InitRbacRoute(group, auther)
}

func NewRbacService(DB *gorm.DB) {
	dao.InitDB(DB)

}
