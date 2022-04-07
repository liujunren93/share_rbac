package test

import (
	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_rbac"
	"github.com/liujunren93/share_rbac/service"
	"github.com/liujunren93/share_utils/databases/gorm"
	"testing"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 16:19
 */

func InitRbacDB() {
	mysql, err := gorm.NewMysql(&gorm.Mysql{
		LogMode:  true,
		Host:     "node1",
		User:     "root",
		Password: "root",
		Port:     3306,
		Database: "test",
	}, nil)
	if err != nil {
		panic(err)
	}
	share_rbac.InitRbacDB(mysql)
}
func initServer()  {
	engine := gin.Default()
	group := engine.Group("rbac")
	share_rbac.InitRbacRoute(group)
	engine.Run(":9092")
}
func initgrpc()  {
	service.NewGrpcService()
}

func TestInitGrpc(t *testing.T) {
	InitRbacDB()
	 initgrpc()


}

func TestInitRbacRoute(t *testing.T) {
	 initServer()
}


