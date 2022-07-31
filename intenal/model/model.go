package model

import (
	"strconv"
	"strings"
	"time"

	"github.com/liujunren93/share_rbac/rbac_pb"
	"github.com/liujunren93/share_utils/common/metadata"
	shErr "github.com/liujunren93/share_utils/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 10:12
 */
type ModelPL struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	DomainID  uint           `gorm:"domain;type:int;not null;default:0;comment:''" json:"domain_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	PL        string         `gorm:"pl;type:varchar(20);not null;default:'';comment:'level_uid,权限锁管理员只能操作数据pl>管理员pl的数据 和自己锁上的数据'" json:"pl"`
}

func (m *ModelPL) getPL() (level, uid int) {
	if len(m.PL) > 0 {
		pl := strings.Split(m.PL, "_")
		if len(pl) != 2 {
			return 0, 0
		}
		level, _ := strconv.Atoi(pl[0])
		uid, _ := strconv.Atoi(pl[1])
		return level, uid
	}
	return 0, 0
}

func modelInfo(tx *gorm.DB) ModelPL {
	var mode ModelPL
	// buf := clause.Builder{}
	// for range s.Clauses {}
	var wheres []clause.Expr
	var ors []clause.Expr
	for _, c := range tx.Statement.Clauses["WHERE"].Expression.(clause.Where).Exprs {

		if exp, ok := c.(clause.Expr); ok {
			wheres = append(wheres, exp)

		}
		if or, ok := c.(clause.OrConditions); ok {
			for _, oc := range or.Exprs {
				if exp, ok := oc.(clause.Expr); ok {
					ors = append(ors, exp)
				}
			}

		}
	}

	tx = tx.Select("id,pl").Table(tx.Statement.Table).Where("pl!=''")
	for _, where := range wheres {
		tx = tx.Where(where.SQL, where.Vars...)
	}

	for _, or := range ors {
		tx = tx.Or(or.SQL, or.Vars...)
	}
	tx.First(&mode)

	return mode
}

func (m *ModelPL) BeforeUpdate(tx *gorm.DB) (err error) {

	return m.checkPl(tx)
}

// 在同一个事务中更新数据
func (m *ModelPL) BeforeDelete(tx *gorm.DB) (err error) {
	return m.checkPl(tx)
}

// func (m *Model) AfterFind(tx *gorm.DB) (err error) {
// 	var session rbac_pb.Session
// 	err, ok := metadata.GetMessage(tx.Statement.Context, rbac_pb.SESSION_SHARE_RBAC_METADATA_KEY.String(), &session)

// 	fmt.Printf("session:%+v", &session)
// 	return
// }
func (m *ModelPL) checkPl(tx *gorm.DB) error {
	var session rbac_pb.Session
	model := modelInfo(tx)
	if model.PL == "" {
		return nil
	}
	err, ok := metadata.GetMessage(tx.Statement.Context, rbac_pb.SESSION_SHARE_RBAC_METADATA_KEY.String(), &session)
	if err != nil || !ok {
		return shErr.New(shErr.StatusDBPermissionDenied, "permission denied")
	}
	level, uid := model.getPL()
	if int(session.UID) == uid || int(session.PL) < level {
		return nil
	}

	return shErr.NewPublic(shErr.StatusDBPermissionDenied, "permission denied")
}

type ModelSmp struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	PL        uint           `gorm:"pl;type:int;not null;default:0;comment:'permission level'" json:"pl"`
}
