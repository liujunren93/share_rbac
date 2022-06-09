package dao

import (
	model2 "github.com/liujunren93/share_rbac/intenal/model"
	"gorm.io/gorm"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 11:52
 */

var DB *gorm.DB

func DomianDB(domain int64) *gorm.DB {
	return DB.Where("domain_id=?", domain)
}

func pageSize(size int64) int {
	if size == 0 {
		return 25
	}

	return int(size)
}

func offset(size, offset int64) int {
	return pageSize(size) * int(offset)
}

func InitDB(db *gorm.DB) error {
	DB = db
	list := []interface{}{
		&model2.RbacAdmin{},
		&model2.RbacRole{},
		&model2.RbacPermission{},
		&model2.RbacRolePermission{},
		&model2.RbacPath{},
		&model2.RbacPermissionPath{},
		&model2.RbacDomain{},
		&model2.RbacRoleUser{},
	}
	return db.AutoMigrate(list...)
}
