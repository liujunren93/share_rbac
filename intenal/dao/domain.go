package dao

import (
	"github.com/liujunren93/share_rbac/intenal/entity"
	"github.com/liujunren93/share_rbac/intenal/model"
	"github.com/liujunren93/share_rbac/log"
	"github.com/liujunren93/share_rbac/pb"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/helper"
)

/**
* @Author: liujunren
* @Date: 2022/3/24 11:58
 */
type Domain struct {
}

func (Domain) List(req *pb.DomainListReq) entity.DomainListRes {
	var res entity.DomainListRes
	db := DB
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Domain != "" {
		db = db.Where("name like ?", "%"+req.Domain+"%")
	}
	db.Model(&model.RbacDomain{}).Count(&res.Total)

	db.Limit(pageSize(req.PageSize)).Offset(offset(req.PageSize, req.Page)).Find(&res.List)
	return res
}

func (Domain) Info(req *pb.DefaultPkReq) model.RbacDomain {
	var info model.RbacDomain
	DB.Where("id = ?", req.Pk.(*pb.DefaultPkReq_ID).ID).First(&info)
	return info
}

func (Domain) Create(req *pb.DomainCreateReq) errors.Error {
	err := DB.Create(&model.RbacDomain{
		Name:   req.Name,
		Domain: req.Domain,
		Status: int8(req.Status),
	}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternalErr(err)
	}
	return nil
}
func (Domain) Update(req *pb.DomainUpdateReq) errors.Error {
	zero := helper.Struct2MapSnakeNoZero(req)
	delete(zero, "id")
	err := DB.Where("id=?", req.Id).Model(&model.RbacDomain{}).Updates(zero).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternalErr(err)
	}
	return nil
}

func (Domain) Del(req *pb.DefaultPkReq) errors.Error {

	err := DB.Where("id=?", req.Pk).Delete(&model.RbacDomain{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternalErr(err)
	}
	return nil
}
