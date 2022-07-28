package dao

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/liujunren93/share_rbac/intenal/entity"
	"github.com/liujunren93/share_rbac/intenal/model"
	"github.com/liujunren93/share_rbac/log"
	pb "github.com/liujunren93/share_rbac/rbac_pb"
	"github.com/liujunren93/share_utils/common/set"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/helper"
	"gorm.io/gorm"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 15:47
 */

type Path struct {
	dao
}

func NewPath(ctx context.Context) Path {
	return Path{dao{ctx}}
}

func (dao Path) List(req *pb.PathListReq) entity.PathListRes {
	var res entity.PathListRes
	db := DB(dao.Ctx)
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if req.PathType != 0 {
		db = db.Where("path_type = ?", req.PathType)
	}
	if req.Path != "" {
		db = db.Where("path like ?", "%"+req.Path+"%")
	}
	db = db.Where("domain_id = ? or domain_id=-1", req.DomainID)

	if req.PageSize > 0 {
		res.Total = dao.count(db)
		db = db.Limit(pageSize(req.PageSize)).Offset(offset(req.PageSize, req.Page))
	}
	if req.SortField != "" {
		db = db.Order(req.SortField + " " + req.SortOrder)
	}

	res.List = dao.list(db, nil)

	return res
}
func (dao Path) count(db *gorm.DB, ids ...uint) int64 {

	if len(ids) != 0 {
		db = db.Where("id in ?", ids)
	}
	var total int64
	db.Model(&model.RbacPath{}).Count(&total)
	return total
}

func (dao Path) list(db *gorm.DB, selectField []string, ids ...uint) []model.RbacPath {
	var list []model.RbacPath
	if len(ids) != 0 {
		db = db.Where("id in ?", ids)
	}
	if len(selectField) > 0 {
		db.Select(strings.Join(selectField, ","))
	}
	db.Find(&list)
	return list
}
func (dao Path) pathMap(db *gorm.DB, selectField []string, ids ...uint) map[uint]model.RbacPath {
	pathlist := dao.list(db, selectField, ids...)
	var pathM = make(map[uint]model.RbacPath, len(pathlist))
	for _, p := range pathlist {
		pathM[p.ID] = p
	}
	return pathM
}

func (dao Path) Create(req *pb.PathCreateReq) (uint, errors.Error) {
	first := DB(dao.Ctx).Where("name=?", req.Name).First(&model.RbacPath{})
	if first.RowsAffected > 0 {
		return 0, errors.NewDBDuplication("name")
	}
	path := model.RbacPath{
		Key:       req.Key,
		Name:      req.Name,
		Component: req.Component,
		Redirect:  req.Redirect,
		ParentID:  uint(req.ParentID),
		Path:      req.Path,
		PathType:  int8(req.PathType),
		DomainID:  int(req.DomainID),
		ApiPath:   req.ApiPath,
		Method:    req.Method,
		Action:    req.Action,
	}
	if req.Meta != nil {
		marshal, err := json.Marshal(req.Meta)
		if err != nil {
			log.Logger.Error(err)
			return 0, errors.NewInternalError(err)
		}
		path.Meta = string(marshal)
	}
	err := DB(dao.Ctx).Create(&path).Error
	if err != nil {
		return 0, errors.NewDBInternal(err)
	}
	return path.ID, nil

}

func (dao Path) Info(req *pb.DefaultPkReq) model.RbacPath {
	var info model.RbacPath
	DB(dao.Ctx).Where("id=?", req.Pk.(*pb.DefaultPkReq_ID).ID).First(&info)
	return info
}

func (dao Path) Update(req *pb.PathUpdateReq) errors.Error {
	var info model.RbacAdmin
	first := DB(dao.Ctx).Where("id!=? and name=?", req.ID, req.Name).First(&info)
	if first.RowsAffected > 0 {
		return errors.NewDBDuplication("account")
	}
	snake := helper.Struct2MapSnake(req)
	delete(snake, "id")
	if req.Meta != nil {
		b, _ := json.Marshal(&req.Meta)
		snake["meta"] = string(b)
	} else {
		delete(snake, "meta")
	}

	err := DB(dao.Ctx).Where("id=? ", req.ID).Model(&model.RbacPath{}).Updates(snake).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}
