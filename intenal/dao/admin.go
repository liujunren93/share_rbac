package dao

import (
	"context"
	"strings"

	"github.com/liujunren93/share_rbac/intenal/entity"
	"github.com/liujunren93/share_rbac/intenal/model"
	"github.com/liujunren93/share_rbac/log"
	pb "github.com/liujunren93/share_rbac/rbac_pb"
	"github.com/liujunren93/share_utils/common/list"
	"github.com/liujunren93/share_utils/common/set"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/helper"
	"gorm.io/gorm"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 15:13
 */

type Admin struct {
	dao
}

func NewAdmin(ctx context.Context) Admin {
	return Admin{
		dao: dao{ctx},
	}
}

func (dao Admin) List(req *pb.AdminListReq) (res entity.AdminListRes) {
	db := DB(dao.Ctx).Where("domain_id=?", req.DomainID)
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Account != "" {
		db = db.Where("account like ?", "%"+req.Account+"%")
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}

	db.Model(&model.RbacAdmin{}).Count(&res.Total)
	if req.SortField != "" {
		db = db.Order(req.SortField + " " + req.SortOrder)
	}
	db.Limit(pageSize(req.PageSize)).Offset(offset(req.PageSize, req.Page)).Find(&res.List)
	return
}

func (dao Admin) Create(req *pb.AdminCreateReq) errors.Error {
	first := DB(dao.Ctx).Where("domain_id=? and account=?", req.DomainID, req.Account).First(&model.RbacAdmin{})
	if first.RowsAffected > 0 {
		return errors.NewDBDuplication("account")
	}
	password, err := helper.NewPassword(req.Account, req.Password, 10)
	if err != nil {
		log.Logger.Error(err)
		return errors.NewInternalError(err)
	}

	err = DB(dao.Ctx).Create(&model.RbacAdmin{
		DomainID: uint(req.DomainID),
		Account:  req.Account,
		Name:     req.Name,
		Password: password,
		Status:   uint(req.Status),
	}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil

}

func (dao Admin) Info(req *pb.DefaultPkReq, fields ...string) model.RbacAdmin {
	var info model.RbacAdmin
	db := DB(dao.Ctx)
	if len(fields) != 0 {
		db = db.Select(strings.Join(fields, ","))
	}
	db.Where("id=? and domain_id=?", req.Pk.(*pb.DefaultPkReq_ID).ID, req.DomainID).First(&info)
	return info
}

func (dao Admin) Update(req *pb.AdminUpdateReq) errors.Error {
	var info model.RbacAdmin
	first := DB(dao.Ctx).Where("id=? and domain_id=? ", req.ID, req.DomainID).First(&info)
	if first.RowsAffected == 0 {
		return errors.NewDBNoData("")
	}
	if req.Password != "" {
		password, err := helper.NewPassword(info.Account, req.Password, 10)
		if err != nil {
			log.Logger.Error(err)
			return errors.NewInternalError(err)
		}
		req.Password = password
	}
	snake := helper.Struct2MapSnakeNoZero(req)
	delete(snake, "id")
	delete(snake, "domain_id")
	err := DB(dao.Ctx).Where("id=?", req.ID).Model(model.RbacAdmin{}).Updates(snake).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}
func (dao Admin) Login(req *pb.LoginReq) (*pb.LoginResData, errors.Error) {

	var info model.RbacAdmin
	var res pb.LoginResData
	first := DB(dao.Ctx).Where(" account=?", req.Account).First(&info)
	if first.RowsAffected == 0 {
		return &res, errors.NewUnauthorized("")
	}
	domainInfo := NewDomain(dao.Ctx).info(int64(info.DomainID))
	if domainInfo.Status != 1 {
		return nil, errors.New(errors.StatusDomainDisable, "域不存在或被禁用")
	}
	err := helper.CheckPassword(info.Account, info.Password, req.Password)
	if err != nil {
		return &res, errors.NewUnauthorized("")
	}
	res.UID = int64(info.ID)
	res.Name = info.Name
	res.DomainID = int64(info.DomainID)
	u := NewRole(dao.Ctx).getRoleIdsByUID(info.ID)
	res.PL = int64(info.PL)
	res.RoleIDs = helper.TransSliceType[uint, int64](u)
	return &res, nil
}
func (dao Admin) getPwd(account, pwd string) (string, error) {
	return helper.NewPassword(account, pwd, 10)
}

func (dao Admin) Registry(req *pb.RegistryReq) (Err errors.Error) {
	tx := DB(dao.Ctx).Begin()
	defer func() {
		if Err != nil {
			err := tx.Rollback().Error
			if err != nil {
				log.Logger.Error("db Rollback", err)
			}
		} else {
			err := tx.Commit().Error
			if err != nil {
				log.Logger.Error("db Commit", err)
			}
		}
	}()
	domain := model.RbacDomain{
		Status: 1,
		Name:   req.Domain,
	}
	err := NewDomain(dao.Ctx).create(tx, &domain)
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	password, err := dao.getPwd(req.Account, req.Password)
	if err != nil {
		log.Logger.Error("dao.admin.Registry.getPwd", err)
		return errors.New(errors.StatusInternalServerError, err)
	}
	admin := model.RbacAdmin{
		DomainID: domain.ID,
		Account:  req.Account,
		Name:     req.Name,
		Password: password,
		Status:   0,
	}
	err = tx.Create(&admin).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	role := model.RbacRole{
		DomainID: domain.ID,
		Name:     "root",
		Desc:     "root",
		Status:   1,
	}
	daoRole := NewRole(dao.Ctx)
	err = daoRole.create(tx, &role)
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	err = dao.setRole(tx, domain.ID, admin.ID, role.ID)
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	permissionList := NewPermission(dao.Ctx).list(tx.Where("domain_id=-1 and status=1"), 0, 0)
	var permissinIds []uint
	for _, permission := range permissionList {
		permissinIds = append(permissinIds, permission.ID)
	}
	err = daoRole.rolePermissionSet(tx, role.ID, domain.ID, permissinIds...)
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}

func (dao Admin) Del(req *pb.DefaultPkReq) errors.Error {
	err := DB(dao.Ctx).Where("id=? and domain_id=?", req.Pk.(*pb.DefaultPkReq_ID).ID, req.DomainID).Delete(&model.RbacAdmin{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	err = DB(dao.Ctx).Where("uid = ? and domain_id=? ", req.Pk.(*pb.DefaultPkReq_ID).ID, req.DomainID).Delete(&model.RbacRoleUser{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil

}

func (dao Admin) AdminInfo(req *pb.DefaultPkReq) *pb.LoginResData {
	var res pb.LoginResData
	ra := dao.Info(req)
	u := NewRole(dao.Ctx).getRoleIdsByUID(ra.ID)
	res.Name = ra.Name
	res.PL = int64(ra.PL)
	res.UID = int64(ra.ID)
	res.RoleIDs = helper.TransSliceType[uint, int64](u)
	return &res
}

func (dao Admin) MenuTree(roleIDs []int64) interface{} {
	pathList := NewPath(dao.Ctx).GetRolePath(1, roleIDs)
	var li = make([]list.TreeNoder, 0, len(pathList))
	for _, p := range pathList {
		li = append(li, &entity.MenuTree{
			Menu: entity.Menu{
				ID:        p.ID,
				ParentID:  p.ParentID,
				Path:      p.Path,
				Name:      p.Name,
				Redirect:  p.Redirect,
				Component: p.Component,
				Meta:      p.Meta,
			},
		})
	}

	// tree := list.List2Tree(li, &entity.MenuTree{})
	return li
}

func (dao Admin) RoleList(req *pb.AdminRoleListReq) []model.RbacRole {
	var arList []model.RbacRoleUser
	d := DB(dao.Ctx).Where("domain_id=? and uid=?  ", req.DomainID, req.UID).Find(&arList)
	if d.RowsAffected == 0 {
		return nil
	}
	us := set.NewSet[uint]()
	for _, v := range arList {
		us.Add(v.RoleID)
	}
	return NewRole(dao.Ctx).list(DB(dao.Ctx), req.DomainID, "", us.List())
}

func (dao Admin) SetRole(req *pb.AdminRoleSetReq) errors.Error {

	return dao.setRole(DB(dao.Ctx), uint(req.DomainID), uint(req.UID), helper.TransSliceType[int64, uint](req.RoleIDs)...)
}

func (dao Admin) setRole(tx *gorm.DB, domainId, uid uint, roleIds ...uint) errors.Error {
	var err error
	if len(roleIds) == 0 {
		err := tx.Where("domain_id=? and uid=? ", domainId, uid).Delete(&model.RbacRoleUser{}).Error
		if err != nil {
			log.Logger.Error(err)
			return errors.NewDBInternal(err)
		}
	} else {
		err := tx.Where("domain_id=? and uid=? and role_id not in ?", domainId, uid, roleIds).Delete(&model.RbacRoleUser{}).Error
		if err != nil {
			log.Logger.Error(err)
			return errors.NewDBInternal(err)
		}
	}

	var list []model.RbacRoleUser
	tx.Where("domain_id=? and uid=? ", domainId, uid).Find(&list)
	us := set.NewSet[uint]()
	for _, v := range list {
		us.Add(v.RoleID)
	}
	newItems := us.GetNewitems(roleIds)
	if len(newItems) == 0 {
		return nil
	}
	var newData []model.RbacRoleUser
	for _, v := range newItems {
		newData = append(newData, model.RbacRoleUser{
			DomainID: domainId,
			RoleID:   v,
			UID:      uid,
		})
	}

	err = tx.Create(&newData).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}
