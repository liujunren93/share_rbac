package entity

import "github.com/liujunren93/share_rbac/intenal/model"

/**
* @Author: liujunren
* @Date: 2022/3/24 12:00
 */
type DomainListRes struct {
	List  []model.RbacDomain `json:"list"`
	Total int64              `json:"total"`
}
