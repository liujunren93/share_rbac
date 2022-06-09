package dao

import (
	"github.com/liujunren93/share_rbac/intenal/entity"
	"github.com/liujunren93/share_rbac/intenal/model"
	"github.com/liujunren93/share_rbac/log"
	"github.com/liujunren93/share_rbac/pb"
	"github.com/liujunren93/share_utils/common/set"
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

const REDIS_ALL_PERMISSION_KEY = "share_rbac_all_permission"

func InitPerrmission() {

}

func (dao Permission) List(req *pb.PermissionListReq) entity.PermissionListRes {
	var res entity.PermissionListRes
	db := DB.Where("domain_id=-1 or domain_id=?", req.DomainID)
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}
	res.List = dao.list(db, req.PageSize, req.Page)
	res.Total = dao.count(db)

	return res
}

func (Permission) count(db *gorm.DB, ids ...uint) int64 {

	if len(ids) > 0 {
		db = db.Where("id in ?", ids)
	}
	var total int64
	db.Model(&model.RbacPermission{}).Count(&total)
	return total
}

func (Permission) list(db *gorm.DB, limit, page int64, ids ...uint) []model.RbacPermission {
	var list []model.RbacPermission

	if limit >= 0 {
		db = db.Limit(pageSize(limit)).Offset(offset(limit, page))
	}
	if len(ids) > 0 {
		db = db.Where("id in ?", ids)
	}
	db.Find(&list)
	return list
}

func (dao Permission) permissionMap(db *gorm.DB, limit, page int64, ids ...uint) map[uint]model.RbacPermission {
	plist := dao.list(db, limit, page, ids...)
	var pmap = make(map[uint]model.RbacPermission, len(plist))
	for _, v := range plist {
		pmap[v.ID] = v
	}
	return pmap
}

func (Permission) Info(req *pb.DefaultPkReq) model.RbacPermission {
	var info model.RbacPermission
	DB.Where("id=? and domain_id=?", req.Pk.(*pb.DefaultPkReq_ID).ID, req.DomainID).First(&info)
	return info
}

func (Permission) Create(req *pb.PermissionCreateReq) (uint, errors.Error) {
	var info model.RbacPermission
	first := DB.Where("name=?", req.Name).First(&info)
	if first.RowsAffected != 0 {
		return 0, errors.NewDBDuplication("name")
	}
	permission := model.RbacPermission{
		DomainID: int(req.DomainID),
		Name:     req.Name,
		Desc:     req.Desc,
	}
	err := DB.Create(&permission).Error
	if err != nil {
		log.Logger.Error(err)
		return 0, errors.NewDBInternal(err)
	}
	return permission.ID, nil
}

func (Permission) Update(req *pb.PermissionUpdateReq) errors.Error {
	var info model.RbacPermission
	first := DB.Where("domain_id=? and name=? and id!=?", req.DomainID, req.Name, req.ID).First(&info)
	if first.RowsAffected != 0 {
		return errors.NewDBNoData("")
	}
	snake := helper.Struct2MapSnakeNoZero(req)
	delete(snake, "id")
	err := DB.Model(&model.RbacPermission{}).Where("id=? and domain_id=?", req.ID, req.DomainID).Updates(snake).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}

func (Permission) Del(req *pb.DefaultPkReq) errors.Error {
	err := DB.Where("id=? and domain_id=?", req.Pk.(*pb.DefaultPkReq_ID).ID, req.DomainID).Delete(&model.RbacPermission{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}

func (Permission) PathList(req *pb.PermissionPathListReq) []model.RbacPath {
	var list []model.RbacPermissionPath
	DB.Where("permission_id=?", req.PermissionID).Find(&list)
	uintSet := set.NewSet[uint]()
	for _, item := range list {
		uintSet.Add(item.PathID)
	}

	if uintSet.Len() == 0 {
		return nil
	}
	return Path{}.list(DB, []string{"id,title,name,path"}, uintSet.List()...)
}

func (Permission) PathSet(req *pb.PermissionPathSetReq) errors.Error {
	var err error
	if len(req.PathIDs) == 0 {
		err = DB.Where("permission_id=?", req.PermissionID, req.PathIDs).Delete(&model.RbacPermissionPath{}).Error
	} else {
		err = DB.Where("permission_id=? and path_id not in ?", req.PermissionID, req.PathIDs).Delete(&model.RbacPermissionPath{}).Error
	}
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	var list []model.RbacPermissionPath
	DB.Where("permission_id=? and path_id in ?", req.PermissionID, req.PathIDs).Find(&list)
	uintSet := set.NewSet[uint]()
	for _, path := range list {
		uintSet.Add(path.PathID)
	}
	pathIds := uintSet.GetNewitems(helper.TransSliceType[int64, uint](req.PathIDs))

	if len(pathIds) == 0 {
		return nil
	}
	var newList = make([]model.RbacPermissionPath, 0, len(pathIds))
	for _, id := range pathIds {
		newList = append(newList, model.RbacPermissionPath{
			PermissionID: uint(req.PermissionID),
			PathID:       id,
		})
	}
	err = DB.Create(&newList).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}

func (Permission) getPermissionIDsByRoleIDs(roles []int64) []uint {
	var rolePermissionList []model.RbacRolePermission
	DB.Where("role_id in ?", roles).Find(&rolePermissionList)
	uintSet := set.NewSet[uint]()
	for _, permission := range rolePermissionList {
		uintSet.Add(permission.PermissionID)
	}
	return uintSet.List()
}

func (dao Permission) PermissionPathList(pids []uint) []model.RbacPermissionPath {
	var list []model.RbacPermissionPath
	DB.Where("permission_id in ?", pids).Find(&list)
	return list
}

func (dao Permission) PermissionPathMap(pids []uint) map[uint][]uint {
	var ppMap = make(map[uint][]uint)
	rpp := dao.PermissionPathList(pids)
	for _, v := range rpp {
		ppMap[v.PermissionID] = append(ppMap[v.PermissionID], v.PathID)
	}
	return ppMap
}
