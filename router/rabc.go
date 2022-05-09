package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_rbac/middleware"
	"github.com/liujunren93/share_rbac/service/ctrl"
	"github.com/liujunren93/share_utils/auth"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 11:39
 */

func InitRbacRoute(group *gin.RouterGroup, auther auth.Auther) {
	var rbac = ctrl.Rbac{Auther: auther}
	domian := group.Group("domain")
	{
		domian.GET("", rbac.DomainList)
		domian.POST("", rbac.DomainCreate)
		domian.PUT("/:id", rbac.DomainUpdate)
		domian.DELETE("/:id", rbac.DomainDel)
		domian.GET("/:id", rbac.DomainInfo)
	}
	auth := group.Group("auth")
	{
		auth.POST("/login", rbac.Login)
		auth.Use(middleware.Auth(auther))
		auth.GET("/userInfo", rbac.UserInfo)
		auth.GET("/menu", rbac.UserMenu)
	}

	group.Use(middleware.Auth(auther))

	admin := group.Group("admin")
	{
		admin.GET("", rbac.AdminList)
		admin.POST("", rbac.AdminCreate)
		admin.PUT("/:id", rbac.AdminUpdate)
		admin.DELETE("/:id", rbac.AdminDel)
		admin.GET("/:id", rbac.AdminInfo)
	}

	group.GET("adminPermission", rbac.UserPermission)

	adminRole := group.Group("adminRole")
	{
		adminRole.GET("/:id", rbac.AdminRoleList)
		adminRole.POST("/:id", rbac.AdminRoleSet)

	}

	permission := group.Group("permission")
	{
		permission.GET("", rbac.PermissionList)
		permission.POST("", rbac.PermissionCreate)
		permission.PUT("/:id", rbac.PermissionUpdate)
		permission.GET("/:id", rbac.PermissionInfo)
		permission.DELETE("/:id", rbac.PermissionDel)
	}

	permissionPath := group.Group("permissionPath")
	{
		permissionPath.GET("/:id", rbac.PermissionPathList)
		permissionPath.PUT("/:id", rbac.PermissionPathSet)

	}

	role := group.Group("role")
	{
		role.GET("", rbac.RoleList)
		role.POST("", rbac.RoleCreate)
		role.PUT("/:id", rbac.RoleUpdate)
		role.GET("/:id", rbac.RoleInfo)
		role.DELETE("/:id", rbac.RoleDel)
	}
	rolePermission := group.Group("rolePermission")
	{
		rolePermission.GET("/:id", rbac.RolePermissionList)
		rolePermission.POST("/:id", rbac.RolePermissionSet)
	}

	path := group.Group("path")
	{
		path.GET("", rbac.PathList)
		path.POST("", rbac.PathCreate)
		path.PUT("/:id", rbac.PathUpdate)
		path.DELETE("/:id", rbac.PathDel)
		path.GET("/:id", rbac.PathInfo)
	}
	group.GET("menu", rbac.UserMenu)

}
