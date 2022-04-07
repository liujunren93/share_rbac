package dao

import (
	"github.com/liujunren93/share_rbac/intenal/entity"
	"github.com/liujunren93/share_rbac/intenal/model"
	"github.com/liujunren93/share_rbac/log"
	"github.com/liujunren93/share_rbac/pb"
	"github.com/liujunren93/share_utils/commond/set"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/helper"
	"gorm.io/gorm"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 11:45
 */
type Permission struct {
}

func (dao Permission) List(req *pb.PermissionListReq) entity.PermissionListRes {
	var res entity.PermissionListRes
	db := DB.Where("domain_id=-1 or do_main=?", req.DomainID)
	res.List = dao.list(db, req.Name, req.PageSize, req.Page)
	res.Total = dao.count(db, req.Name)

	return res
}

func (Permission) count(db *gorm.DB, name string, ids ...uint) int64 {

	if name == "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	if len(ids) > 0 {
		db = db.Where("id in ?", ids)
	}
	var total int64
	db.Model(&model.RbacPermission{}).Count(&total)
	return total
}

func (Permission) list(db *gorm.DB, name string, limit, page int64, ids ...uint) []model.RbacPermission {
	var list []model.RbacPermission

	if limit >= 0 {
		db = db.Limit(pageSize(limit)).Offset(offset(limit, page))
	}
	if name == "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	if len(ids) > 0 {
		db = db.Where("id in ?", ids)
	}
	db.Find(&list)
	return list
}

func (Permission) Info(req *pb.DefaultPkReq) model.RbacPermission {
	var info model.RbacPermission
	DB.Where("id=?", req.Pk).First(&info)
	return info
}

func (Permission) Create(req *pb.PermissionCreateReq) errors.Error {
	var info model.RbacPermission
	first := DB.Where("name=?", req.Name).First(&info)
	if first.RowsAffected != 0 {
		return errors.NewDBDuplication("name")
	}
	err := DB.Create(&model.RbacPermission{
		DomainID: uint(req.DomainID),
		Name:     req.Name,
		Desc:     req.Desc,
	}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternalErr(err)
	}
	return nil
}

func (Permission) Update(p *pb.PermissionUpdateReq) errors.Error {
	var info model.RbacPermission
	first := DB.Where("name=? and id!=?", p.Name, p.ID).First(&info)
	if first.RowsAffected != 0 {
		return errors.NewDBNoData("")
	}
	snake := helper.Struct2MapSnakeNoZero(&p)
	delete(snake, "id")
	err := DB.Model(&model.RbacPermission{}).Updates(snake).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternalErr(err)
	}
	return nil
}

func (Permission) Del(req *pb.PermissionDelReq) errors.Error {
	err := DB.Where("id=? and do_main=?", req.ID, req.DomainID).Delete(&model.RbacPermission{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternalErr(err)
	}
	return nil
}

func (Permission) PathList(req *pb.PermissionPathListReq) []model.RbacPath {
	var list []model.RbacPermissionPath
	DB.Where("permission_id=?", req.PermissionID).Find(&list)
	uintSet := set.NewUintSet()
	for _, item := range list {
		uintSet.Add(item.PathID)
	}
	return Path{}.list("", -1, 0, 0, uintSet.List()...)
}

func (Permission) PathSet(req *pb.PermissionPathSetReq) errors.Error {
	var list []model.RbacPermissionPath
	DB.Where("permission_id=? and path_id in ?", req.PermissionID, req.PathIDs).Find(&list)
	uintSet := set.NewUintSet()
	for _, path := range list {
		uintSet.Add(path.PathID)
	}
	DB.Where("permission_id=? and path_id not in ?", req.PermissionID, req.PathIDs).Delete(&model.RbacPermissionPath{})
	pathIds := uintSet.GetAddItems(req.PathIDs)
	var newList = make([]model.RbacPermissionPath, 0, len(pathIds))
	for _, id := range pathIds {
		newList = append(newList, model.RbacPermissionPath{
			PermissionID: uint(req.PermissionID),
			PathID:       id,
		})
	}
	err := DB.Create(&newList).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternalErr(err)
	}
	return nil
}
