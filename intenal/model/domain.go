package model

/**
* @Author: liujunren
* @Date: 2022/2/28 10:11
 */

type RbacDomain struct {
	ModelPL
	Status int8   `gorm:"status;type:int;not null;default:1;comment:'1:启用,2:禁用'" json:"status" binding:"required"`
	Name   string `gorm:"name;type:varchar(100);not null;default:'';comment:''" json:"name"  binding:"required"`
	Domain string `gorm:"domain;type:varchar(255);not null;default:'';comment:'域名'" json:"domain"  binding:"required"`
}
