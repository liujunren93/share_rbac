package dao

import (
	"context"

	"github.com/liujunren93/share_rbac/intenal/entity"
	"github.com/liujunren93/share_rbac/intenal/model"
	"github.com/liujunren93/share_rbac/log"
	pb "github.com/liujunren93/share_rbac/rbac_pb"
	"github.com/liujunren93/share_utils/errors"
	"github.com/liujunren93/share_utils/helper"
	"gorm.io/gorm"
)

/**
* @Author: liujunren
* @Date: 2022/3/24 11:58
 */
type Domain struct {
	dao
}

func NewDomain(ctx context.Context) Domain {
	return Domain{dao{ctx}}
}

func (dao Domain) List(req *pb.DomainListReq) entity.DomainListRes {
	var res entity.DomainListRes
	db := DB(dao.Ctx)
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Domain != "" {
		db = db.Where("domain like ?", "%"+req.Domain+"%")
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.Domain)
	}
	db.Model(&model.RbacDomain{}).Count(&res.Total)
	if req.SortField != "" {
		db = db.Order(req.SortField + " " + req.SortOrder)
	}
	db.Limit(pageSize(req.PageSize)).Offset(offset(req.PageSize, req.Page)).Find(&res.List)
	return res
}

func (dao Domain) Info(req *pb.DefaultPkReq) model.RbacDomain {
	var info model.RbacDomain
	dao.info(req.Pk.(*pb.DefaultPkReq_ID).ID)
	return info
}

func (dao Domain) info(id int64) model.RbacDomain {
	var info model.RbacDomain
	DB(dao.Ctx).Where("id = ?", id).First(&info)
	return info
}

func (dao Domain) Create(req *pb.DomainCreateReq) errors.Error {
	domain := model.RbacDomain{
		Name:   req.Name,
		Domain: req.Domain,
		Status: int8(req.Status),
	}
	err := dao.create(DB(dao.Ctx), &domain)
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}
func (dao Domain) create(tx *gorm.DB, domain *model.RbacDomain) error {
	err := tx.Create(domain).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao Domain) Update(req *pb.DomainUpdateReq) errors.Error {
	zero := helper.Struct2MapSnakeNoZero(req)
	delete(zero, "id")
	err := DB(dao.Ctx).Where("id=?", req.ID).Model(&model.RbacDomain{}).Updates(zero).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}

func (dao Domain) Del(req *pb.DefaultPkReq) errors.Error {

	err := DB(dao.Ctx).Where("id=?", req.Pk.(*pb.DefaultPkReq_ID).ID).Delete(&model.RbacDomain{}).Error
	if err != nil {
		log.Logger.Error(err)
		return errors.NewDBInternal(err)
	}
	return nil
}
