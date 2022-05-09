package test

import (
	"testing"

	"github.com/liujunren93/share_rbac/log"
	"github.com/liujunren93/share_utils/middleware"

	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_rbac"
	"github.com/liujunren93/share_rbac/service"
	"github.com/liujunren93/share_utils/auth"
	"github.com/liujunren93/share_utils/auth/jwt"
	"github.com/liujunren93/share_utils/databases/gorm"
	"github.com/sirupsen/logrus"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 16:19
 */

func init() {
	log.Logger = logrus.New()
}
func InitRbacDB() {
	mysql, err := gorm.NewMysql(&gorm.Mysql{
		LogMode:  true,
		Host:     "192.168.0.103",
		User:     "root",
		Password: "root",
		Port:     3306,
		Database: "test",
	}, nil)
	if err != nil {
		panic(err)
	}
	share_rbac.NewRbacService(mysql)
}
func initServer() {
	engine := gin.Default()
	engine.Use(middleware.Cors)
	group := engine.Group("rbac")

	share_rbac.RegisterRouter(group, jwt.NewAuth(auth.WithExpiry(7200), auth.WithSecret("www.sharelie.com")))
	engine.Run("0.0.0.0:9091")
}
func initgrpc() {
	service.NewGrpcService()
}

func TestInitGrpc(t *testing.T) {
	InitRbacDB()
	initgrpc()

}

func TestInitRbacRoute(t *testing.T) {
	initServer()
}
