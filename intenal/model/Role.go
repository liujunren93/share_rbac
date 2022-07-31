package model

/**
* @Author: liujunren
* @Date: 2022/2/28 11:01
 */

type RbacRole struct {
	ModelPL
	DomainID uint   `gorm:"domain_id;type:int;not null;default:0;comment:'domain_id'"`
	Name     string `gorm:"name;type:varchar(100);not null;default:'';comment:'name'" json:"name"`
	Desc     string `gorm:"desc;type:varchar(100);not null;default:'';comment:'desc'" json:"desc"`
	Status   uint   `gorm:"status;type:int;not null;default:1;comment:'1:启用,2:禁用'" json:"status"`
}

type RbacRoleUser struct {
	ModelPL
	DomainID uint `gorm:"domain_id;type:int;not null;default:0;comment:'domain_id'" json:"domain_id"`
	RoleID   uint `gorm:"role_id;type:int;not null;default:0;comment:'role_id'" json:"role_id"`
	UID      uint `gorm:"uid;type:int;not null;default:0;comment:'uid'" `
}

type RbacRolePermission struct {
	ModelPL
	DomainID     uint `gorm:"domain_id;type:int;not null;default:0;comment:'domain_id'"`
	RoleID       uint `gorm:"role_id;type:int;not null;default:0;comment:'role_id'"`
	PermissionID uint `gorm:"permission_id;type:int;not null;default:0;comment:'permission_id'"`
}
