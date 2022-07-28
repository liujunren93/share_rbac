package share_rbac

import (
	"context"

	"github.com/liujunren93/share_rbac/log"
	pb "github.com/liujunren93/share_rbac/rbac_pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"

	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/server"
	"github.com/liujunren93/share_rbac/intenal/dao"
	"github.com/liujunren93/share_rbac/service/api/ctrl"
	"github.com/liujunren93/share_rbac/service/api/middleware"
	"github.com/liujunren93/share_rbac/service/rpc"
	"github.com/liujunren93/share_utils/common/auth"
	"github.com/liujunren93/share_utils/common/gin/router"
	"github.com/liujunren93/share_utils/common/mq"
	utmiddleware "github.com/liujunren93/share_utils/middleware"
	"github.com/liujunren93/share_utils/wrapper/metadata"
	"gorm.io/gorm"
)

type Rbac struct {
	mq         mq.Mqer
	auther     auth.Auther
	grpcClient client.Client
}

func NewRbac(mq mq.Mqer) *Rbac {
	return &Rbac{mq: mq}
}

func (r *Rbac) SetLogger(logger *logrus.Logger) {

	log.InitLogger(logger)
	log.Logger.Debug("rbac.setlogger")
}
func (r *Rbac) UpAuther(auther auth.Auther) {
	log.Logger.Debug("UpAuther", auther)
	r.auther = auther
	ctrl.UpdateAuther(r.auther)

	middleware.UpdateAuther(auther)
}
func session(ctx context.Context) (proto.Message, error) {
	if ctx, ok := ctx.(*gin.Context); ok {
		domainId := ctx.GetInt64(pb.SESSION_SHARE_RBAC_DOMAIN_ID.String())
		uid := ctx.GetInt64(pb.SESSION_SHARE_RBAC_UID.String())
		pl := ctx.GetInt64(pb.SESSION_SHARE_RBAC_PL.String())
		roleIDs, ok := ctx.Get(pb.SESSION_SHARE_RBAC_ROLE_IDS.String())
		session := pb.Session{
			UID:      uid,
			DomainID: domainId,
			PL:       pl,
		}
		if ok {
			session.RoleIDs = roleIDs.([]int64)
		}
		return &session, nil

	} else {
		return nil, nil
	}

}
func (r *Rbac) NewApiService(ctx context.Context, engine *gin.Engine, auther auth.Auther, cli *client.Client, namespace, serverName string) (unLogin, Login router.Router, err error) {

	r.auther = auther
	cli.AddOptions(client.WithCallWrappers(metadata.NewClientWrapperMessage(pb.SESSION_SHARE_RBAC_METADATA_KEY.String(), session)))
	if namespace != "" {
		cli.AddOptions(client.WithNamespace(namespace))
	}

	cci, err := cli.Client(serverName)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	ctrl.InitRbacCtrl(ctx, r.auther, r.mq, cci)
	middleware.UpdateAuther(r.auther)
	return r.initRbacRoute(engine)
}

func (r *Rbac) NewGrpcService(DB *gorm.DB, ser *server.GrpcServer) error {
	err := dao.InitDB(DB)
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	pb.RegisterRbacServer(ser.Server(), rpc.NewRbacServer(r.mq))
	return nil
}

func (r *Rbac) initRbacRoute(engine *gin.Engine) (unLogin, login router.Router, err error) {

	unLogin = router.NewRouter(engine)
	var rbac = ctrl.RbacCtrl

	unLogin.Use(utmiddleware.Cors, middleware.Session)
	login = unLogin.Group("")
	login.Use(middleware.Auth, middleware.Rbac)
	rbacRouter := unLogin.Group("rbac")
	rbacRouter.Use(middleware.Rbac)

	auth := rbacRouter.Group("auth").White("rbac")
	{
		auth.POST("/login", rbac.Login)
		auth.POST("/register", rbac.Registry)
		auth.POST("/refreshToken", rbac.RefreshToken)
		auth.Use(middleware.Auth)

		auth.GET("/userInfo", rbac.UserInfo)
		auth.GET("/permission", rbac.Permission)
		auth.GET("/menu", rbac.UserMenu)
	}
	loginRouter := rbacRouter.Group("")
	loginRouter.Use(middleware.Auth)

	admin := loginRouter.Group("admin")
	{
		admin.GET("", rbac.AdminList)
		admin.POST("", rbac.AdminCreate)
		admin.PUT("/:id", rbac.AdminUpdate)
		admin.DELETE("/:id", rbac.AdminDel)
		admin.GET("/:id", rbac.AdminInfo)
	}

	loginRouter.GET("adminPermission", rbac.UserPermission)
	domian := loginRouter.Group("domain")

	{
		domian.GET("", rbac.DomainList)
		domian.POST("", rbac.DomainCreate)
		domian.PUT("/:id", rbac.DomainUpdate)
		domian.DELETE("/:id", rbac.DomainDel)
		domian.GET("/:id", rbac.DomainInfo)
	}
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
	return
}
