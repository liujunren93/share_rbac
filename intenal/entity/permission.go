package entity

import "github.com/liujunren93/share_rbac/intenal/model"

/**
* @Author: liujunren
* @Date: 2022/3/9 9:47
 */

type PermissionListRes struct {
	List  []model.RbacPermission `json:"list"`
	Total int64                  `json:"total"`
}

type UserPermissionListRes struct {
	List UserPermission `json:"list"`
}

type UserPermission struct {
	PID        uint     `json:"-"`
	PathID     uint     `json:"-"`
	Path       string   `json:"permissionId"`
	ActionList []string `json:"actionList"`
}
