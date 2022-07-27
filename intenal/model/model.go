package model

import (
	"errors"
	"strconv"
	"time"

	"github.com/liujunren93/share_utils/common/metadata"
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
	PL        int            `gorm:"pl;type:int;not null;default:0;comment:'permisssion level 只能操作>=pl的数据'" json:"pl"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	return m.checkPl(tx)
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	return m.checkPl(tx)
}

// 在同一个事务中更新数据
func (m *Model) AfterDelete(tx *gorm.DB) (err error) {
	return m.checkPl(tx)
}

func (m *Model) checkPl(tx *gorm.DB) error {
	pl, ok := metadata.GetVal(tx.Statement.Context, "rbac_pl")
	if ok {
		return errors.New("no data permission")
	}
	plint, err := strconv.Atoi(pl)
	if err != nil {
		return err
	}
	if m.PL < plint {
		return errors.New("no data permission")
	}
	return nil
}

type ModelSmp struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