func (dao Path) Del(req *pb.DefaultPkReq) errors.Error {
	err := DB(dao.Ctx).Where("id=?", req.Pk.(*pb.DefaultPkReq_ID).ID).Delete(&model.RbacPath{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return NewPermission(dao.Ctx).pathDel(uint(req.Pk.(*pb.DefaultPkReq_ID).ID))
}

func (dao Path) DelByids(ids []uint) errors.Error {
	err := DB(dao.Ctx).Where("id in ?", ids).Delete(&model.RbacPath{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}

// func (dao Path) getPermissionByIDs(roleIDs ...uint) {
// 	var list []model.RbacPermissionPath
// 	DB(dao.Ctx).Where("role_id=?", roleIDs).Find(&list)

// }

//GetPathByUid 获取用户path
func (dao Path) GetRolePath(pathType int8, roleIDs []int64) []model.RbacPath {
	pids := NewPermission(dao.Ctx).getPermissionIDsByRoleIDs(roleIDs)
	return dao.getPathByPIDs(pathType, pids...)

}

//GetPathByUid 获取用户path
func (dao Path) GetPathWithPermissionByUid(pathType int8, uid uint) []entity.UserPermission {
	var resList []entity.UserPermission
	role := Role{}
	roles := role.getRoleIdsByUID(uid)
	r := helper.TransSliceType[uint, int64](roles)
	daoPermission := NewPermission(dao.Ctx)
	pids := daoPermission.getPermissionIDsByRoleIDs(r)
	permissionPathList := daoPermission.PermissionPathList(pids)
	pathSet := set.NewSet[uint]()
	permissionSet := set.NewSet[uint]()
	for _, v := range permissionPathList {
		resList = append(resList, entity.UserPermission{
			PID:        v.PermissionID,
			PathID:     v.PathID,
			ActionList: []string{},
		})
		permissionSet.Add(v.PermissionID)
		pathSet.Add(v.PathID)
	}
	pmap := daoPermission.permissionMap(DB(dao.Ctx), -1, 0, permissionSet.List()...)
	pathList := NewPath(dao.Ctx).pathMap(DB(dao.Ctx), []string{"id", "name"}, pathSet.List()...)
	for i, v := range resList {
		resList[i].Path = pathList[v.PathID].Name
		if v, ok := pmap[v.PID]; ok {
			resList[i].ActionList = append(resList[i].ActionList, v.Name)
		}
	}
	return resList
}

//GetPathByUid 获取用户path
func (dao Path) GetPathActionsByRoles(pathType int8, roles []int64) map[string][]string {
	var res = make(map[string][]string)
	daoPermission := NewPermission(dao.Ctx)
	pids := daoPermission.getPermissionIDsByRoleIDs(roles)
	permissionPathList := daoPermission.PermissionPathList(pids)
	pathSet := set.NewSet[uint]()
	// permissionSet := set.NewSet[uint]()
	for _, v := range permissionPathList {
		// tmp[v.PathID] = append(tmp[v.PathID], v.PermissionID)
		// permissionSet.Add(v.PermissionID)
		pathSet.Add(v.PathID)
	}
	// pmap := Permission(dao.Ctx)permissionMap(DB, -1, 0, permissionSet.List()...)
	pathList := NewPath(dao.Ctx).list(DB(dao.Ctx), []string{"id", "method", "`key`", "api_path", "action"}, pathSet.List()...)
	for _, v := range pathList {
		if v.Key == "" || v.Action == "" {
			continue
		}
		res[v.Key] = append(res[v.Key], v.Action)
	}
	return res
}

func (dao Path) getPathByPIDs(pathType int8, permissionIDs ...uint) []model.RbacPath {
	var list []model.RbacPermissionPath
	DB(dao.Ctx).Where("permission_id in ?", permissionIDs).Find(&list)
	if len(list) == 0 {
		return nil
	}
	uintSet := set.NewSet[uint]()
	for _, v := range list {
		uintSet.Add(v.PathID)
	}
	db := DB(dao.Ctx).Where("path_type=?", pathType).Order("sort asc").Debug()

	return dao.list(db, nil, uintSet.List()...)

}
