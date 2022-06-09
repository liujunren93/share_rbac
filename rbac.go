package share_rbac

import (
	"context"

	"github.com/liujunren93/share_rbac/log"

	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/server"
	"github.com/liujunren93/share_rbac/intenal/dao"
	"github.com/liujunren93/share_rbac/pb"
	"github.com/liujunren93/share_rbac/service/api/ctrl"
	"github.com/liujunren93/share_rbac/service/api/middleware"
	"github.com/liujunren93/share_rbac/service/rpc"
	"github.com/liujunren93/share_utils/common/auth"
	"github.com/liujunren93/share_utils/common/gin/router"
	"github.com/liujunren93/share_utils/common/mq"
	"github.com/liujunren93/share_utils/wrapper/breaker"
	"github.com/liujunren93/share_utils/wrapper/metadata"
	"gorm.io/gorm"
)

type Rbac struct {
	serverName string
	mq         mq.Mqer
	grpcClient client.Client
}

func NewRbac(mq mq.Mqer) *Rbac {
	return &Rbac{mq: mq}
}

func session(ctx context.Context) string {
	if ctx, ok := ctx.(*gin.Context); ok {
		domainId := ctx.GetInt64(pb.SESSION_DOMAIN_ID.String())
		uid := ctx.GetInt64(pb.SESSION_UID.String())
		roleIDs, ok := ctx.Get(pb.SESSION_ROLE_IDS.String())
		session := pb.Session{
			UID:      uid,
			DomainID: domainId,
		}
		if ok {
			session.RoleIDs = roleIDs.([]int64)
		}
		return session.String()
	} else {
		return ""
	}

}
func (r *Rbac) NewApiService(group *gin.RouterGroup, auther auth.Auther, cli *client.Client) {
	cli.AddOptions(client.WithCallWrappers(metadata.NewClientWrapper("rbac_session", session)), client.WithCallWrappers(breaker.NewClientWrapper()))
	cci, err := cli.Client(r.serverName)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	ctrl.InitRbacCtrl(auther, r.mq, cci)
	r.initRbacRoute(group, auther)
}

func (*Rbac) NewGrpcService(DB *gorm.DB, ser *server.GrpcServer) {
	dao.InitDB(DB)
	pb.RegisterRbacServer(ser.Server(), rpc.NewRbacServer())
	ser.Run()
}

func (r *Rbac) initRbacRoute(routerGroup *gin.RouterGroup, auther auth.Auther) (noAuth, auth router.RouterGroup, err error) {
	session := router.NewRouterGroup(routerGroup)

	var rbac = ctrl.RbacCtrl
	session.Use(middleware.Session(auther), middleware.Rbac)
	domian := session.Group("domain")
	{
		domian.GET("", rbac.DomainList)
		domian.POST("", rbac.DomainCreate)
		domian.PUT("/:id", rbac.DomainUpdate)
		domian.DELETE("/:id", rbac.DomainDel)
		domian.GET("/:id", rbac.DomainInfo)
	}

	rauth := session.Group("auth").White("rbac")
	{
		rauth.POST("/login", rbac.Login)
		rauth.Use(middleware.Auth(auther))
		rauth.GET("/userInfo", rbac.UserInfo)
		rauth.GET("/permission", rbac.Permission)
		rauth.GET("/menu", rbac.UserMenu)
	}
	loginRouter := session.Group("")
	loginRouter.Use(middleware.Auth(auther))

	admin := loginRouter.Group("admin")
	{
		admin.GET("", rbac.AdminList)
		admin.POST("", rbac.AdminCreate)
		admin.PUT("/:id", rbac.AdminUpdate)
		admin.DELETE("/:id", rbac.AdminDel)
		admin.GET("/:id", rbac.AdminInfo)
	}

	loginRouter.GET("adminPermission", rbac.UserPermission)

	adminRole := loginRouter.Group("adminRole")
	{
		adminRole.GET("/:id", rbac.AdminRoleList)

	}
	account := loginRouter.Group("account")
	{
		account.PUT("base", rbac.AccountEdit)
		account.GET("base", rbac.AccountInfo).White("rbac")
	}

	permission := loginRouter.Group("permission")
	{
		permission.GET("", rbac.PermissionList)
		permission.POST("", rbac.PermissionCreate)
		permission.PUT("/:id", rbac.PermissionUpdate)
		permission.GET("/:id", rbac.PermissionInfo)
		permission.DELETE("/:id", rbac.PermissionDel)
	}

	permissionPath := loginRouter.Group("permissionPath")
	{
		permissionPath.GET("/:id", rbac.PermissionPathList)
		permissionPath.PUT("/:id", rbac.PermissionPathSet)
	}

	role := loginRouter.Group("role")
	{
		role.GET("", rbac.RoleList)
		role.POST("", rbac.RoleCreate)
		role.PUT("/:id", rbac.RoleUpdate)
		role.GET("/:id", rbac.RoleInfo)
		role.DELETE("/:id", rbac.RoleDel)
	}
	rolePermission := loginRouter.Group("rolePermission")
	{
		rolePermission.GET("/:id", rbac.RolePermissionList)
		rolePermission.PUT("/:id", rbac.RolePermissionSet)
	}

	path := loginRouter.Group("path")
	{
		path.GET("", rbac.PathList)
		path.POST("", rbac.PathCreate)
		path.PUT("/:id", rbac.PathUpdate)
		path.DELETE("/:id", rbac.PathDel)
		path.GET("/:id", rbac.PathInfo)
	}
	loginRouter.GET("menu", rbac.UserMenu)
	return session, loginRouter, nil
}
