package model

import (
	"time"

	"gorm.io/gorm"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 10:12
 */
type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	DomainID  uint           `gorm:"domain;type:int;not null;default:0;comment:''" json:"domain_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ModelSmp struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
