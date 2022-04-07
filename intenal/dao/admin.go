package dao

import (
	"github.com/liujunren93/share_rbac/intenal/entity"
	"github.com/liujunren93/share_rbac/intenal/model"
	"github.com/liujunren93/share_rbac/log"
	"github.com/liujunren93/share_rbac/pb"
	"github.com/liujunren93/share_utils/commond/list"
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
	db.Model(&model.RbacAdmin{}).Count(&res.Total)
	DB.Limit(pageSize(req.PageSize)).Offset(offset(req.PageSize, req.Page)).Find(&res.List)
	return
}

func (Admin) Create(req *pb.AdminCreateReq) errors.Error {
	first := DB.Where("account=?", req.Account).First(&model.RbacAdmin{})
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
	}).Error
	if err != nil {
		return errors.NewDBInternalErr(err)
	}
	return nil

}

func (Admin) Info(req *pb.DefaultPkReq) model.RbacAdmin {
	var info model.RbacAdmin
	DB.Where("id=?", req.Pk).First(&info)
	return info
}

func (Admin) Update(req *pb.AdminUpdateReq) errors.Error {
	var info model.RbacAdmin
	first := DB.Where("id=? ", req.ID).First(&info)
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
	err := DB.Where("id=?", req.ID).Model(model.RbacAdmin{}).Updates(snake).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternalErr(err)
	}
	return nil
}
func (Admin) CheckPwd(req *pb.CheckPwdReq) errors.Error {
	var info model.RbacAdmin
	first := DB.Where("account=? ", req.Account).First(&info)
	if first.RowsAffected > 0 {
		return errors.NewUnauthorized("")
	}
	err := helper.CheckPassword(info.Account, info.Password, req.Password)
	if err != nil {
		return errors.NewUnauthorized("")
	}
	return nil
}
func (Admin) Del(Id int64) errors.Error {
	err := DB.Where("id=?", Id).Delete(&model.RbacAdmin{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternalErr(err)
	}
	return nil

}

func (Admin) MenuTree(uid uint) interface{} {
	pathList := Path{}.GetPathByUid(1, uid)
	var li = make([]list.TreeNoder, 0, len(pathList))
	for _, p := range pathList {
		li = append(li, &entity.MenuTree{
			RbacPath:  p,
			Childrens: nil,
		})
	}
	tree := list.List2Tree(li, &entity.MenuTree{})
	return tree
}
