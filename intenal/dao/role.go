package dao

import (
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
* @Date: 2022/2/28 15:13
 */

type Role struct {
}

func (Role) List(req *pb.RoleListReq) (res entity.AdminListRes) {
	db := DB.Where("domain_id=?", req.DomainID)
	if req.Name == "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	db.Model(&model.RbacAdmin{}).Count(&res.Total)
	DB.Limit(int(req.PageSize)).Offset(int(req.Page * req.PageSize)).Find(&res.List)
	return
}

func (Role) Create(req *pb.RoleCreateReq) errors.Error {
	first := DB.Where("name=?", req.Name).First(&model.RbacRole{})
	if first.RowsAffected > 0 {
		return errors.NewDBDuplication("name")
	}
	err := DB.Create(&model.RbacRole{
		DomainID: uint(req.DomainID),
		Name:     req.Name,
		Desc:     req.Desc,
	}).Error
	if err != nil {
		return errors.NewDBInternalErr(err)
	}
	return nil

}

func (Role) Info(Id *pb.DefaultPkReq) model.RbacRole {
	var info model.RbacRole
	DB.Where("id=?", Id).First(&info)
	return info
}

func (Role) Update(req *pb.RoleUpdateReq) errors.Error {
	var info model.RbacAdmin
	first := DB.Where("id!=? and name=?", req.ID, req.Name).First(&info)
	if first.RowsAffected > 0 {
		return errors.NewDBDuplication("account")
	}
	snake := helper.Struct2MapSnakeNoZero(req)
	delete(snake, "id")
	err := DB.Where("id=?", req.ID).Updates(snake).Error
	if err != nil {
		return errors.NewDBInternalErr(err)
	}
	return nil
}
func (Role) Del(id int64) errors.Error {

	err := DB.Where("id=?", id).Delete(&model.RbacRole{}).Error
	if err != nil {
		return errors.NewDBInternalErr(err)
	}
	return nil
}

func (Role) RolePermissionList(req *pb.RolePermissionListReq) []model.RbacPermission {
	var list []model.RbacRolePermission

	DB.Where("domain_id=? and role_id=?", req.DomainID, req.RoleID).Find(&list)
	uintSet := set.NewUintSet()
	for _, permission := range list {
		uintSet.Add(permission.PermissionID)
	}
	return Permission{}.list(DB, "", -1, 0, uintSet.List()...)

}

func (Role) RolePermissionSet(req *pb.RolePermissionSetReq) errors.Error {
	var list []model.RbacRolePermission
	DB.Where("role_id=?", req.RoleID).Find(&list)
	uintSet := set.NewUintSet()
	for _, permission := range list {
		uintSet.Add(permission.PermissionID)
	}
	err := DB.Where("role_id=? and permission_id not in ?", req.RoleID, req.PermissionIDS).Model(&model.RbacRolePermission{}).UpdateColumn("status", 2).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternalErr(err)
	}
	err = DB.Where("role_id=? and permission_id  in ?", req.RoleID, req.PermissionIDS).Model(&model.RbacRolePermission{}).UpdateColumn("status", 1).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternalErr(err)
	}
	addItems := uintSet.GetAddItems(req.PermissionIDS)
	var newData []model.RbacRolePermission
	for _, id := range addItems {
		newData = append(newData, model.RbacRolePermission{
			RoleID:       uint(req.RoleID),
			PermissionID: id,
			DomainID:     uint(req.DomainID),
		})
	}
	err = DB.Create(&newData).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternalErr(err)
	}
	return nil
}

func (Role) getRoleIdsByUID(uid uint) []uint {
	var list []model.RbacRoleUser

	DB.Where("uid=?", uid).Find(&list)
	uintSet := set.NewUintSet()
	for _, user := range list {
		uintSet.Add(user.RoleID)
	}
	return uintSet.List()
}
func (Role) getPermissionIDsByRoleIDs(roles []uint) []uint {
	var rolePermissionList []model.RbacRolePermission
	DB.Where("role_id in ?", roles).Find(&rolePermissionList)
	uintSet := set.NewUintSet()
	for _, permission := range rolePermissionList {
		uintSet.Add(permission.PermissionID)
	}
	return uintSet.List()
}
