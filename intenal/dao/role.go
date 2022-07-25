package dao

import (
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
}

func (r Role) List(req *pb.RoleListReq) (res entity.RoleListRes) {
	db := DB.Where("domain_id=?", req.DomainID)
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

func (Role) list(db *gorm.DB, domainId int64, name string, ids []uint) []model.RbacRole {
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
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil

}

func (Role) Info(req *pb.DefaultPkReq) model.RbacRole {
	var info model.RbacRole
	DB.Where("id=? and domain_id=?", req.Pk.(*pb.DefaultPkReq_ID).ID, req.DomainID).First(&info)
	return info
}

func (Role) Update(req *pb.RoleUpdateReq) errors.Error {
	var info model.RbacRole
	first := DB.Where("domain_id=? and id!=? and name=?", req.DomainID, req.ID, req.Name).First(&info)
	if first.RowsAffected > 0 {
		return errors.NewDBDuplication("account")
	}
	snake := helper.Struct2MapSnakeNoZero(req)
	delete(snake, "id")
	err := DB.Where("id=? and domain_id=?", req.ID, req.DomainID).Model(&model.RbacRole{}).Updates(snake).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}
func (Role) Del(id, domainId int64) errors.Error {

	err := DB.Where("id=? and domain_id=?", id, domainId).Delete(&model.RbacRole{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	err = DB.Where("role_id=? and domain_id=?", id, domainId).Delete(&model.RbacRoleUser{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}

func (Role) RolePermissionList(req *pb.RolePermissionListReq) []model.RbacPermission {
	var list []model.RbacRolePermission

	DB.Where("domain_id=? and role_id=?", req.DomainID, req.RoleID).Find(&list)
	uintSet := set.NewSet[uint]()
	for _, permission := range list {
		uintSet.Add(permission.PermissionID)
	}
	if uintSet.Len() == 0 {
		return nil
	}
	return Permission{}.list(DB, -1, 0, uintSet.List()...)

}
func (Role) rolePermission(domainId int64, roleId []uint) []model.RbacRolePermission {
	var list []model.RbacRolePermission
	DB.Where("domain_id=? and role_id in ?", domainId, roleId).Find(&list)
	return list
}

func (Role) RolePermissionSet(req *pb.RolePermissionSetReq) errors.Error {
	var err error
	if len(req.PermissionIDS) == 0 {
		err = DB.Where("role_id=?", req.RoleID).Delete(&model.RbacRolePermission{}).Error
	} else {
		err = DB.Where("role_id=? and permission_id not in ?", req.RoleID, req.PermissionIDS).Delete(&model.RbacRolePermission{}).Error
	}
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	var list []model.RbacRolePermission
	DB.Where("role_id=? and domain_id=?", req.RoleID, req.DomainID).Find(&list)
	uintSet := set.NewSet[uint]()
	for _, permission := range list {
		uintSet.Add(permission.PermissionID)
	}
	addItems := uintSet.GetNewitems(helper.TransSliceType[int64, uint](req.PermissionIDS))
	if len(addItems) == 0 {
		return nil
	}
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
		return errors.NewDBInternal(err)
	}
	return nil
}

func (Role) getRoleIdsByUID(uid uint) []uint {
	var list []model.RbacRoleUser

	DB.Where("uid=?", uid).Find(&list)
	uintSet := set.NewSet[uint]()
	for _, user := range list {
		uintSet.Add(user.RoleID)
	}
	return uintSet.List()
}

func (r Role) GetDomainPolicy(domainId int64) []entity.DomainPolicy {
	roleList := r.list(DB, domainId, "", nil)
	rset := set.NewSet[uint]()
	for _, v := range roleList {
		rset.Add(v.ID)
	}
	var list []entity.DomainPolicy
	var pSet = set.NewSet[uint]()
	rolePermissionList := r.rolePermission(domainId, rset.List())
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
	permissionPath := Permission{}.PermissionPathMap(pSet.List())
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

	pathMap := Path{}.pathMap(DB, []string{"id", "api_path", "method"}, pathSet.List()...)

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
