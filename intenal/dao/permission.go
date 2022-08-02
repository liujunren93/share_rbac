package dao

import (
	"context"

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
* @Date: 2022/2/28 11:45
 */
type Permission struct {
	dao
}

const REDIS_ALL_PERMISSION_KEY = "share_rbac_all_permission"

func NewPermission(ctx context.Context) Permission {
	return Permission{dao: dao{ctx}}
}

func (dao Permission) List(req *pb.PermissionListReq) entity.PermissionListRes {
	var res entity.PermissionListRes
	db := DB(dao.Ctx).Where("domain_id=-1 or domain_id=?", dao.GetSession().DomainID)
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}
	res.Total = dao.count(db)
	if req.SortField != "" {
		db = db.Order(req.SortField + " " + req.SortOrder)
	}
	res.List = dao.list(db, req.PageSize, req.Page)

	return res
}

func (dao Permission) count(db *gorm.DB, ids ...uint) int64 {

	if len(ids) > 0 {
		db = db.Where("id in ?", ids)
	}
	var total int64
	db.Model(&model.RbacPermission{}).Count(&total)
	return total
}

func (dao Permission) list(db *gorm.DB, limit, page int64, ids ...uint) []model.RbacPermission {
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

func (dao Permission) Info(req *pb.DefaultPkReq) model.RbacPermission {
	var info model.RbacPermission
	DB(dao.Ctx).Where("id=? and domain_id=?", req.Pk.(*pb.DefaultPkReq_ID).ID, dao.GetSession().DomainID).First(&info)
	return info
}

func (dao Permission) Create(req *pb.PermissionCreateReq) (uint, errors.Error) {
	var info model.RbacPermission
	session := dao.GetSession()
	domainId := session.DomainID
	if req.IsPublic && session.DomainID == 1 { // 只有平台才能设置公共权限
		domainId = -1
	}
	first := DB(dao.Ctx).Where("domain_id=? and name=?", domainId, req.Name).First(&info)
	if first.RowsAffected != 0 {
		return 0, errors.NewDBDuplication("权限名重复")
	}
	permission := model.RbacPermission{
		Name:     req.Name,
		Desc:     req.Desc,
		DomainID: int(domainId),
	}

	err := DB(dao.Ctx).Create(&permission).Error
	if err != nil {
		log.Logger.Error(err)
		return 0, errors.NewDBInternal(err)
	}
	return permission.ID, nil
}

func (dao Permission) Update(req *pb.PermissionUpdateReq) (int, errors.Error) {
	var info model.RbacPermission
	session := dao.GetSession()
	domainId := session.DomainID
	if req.IsPublic && session.DomainID == 1 { // 只有平台才能设置公共权限
		domainId = -1
	}
	first := DB(dao.Ctx).Where("domain_id=? and name=? and id!=?", domainId, req.Name, req.ID).First(&info)
	if first.RowsAffected != 0 {
		return 0, errors.NewDBNoData("")
	}
	snake := helper.Struct2MapSnakeNoZero(req)
	if req.IsLock {
		snake["pl"] = dao.NewPL()
	}

	snake["domain_id"] = domainId

	err := DB(dao.Ctx).Model(&model.RbacPermission{}).Where("id=? and domain_id=?", req.ID, domainId).Updates(snake).Error
	if err != nil {
		log.Logger.Error("Permission.Update.Updates", err)
		return 0, errors.NewDBInternal(err)
	}
	return int(domainId), nil
}

func (dao Permission) Del(req *pb.DefaultPkReq) (domainId int, err errors.Error) {
	var er error
	if dao.GetSession().PL == 1 {
		domainId = -1
		var info model.RbacPermission
		er = DB(dao.Ctx).Where("id=?", req.Pk.(*pb.DefaultPkReq_ID)).Delete(&info).Error
		domainId = info.DomainID
	} else {
		er = DB(dao.Ctx).Where("id=? and domain_id=?", req.Pk.(*pb.DefaultPkReq_ID).ID, dao.GetSession().DomainID).Delete(&model.RbacPermission{}).Error
		domainId = int(dao.GetSession().DomainID)
	}
	if er != nil {
		log.Logger.Error("Permission.Del", err, req)
		return 0, errors.NewDBInternal(err)
	}
	return domainId, nil
}

func (dao Permission) PathList(req *pb.PermissionPathListReq) []model.RbacPath {
	var list []model.RbacPermissionPath
	DB(dao.Ctx).Where("permission_id=?", req.PermissionID).Find(&list)
	uintSet := set.NewSet[uint]()
	for _, item := range list {
		uintSet.Add(item.PathID)
	}

	if uintSet.Len() == 0 {
		return nil
	}
	return NewPath(dao.Ctx).list(DB(dao.Ctx), []string{"id,title,name,path"}, uintSet.List()...)
}
func (dao Permission) pathDel(tx *gorm.DB, id uint) errors.Error {
	err := tx.Where("path_id=?", id).Delete(&model.RbacPermissionPath{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}
func (dao Permission) PathSet(req *pb.PermissionPathSetReq) errors.Error {
	var err error
	if len(req.PathIDs) == 0 {
		err = DB(dao.Ctx).Where("permission_id=?", req.PermissionID).Delete(&model.RbacPermissionPath{}).Error
	} else {
		err = DB(dao.Ctx).Where("permission_id=? and path_id not in ?", req.PermissionID, req.PathIDs).Delete(&model.RbacPermissionPath{}).Error
	}
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	var list []model.RbacPermissionPath
	DB(dao.Ctx).Where("permission_id=? and path_id in ?", req.PermissionID, req.PathIDs).Find(&list)
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
	err = DB(dao.Ctx).Create(&newList).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}

func (dao Permission) getPermissionIDsByRoleIDs(roles []int64) []uint {
	var rolePermissionList []model.RbacRolePermission
	DB(dao.Ctx).Where("role_id in ?", roles).Find(&rolePermissionList)
	uintSet := set.NewSet[uint]()
	for _, permission := range rolePermissionList {
		uintSet.Add(permission.PermissionID)
	}
	return uintSet.List()
}

func (dao Permission) PermissionPathList(pids []uint) []model.RbacPermissionPath {
	var list []model.RbacPermissionPath
	DB(dao.Ctx).Where("permission_id in ?", pids).Find(&list)
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

func initRootPermission(tx *gorm.DB, domainId uint) {

}
