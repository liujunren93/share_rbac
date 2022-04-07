package entity

import "github.com/liujunren93/share_rbac/intenal/model"

/**
* @Author: liujunren
* @Date: 2022/3/9 9:47
 */

type PermissionListRes struct {
	List  []model.RbacPermission `json:"list"`
	Total int64                `json:"total"`
}
