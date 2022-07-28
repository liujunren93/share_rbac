package dao

import (
	"context"

	model2 "github.com/liujunren93/share_rbac/intenal/model"
	"github.com/liujunren93/share_rbac/rbac_pb"
	"github.com/liujunren93/share_utils/common/metadata"
	"gorm.io/gorm"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 11:52
 */

var Db *gorm.DB

type dao struct {
	Ctx context.Context
}

func (d dao) GetSession() *rbac_pb.Session {
	var sess *rbac_pb.Session
	metadata.GetMessage(d.Ctx, rbac_pb.SESSION_SHARE_RBAC_METADATA_KEY.String(), sess)
	return sess
}

func DB(ctx context.Context) *gorm.DB {
	return Db.WithContext(ctx)
}

// func getPL(ctx context.Context) (int, error) {
// 	pl, ok := metadata.GetVal(ctx, rbac_pb.SESSION_SHARE_RBAC_METADATA_KEY.String())
// 	if ok {
// 		return 0, errors.New("no data permission")
// 	}
// 	plint, err := strconv.Atoi(pl)
// 	if err != nil {
// 		return plint, err
// 	}
// 	if m.PL < plint {
// 		return errors.New("no data permission")
// 	}
// 	return nil
// }

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
	Db = db
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
