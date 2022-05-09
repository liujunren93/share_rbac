package entity

import "github.com/liujunren93/share_rbac/intenal/model"

/**
* @Author: liujunren
* @Date: 2022/2/28 15:37
 */

type RoleListReq struct {
	PageSize int    `form:"page_size"`
	Page     int    `form:"page"`
	Name     string `form:"name"`
}

type RoleListRes struct {
	Total int64            `json:"total"`
	List  []model.RbacRole `json:"list"`
}
