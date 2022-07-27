package test

import (
	"net/http"
	"testing"

	"github.com/liujunren93/share_utils/server"
	"gorm.io/gorm"

	"github.com/liujunren93/share_rbac"

	_ "net/http/pprof"

	g "github.com/liujunren93/share_utils/databases/gorm"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 16:19
 */
var app *share_rbac.Rbac

func init() {
	// r, _ := dbredis.NewRedis(&re.Options{
	// 	Network: "tcp",
	// 	Addr:    "node1:6379",
	// })
	// app = share_rbac.NewRbac(redis.NewMq(r), jwt.NewAuth(auth.WithExpiry(10000), auth.WithSecret("www.sharelie.com")))
	// log.Logger = logrus.New()
}
func InitRbacDB() *gorm.DB {
	mysql, err := g.NewMysql(&g.Mysql{
		Debug:    true,
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
	// engine := gin.Default()
	// engine.Use(middleware.Cors)

	// c := grpc.NewClient(grpc.WithBuildTargetFunc(func(args ...string) string { return "127.0.0.1:19092" }))
	// shareClient, err := c.GetShareClient()
	// if err != nil {
	// 	panic(err)
	// }
	// go func() {
	// 	var i int
	// 	for {
	// 		time.Sleep(time.Second * 20)
	// 		i++
	// 		app.UpAuther(jwt.NewAuth(auth.WithExpiry(100000), auth.WithSecret(strconv.Itoa(i))))
	// 	}

	// }()

	// // app.NewApiService(context.TODO(), engine, shareClient, "", "rbac")
	// engine.Run(":9091")
}

func TestInitGrpc(t *testing.T) {
	go http.ListenAndServe("0.0.0.0:6061", nil)
	d := InitRbacDB()
	s := server.Server{Address: "0.0.0.0:19092", Mode: "debug"}
	gs, err := s.NewServer()
	if err != nil {
		panic(err)
	}
	app.NewGrpcService(d, gs)
	gs.Run()

}

func TestInitRbacRoute(t *testing.T) {
	go http.ListenAndServe("0.0.0.0:6060", nil)
	initServer()
}
