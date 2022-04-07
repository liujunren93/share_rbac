package share_rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_rbac/intenal/dao"
	"github.com/liujunren93/share_rbac/service/ctrl"
	"gorm.io/gorm"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 11:39
 */

func InitRbacRoute(group *gin.RouterGroup) {
	var rbac = new(ctrl.Rbac)

	domian := group.Group("domain")
	{
		domian.GET("", rbac.DomainList)
		domian.POST("", rbac.DomainCreate)
		domian.PUT("/:id", rbac.DomainUpdate)
		domian.DELETE("/:id", rbac.DomainDel)
		domian.GET("/:id", rbac.DomainInfo)
	}
	admin := group.Group("admin")
	{
		admin.GET("", rbac.AdminList)
		admin.POST("", rbac.AdminCreate)
		admin.PUT("/:id", rbac.AdminUpdate)
		admin.DELETE("/:id", rbac.AdminDel)
	}

	permission := group.Group("permission")
	{
		permission.GET("", rbac.PermissionList)
		permission.POST("", rbac.PermissionCreate)
		permission.PUT("/:id", rbac.PermissionUpdate)
		permission.DELETE("/:id", rbac.PermissionDel)
	}

	permissionPath := group.Group("permissionPath")
	{
		permissionPath.GET("", rbac.PermissionPathList)
		permissionPath.POST("", rbac.PermissionPathSet)

	}

	role := group.Group("role")
	{
		role.GET("", rbac.RoleList)
		role.POST("", rbac.RoleCreate)
		role.PUT("/:id", rbac.RoleUpdate)
		role.DELETE("/:id", rbac.RoleDel)
	}
	rolePermission := group.Group("rolePermission")
	{
		rolePermission.GET("", rbac.RolePermissionList)
		rolePermission.POST("", rbac.RolePermissionSet)
	}

	path := group.Group("path")
	{
		path.GET("", rbac.PathList)
		path.POST("", rbac.PathCreate)
		path.PUT("/:id", rbac.PathUpdate)
		path.DELETE("/:id", rbac.PathDel)
	}

}

func InitRbacDB(db *gorm.DB) {
	dao.InitDB(db)
}
