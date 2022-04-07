package model

/**
* @Author: liujunren
* @Date: 2022/2/28 10:51
 */

type RbacPath struct {
	Model
	Key       string `gorm:"key;type:varchar(100);not null;default:'';comment:'key'" json:"key"`
	Name      string `gorm:"name;type:varchar(100);not null;default:'';comment:'name'" json:"name"`
	Path      string `gorm:"path;type:varchar(100);not null;default:'';comment:'path'" json:"path"`
	Component string `gorm:"component;type:varchar(255);not null;default:'';comment:'component'" json:"component"`
	Redirect  string `gorm:"redirect;type:varchar(255);not null;default:'';comment:'redirect'" json:"redirect"`
	ParentID  uint   `gorm:"parent_id;type:int;not null;default:0;comment:'parent_id'" json:"parentID"`
	PathType  int8   `gorm:"path_type;type:tinyint(1);not null;default:1;comment:'1:生成菜单;2:api'" json:"path_type"`
	Meta      string `gorm:"meta;type:varchar(255);not null;default:'';comment:'meta{permission:}'"`
	Method    string `gorm:"method;type:varchar(20);not null;default:'';comment:'method'" json:"method"`
	Status    uint   `gorm:"status;type:int;not null;default:1;comment:'1:启用,2:禁用'" json:"status"`
}

type RbacPermission struct {
	Model
	DomainID uint   `gorm:"domain_id;type:int;not null;default:0;comment:'domain_id'"`
	Name     string `gorm:"name;type:varchar(100);not null;default:'';comment:'name'" json:"name"`
	Desc     string `gorm:"desc;type:varchar(100);not null;default:'';comment:'desc'" json:"desc"`
}

type RbacPermissionPath struct {
	Model
	DomainID     uint `gorm:"domain_id;type:int;not null;default:0;comment:'domain_id'"`
	PermissionID uint `gorm:"permission_id;type:int;not null;default:0;comment:'permission_id'"`
	PathID       uint `gorm:"path_id;type:int;not null;default:0;comment:'path_id'"`
}
