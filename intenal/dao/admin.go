package dao

import (
	"strings"

	"github.com/liujunren93/share_rbac/intenal/entity"
	"github.com/liujunren93/share_rbac/intenal/model"
	"github.com/liujunren93/share_rbac/log"
	pb "github.com/liujunren93/share_rbac/rbac_pb"
	"github.com/liujunren93/share_utils/common/list"
	"github.com/liujunren93/share_utils/common/set"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/helper"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 15:13
 */

type Admin struct {
}

func (Admin) List(req *pb.AdminListReq) (res entity.AdminListRes) {
	db := DB.Where("domain_id=?", req.DomainID)
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
	db.Limit(pageSize(req.PageSize)).Offset(offset(req.PageSize, req.Page)).Find(&res.List)
	return
}

func (Admin) Create(req *pb.AdminCreateReq) errors.Error {
	first := DB.Where("domain_id=? and account=?", req.DomainID, req.Account).First(&model.RbacAdmin{})
	if first.RowsAffected > 0 {
		return errors.NewDBDuplication("account")
	}
	password, err := helper.NewPassword(req.Account, req.Password, 10)
	if err != nil {
		log.Logger.Error(err)
		return errors.InternalErrorMsg(err)
	}
	err = DB.Create(&model.RbacAdmin{
		DomainID: uint(req.DomainID),
		Account:  req.Account,
		Name:     req.Name,
		Password: password,
		Status:   uint(req.Status),
	}).Error
	if err != nil {
		return errors.NewDBInternal(err)
	}
	return nil

}

func (Admin) Info(req *pb.DefaultPkReq, fields ...string) model.RbacAdmin {
	var info model.RbacAdmin
	db := DB
	if len(fields) != 0 {
		db = db.Select(strings.Join(fields, ","))
	}
	db.Where("id=? and domain_id=?", req.Pk.(*pb.DefaultPkReq_ID).ID, req.DomainID).First(&info)
	return info
}

func (Admin) Update(req *pb.AdminUpdateReq) errors.Error {
	var info model.RbacAdmin
	first := DB.Where("id=? and domain_id=? ", req.ID, req.DomainID).First(&info)
	if first.RowsAffected == 0 {
		return errors.NewDBNoData("")
	}
	if req.Password != "" {
		password, err := helper.NewPassword(info.Account, req.Password, 10)
		if err != nil {
			log.Logger.Error(err)
			return errors.InternalErrorMsg(err)
		}
		req.Password = password
	}
	snake := helper.Struct2MapSnakeNoZero(req)
	delete(snake, "id")
	delete(snake, "domain_id")
	err := DB.Where("id=?", req.ID).Model(model.RbacAdmin{}).Updates(snake).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}
func (Admin) Login(req *pb.LoginReq) (*pb.LoginResData, errors.Error) {
	domainInfo := Domain{}.info(req.DomainID)
	if domainInfo.Status != 1 {
		return nil, errors.New(errors.StatusDomainDisable, "域不存在或被禁用")
	}
	var info model.RbacAdmin
	var res pb.LoginResData
	first := DB.Where(" domain_id=? and account=? and status=1", req.DomainID, req.Account).First(&info)
	if first.RowsAffected == 0 {
		return &res, errors.NewUnauthorized("")
	}
	err := helper.CheckPassword(info.Account, info.Password, req.Password)
	if err != nil {
		return &res, errors.NewUnauthorized("")
	}
	res.UID = int64(info.ID)
	res.Name = info.Name
	u := Role{}.getRoleIdsByUID(info.ID)
	res.RoleIDs = helper.TransSliceType[uint, int64](u)
	return &res, nil
}
func (Admin) Del(req *pb.DefaultPkReq) errors.Error {
	err := DB.Where("id=? and domain_id=?", req.Pk.(*pb.DefaultPkReq_ID).ID, req.DomainID).Delete(&model.RbacAdmin{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil

}

func (dao Admin) AdminInfo(req *pb.DefaultPkReq) *pb.LoginResData {
	var res pb.LoginResData
	ra := dao.Info(req)
	u := Role{}.getRoleIdsByUID(ra.ID)
	res.Name = ra.Name
	res.UID = int64(ra.ID)
	res.RoleIDs = helper.TransSliceType[uint, int64](u)
	return &res
}

func (Admin) MenuTree(roleIDs []int64) interface{} {
	pathList := Path{}.GetRolePath(1, roleIDs)
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

func (Admin) RoleList(req *pb.AdminRoleListReq) []model.RbacRole {
	var arList []model.RbacRoleUser
	d := DB.Where("domain_id=? and uid=?  ", req.DomainID, req.UID).Find(&arList)
	if d.RowsAffected == 0 {
		return nil
	}
	us := set.NewSet[uint]()
	for _, v := range arList {
		us.Add(v.RoleID)
	}
	return Role{}.list(DB, req.DomainID, "", us.List())
}

func (Admin) SetRole(req *pb.AdminRoleSetReq) errors.Error {
	var err error
	if len(req.RoleIDs) == 0 {
		err := DB.Where("domain_id=? and uid=? ", req.DomainID, req.UID).Delete(&model.RbacRoleUser{}).Error
		if err != nil {
			log.Logger.Error(err)
			return errors.NewDBInternal(err)
		}
	} else {
		err := DB.Where("domain_id=? and uid=? and role_id not in ?", req.DomainID, req.UID, req.RoleIDs).Delete(&model.RbacRoleUser{}).Error
		if err != nil {
			log.Logger.Error(err)
			return errors.NewDBInternal(err)
		}
	}

	var list []model.RbacRoleUser
	DB.Where("domain_id=? and uid=? ", req.DomainID, req.UID).Find(&list)
	us := set.NewSet[uint]()
	for _, v := range list {
		us.Add(v.RoleID)
	}
	newItems := us.GetNewitems(helper.TransSliceType[int64, uint](req.RoleIDs))
	if len(newItems) == 0 {
		return nil
	}
	var newData []model.RbacRoleUser
	for _, v := range newItems {
		newData = append(newData, model.RbacRoleUser{
			DomainID: uint(req.DomainID),
			RoleID:   v,
			UID:      uint(req.UID),
		})
	}

	err = DB.Create(&newData).Error
	if err != nil {
		return errors.NewDBInternal(err)
	}
	return nil
}
