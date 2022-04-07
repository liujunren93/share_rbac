package dao

import (
	"encoding/json"
	"github.com/liujunren93/share_rbac/intenal/entity"
	"github.com/liujunren93/share_rbac/intenal/model"
	"github.com/liujunren93/share_rbac/log"
	"github.com/liujunren93/share_rbac/pb"
	"github.com/liujunren93/share_utils/commond/set"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/helper"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 15:47
 */

type Path struct {
}

func (dao Path) List(req *pb.PathListReq) entity.PathListRes {
	var res entity.PathListRes
	res.List = dao.list(req.Title, req.PageSize, req.Page, 0)
	res.Total = dao.Count(req.Title)
	return res
}
func (Path) Count(title string, ids ...uint) int64 {
	db := DB
	if title == "" {
		db = db.Where("title like ?", "%"+title+"%")
	}
	if len(ids) == 0 {
		db = db.Where("id in ?", ids)
	}
	var total int64
	db.Count(&total)
	return total
}

func (Path) list(title string, limit, page int64, pathType int8, ids ...uint) []model.RbacPath {
	var list []model.RbacPath
	db := DB
	if title == "" {
		db = db.Where("title like ?", "%"+title+"%")
	}
	if len(ids) == 0 {
		db = db.Where("id in ?", ids)
	}
	if limit >= 0 {
		db = db.Limit(pageSize(limit)).Offset(offset(limit, page))
	}
	if pathType != 0 {
		db = db.Where("path_type = ? ", pathType)
	}
	db.Find(&list)
	return list
}

func (Path) Create(req *pb.PathCreateReq) errors.Error {
	first := DB.Where("name=?", req.Name).First(&model.RbacPath{})
	if first.RowsAffected > 0 {
		return errors.NewDBDuplication("name")
	}
	path := model.RbacPath{
		Key:       helper.RandString(32),
		Name:      req.Name,
		Component: req.Component,
		Redirect:  req.Redirect,
		ParentID:  uint(req.ParentID),
		Path:      req.Path,
		PathType:  int8(req.PathType),
	}
	if req.Meta != nil {
		marshal, err := json.Marshal(req.Meta)
		if err != nil {
			log.Logger.Error(err)
			return errors.InternalErrorMsg(err)
		}
		path.Meta = string(marshal)
	}
	err := DB.Create(&path).Error
	if err != nil {
		return errors.NewDBInternalErr(err)
	}
	return nil

}

func (Path) Info(req *pb.DefaultPkReq) model.RbacPath {
	var info model.RbacPath
	DB.Where("id=?", req.Pk).First(&info)
	return info
}

func (Path) Update(req *pb.PathUpdateReq) errors.Error {
	var info model.RbacAdmin
	first := DB.Where("id!=? and name=?", req.ID, req.Name).First(&info)
	if first.RowsAffected > 0 {
		return errors.NewDBDuplication("account")
	}
	snake := helper.Struct2MapSnakeNoZero(req)
	delete(snake, "id")
	err := DB.Where("id=? ", req.ID).Updates(snake).Error
	if err != nil {
		return errors.NewDBInternalErr(err)
	}
	return nil
}
func (Path) Del(req *pb.DefaultPkReq) errors.Error {
	err := DB.Where("id=?", req.Pk).Delete(&model.RbacPath{}).Error
	if err != nil {
		return errors.NewDBInternalErr(err)
	}
	return nil
}

func (Path) getPermissionByIDs(roleIDs ...uint) {
	var list []model.RbacPermissionPath
	DB.Where("role_id=?", roleIDs).Find(&list)

}
func (dao Path) GetPathByUid(pathType int8, uid uint) []model.RbacPath {
	role := Role{}
	roles := role.getRoleIdsByUID(uid)
	pids := role.getPermissionIDsByRoleIDs(roles)
	return dao.getPathByPIDs(pathType, pids...)

}

func (dao Path) getPathByPIDs(pathType int8, permissionIDs ...uint) []model.RbacPath {
	var list []model.RbacPermissionPath
	DB.Where("permission_id in ?", permissionIDs).Find(&list)
	if len(list) == 0 {
		return nil
	}
	uintSet := set.NewUintSet()
	for _, v := range list {
		uintSet.Add(v.PathID)
	}
	return dao.list("", -1, 0, pathType, uintSet.List()...)

}
