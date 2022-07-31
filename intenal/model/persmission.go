package model

/**
* @Author: liujunren
* @Date: 2022/2/28 10:51
 */

type RbacPath struct {
	ModelPL
	Key       string `gorm:"key;type:varchar(100);not null;default:'';comment:'前端权限组'" json:"key,omitempty"`
	Name      string `gorm:"name;type:varchar(100);not null;default:'';comment:'name'" json:"name,omitempty"`
	Action    string `gorm:"action;type:varchar(100);not null;default:'';comment:'方法:前端权限标识'" json:"action,omitempty"`
	Path      string `gorm:"path;type:varchar(100);not null;default:'';comment:'path'" json:"path,omitempty"`
	Component string `gorm:"component;type:varchar(255);not null;default:'';comment:'component'" json:"component,omitempty"`
	Redirect  string `gorm:"redirect;type:varchar(255);not null;default:'';comment:'redirect'" json:"redirect,omitempty"`
	ParentID  uint   `gorm:"parent_id;type:int;not null;default:0;comment:'parent_id'" json:"parent_id"`
	PathType  int8   `gorm:"path_type;type:tinyint(1);not null;default:1;comment:'1:前端;2:api'" json:"path_type,omitempty"`
	Meta      string `gorm:"meta;type:varchar(255);not null;default:'';comment:'meta{permission:}'" json:"meta,omitempty"`
	Method    string `gorm:"method;type:varchar(20);not null;default:'';comment:'api:method'" json:"method,omitempty"`
	ApiPath   string `gorm:"api_path;type:varchar(100);not null;default:'';comment:'api:api_path'" json:"api_path"`
	Status    uint   `gorm:"status;type:int;not null;default:1;comment:'1:启用,2:禁用'" json:"status,omitempty"`
	DomainID  int    `gorm:"domain_id;index;type:int;not null;default:-1;comment:'-1:公共'" json:"-"`
	Sort      int    `gorm:"sort;type:int;not null;default:9999" json:"sort"`
}

type RbacPermission struct {
	ModelPL
	// Type     int8   `gorm:"domain_id;type:tinyint(1);not null;default:1;comment:'1:普通权限,2:action;3:action+普通权限'" json:"-"`
	DomainID int    `gorm:"domain_id;index;type:int;not null;default:0;comment:'-1:基础权限，所有域通用'" json:"-"`
	Name     string `gorm:"name;type:varchar(100);not null;default:'';comment:'name'" json:"name"`
	Status   uint   `gorm:"status;type:int;not null;default:1;comment:'1:启用,2:禁用'" json:"status,omitempty"`
	Desc     string `gorm:"desc;type:varchar(100);not null;default:'';comment:'desc'" json:"desc"`
}

type RbacPermissionPath struct {
	ModelPL
	DomainID     int  `gorm:"domain_id;type:int;not null;default:0;comment:'domain_id'"`
	PermissionID uint `gorm:"permission_id;type:int;not null;default:0;comment:'permission_id'"`
	PathID       uint `gorm:"path_id;type:int;not null;default:0;comment:'path_id'"`
}
