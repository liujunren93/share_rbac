package test

import (
	"context"
	"net/http"
	"testing"

	re "github.com/go-redis/redis/v8"
	"github.com/liujunren93/share_rbac/log"
	"github.com/liujunren93/share_utils/client/grpc"
	"github.com/liujunren93/share_utils/middleware"
	"github.com/liujunren93/share_utils/server"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_rbac"

	_ "net/http/pprof"

	"github.com/liujunren93/share_utils/common/auth"
	"github.com/liujunren93/share_utils/common/auth/jwt"
	"github.com/liujunren93/share_utils/common/mq/redis"
	g "github.com/liujunren93/share_utils/databases/gorm"
	dbredis "github.com/liujunren93/share_utils/databases/redis"
	"github.com/sirupsen/logrus"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 16:19
 */
var app *share_rbac.Rbac

func init() {
	r, _ := dbredis.NewRedis(&re.Options{
		Network: "tcp",
		Addr:    "node1:6379",
	})
	app = share_rbac.NewRbac(redis.NewMq(r))
	log.Logger = logrus.New()
}
func InitRbacDB() *gorm.DB {
	mysql, err := g.NewMysql(&g.Mysql{
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
	return mysql
}
func initServer() {
	engine := gin.Default()
	engine.Use(middleware.Cors)

	c := grpc.NewClient(grpc.WithBuildTargetFunc(func(args ...string) string { return "127.0.0.1:19091" }))
	shareClient, err := c.GetShareClient()
	if err != nil {
		panic(err)
	}

	app.NewApiService(context.TODO(), engine, jwt.NewAuth(auth.WithExpiry(1000), auth.WithSecret("www.sharelie.com")), shareClient, "rbac")
	engine.Run(":9091")
}

func TestInitGrpc(t *testing.T) {
	go http.ListenAndServe("0.0.0.0:6061", nil)
	d := InitRbacDB()
	s := server.Server{Address: "0.0.0.0:19091", Mode: "debug"}
	gs, err := s.NewServer()
	if err != nil {
		panic(err)
	}
	app.NewGrpcService(d, gs)

}

func TestInitRbacRoute(t *testing.T) {
	go http.ListenAndServe("0.0.0.0:6060", nil)
	initServer()
}
