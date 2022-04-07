package entity

import (
	"github.com/liujunren93/share_rbac/intenal/model"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 15:15
 */

type AdminListReq struct {
	PageSize int    `form:"page_size"`
	Page     int    `form:"page"`
	Name     string `form:"name"`
}

type AdminListRes struct {
	List  []model.RbacAdmin `json:"list"`
	Total int64             `json:"total"`
}

type AdminUpdateReq struct {
	Method int `json:"method" binding:"required"` // 1:用戶修改密码，2：管理员修改密码，3：修改用户信息
	model.RbacAdmin
}
