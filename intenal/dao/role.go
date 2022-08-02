package dao

import (
	"context"

	"github.com/liujunren93/share_rbac/intenal/entity"
	"github.com/liujunren93/share_rbac/intenal/model"
	"github.com/liujunren93/share_rbac/log"
	pb "github.com/liujunren93/share_rbac/rbac_pb"
	"github.com/liujunren93/share_utils/common/set"
	"gorm.io/gorm"

	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/helper"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 15:13
 */

type Role struct {
	dao
}

func NewRole(ctx context.Context) Role {
	return Role{dao: dao{ctx}}
}

func (dao Role) List(req *pb.RoleListReq) (res entity.RoleListRes) {
	db := DB(dao.Ctx).Where("domain_id=?", dao.GetSession().DomainID)
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	db.Model(&model.RbacRole{}).Count(&res.Total)
	if req.SortField != "" {
		db = db.Order(req.SortField + " " + req.SortOrder)
	}
	db = db.Limit(int(req.PageSize)).Offset(int(req.Page * req.PageSize)).Find(&res.List)
	return
}

func (dao Role) list(db *gorm.DB, domainId int64, name string, ids []uint) []model.RbacRole {
	var list []model.RbacRole
	db = db.Where("domain_id=?", domainId)
	if ids != nil {
		db = db.Where("id in ?", ids)
	}
	if name != "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	db.Find(&list)
	return list
}

func (dao Role) Create(req *pb.RoleCreateReq) errors.Error {
	role := model.RbacRole{
		DomainID: uint(dao.GetSession().DomainID),
		Name:     req.Name,
		Desc:     req.Desc,
	}
	err := dao.create(DB(dao.Ctx), &role)
	if err != nil {
		return err
	}
	return nil

}

func (dao Role) create(tx *gorm.DB, role *model.RbacRole) errors.Error {
	first := DB(dao.Ctx).Where("domain_id=? and  name=?", role.DomainID, role.Name).First(&model.RbacRole{})
	if first.RowsAffected > 0 {
		return errors.NewDBDuplication("角色名已存在")
	}
	err := DB(dao.Ctx).Create(role).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil

}

func (dao Role) Info(req *pb.DefaultPkReq) model.RbacRole {
	var info model.RbacRole
	DB(dao.Ctx).Where("id=? and domain_id=?", req.Pk.(*pb.DefaultPkReq_ID).ID, dao.GetSession().DomainID).First(&info)
	return info
}

func (dao Role) Update(req *pb.RoleUpdateReq) errors.Error {
	var info model.RbacRole
	first := DB(dao.Ctx).Where("domain_id=? and id!=? and name=?", dao.GetSession().DomainID, req.ID, req.Name).First(&info)
	if first.RowsAffected > 0 {
		return errors.NewDBDuplication("account")
	}
	snake := helper.Struct2MapSnakeNoZero(req)
	if req.IsLock {
		snake["pl"] = dao.NewPL()
	}
	err := DB(dao.Ctx).Where("id=? and domain_id=?", req.ID, dao.GetSession().DomainID).Model(&model.RbacRole{}).Updates(snake).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}
func (dao Role) Del(id, domainId int64) errors.Error {

	err := DB(dao.Ctx).Where("id=? and domain_id=?", id, domainId).Delete(&model.RbacRole{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	err = DB(dao.Ctx).Where("role_id=? and domain_id=?", id, domainId).Delete(&model.RbacRoleUser{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}

func (dao Role) RolePermissionList(req *pb.RolePermissionListReq) []model.RbacPermission {
	var list []model.RbacRolePermission

	DB(dao.Ctx).Where("domain_id=? and role_id=?", dao.GetSession().DomainID, req.RoleID).Find(&list)
	uintSet := set.NewSet[uint]()
	for _, permission := range list {
		uintSet.Add(permission.PermissionID)
	}
	if uintSet.Len() == 0 {
		return nil
	}
	return NewPermission(dao.Ctx).list(DB(dao.Ctx), -1, 0, uintSet.List()...)

}
func (dao Role) rolePermission(domainId int64, roleId []uint) []model.RbacRolePermission {
	var list []model.RbacRolePermission
	DB(dao.Ctx).Where("domain_id=? and role_id in ?", domainId, roleId).Find(&list)
	return list
}

func (dao Role) RolePermissionSet(req *pb.RolePermissionSetReq) errors.Error {
	err := dao.rolePermissionSet(DB(dao.Ctx), uint(req.RoleID), uint(dao.GetSession().DomainID), helper.TransSliceType[int64, uint](req.PermissionIDS)...)
	if err != nil {
		log.Logger.Error("Role.RolePermissionSet", err)
	}
	return err
}

func (dao Role) rolePermissionSet(tx *gorm.DB, roleId, domainId uint, permissionIDS ...uint) errors.Error {
	log.Logger.Debug("rolePermissionSet")
	var err error
	if len(permissionIDS) == 0 {
		err = tx.Where("role_id=?", roleId).Delete(&model.RbacRolePermission{}).Error
	} else {
		err = tx.Where("role_id=? and permission_id not in ?", roleId, permissionIDS).Delete(&model.RbacRolePermission{}).Error
	}
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	var list []model.RbacRolePermission
	tx.Where("role_id=? and domain_id=?", roleId, domainId).Find(&list)
	uintSet := set.NewSet[uint]()
	for _, permission := range list {
		uintSet.Add(permission.PermissionID)
	}
	addItems := uintSet.GetNewitems(permissionIDS)
	log.Logger.Debug("rolePermissionSet.newItem", addItems)
	if len(addItems) == 0 {
		return nil
	}

	var newData []model.RbacRolePermission
	for _, id := range addItems {
		newData = append(newData, model.RbacRolePermission{
			RoleID:       roleId,
			PermissionID: id,
			DomainID:     domainId,
		})
	}

	err = tx.Create(&newData).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}

func (dao Role) getRoleIdsByUID(uid uint) []uint {
	var list []model.RbacRoleUser

	DB(dao.Ctx).Where("uid=?", uid).Find(&list)
	uintSet := set.NewSet[uint]()
	for _, user := range list {
		uintSet.Add(user.RoleID)
	}
	return uintSet.List()
}

func (dao Role) GetDomainPolicy() []entity.DomainPolicy {
	roleList := dao.list(DB(dao.Ctx), dao.GetSession().DomainID, "", nil)
	rset := set.NewSet[uint]()
	for _, v := range roleList {
		rset.Add(v.ID)
	}
	var list []entity.DomainPolicy
	var pSet = set.NewSet[uint]()
	rolePermissionList := dao.rolePermission(dao.GetSession().DomainID, rset.List())
	for _, v := range rolePermissionList {
		pSet.Add(v.PermissionID)
		list = append(list, entity.DomainPolicy{
			RoleID:       v.RoleID,
			PermissionID: v.PermissionID,
			ApiPath:      "",
			Method:       "",
		})
	}
	var pathSet = set.NewSet[uint]()
	permissionPath := NewPermission(dao.Ctx).PermissionPathMap(pSet.List())
	for _, paths := range permissionPath {
		pathSet.Add(paths...)
	}
	if len(pathSet.List()) == 0 {
		return nil
	}
	var resList []entity.DomainPolicy
	for _, v := range list {
		if pids, ok := permissionPath[v.PermissionID]; ok {
			for _, pid := range pids {
				resList = append(resList, entity.DomainPolicy{
					RoleID:       v.RoleID,
					PermissionID: v.PermissionID,
					PathID:       pid,
				})
			}

		}

	}

	pathMap := NewPath(dao.Ctx).pathMap(DB(dao.Ctx), []string{"id", "api_path", "method"}, pathSet.List()...)
	j := 0
	for _, v := range resList {
		if p, ok := pathMap[v.PathID]; ok {
			if p.ApiPath == "" || p.Method == "" {
				continue
			}
			v.ApiPath = p.ApiPath
			v.Method = p.Method
			resList[j] = v
			j++

		}
	}
	return resList[:j]
}
