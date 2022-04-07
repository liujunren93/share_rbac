package model

/**
* @Author: liujunren
* @Date: 2022/2/28 10:11
 */

type RbacAdmin struct {
	Model
	DomainID uint   `gorm:"domain_id;type:int;not null;default:0;comment:'domain_id'"`
	Account  string `gorm:"account;type:varchar(100);not null;default:'';comment:''" json:"account" binding:"required"`
	Name     string `gorm:"name;type:varchar(100);not null;default:'';comment:''" json:"name"  binding:"required"`
	Password string `gorm:"password;type:varchar(100);not null;default:'';comment:''"  json:"-"  binding:"required"`
	Status   uint   `gorm:"status;type:int;not null;default:1;comment:'1:启用,2:禁用'" json:"status"`
}
