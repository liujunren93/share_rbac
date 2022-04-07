package entity

/**
* @Author: liujunren
* @Date: 2022/2/28 15:37
 */

type RoleListReq struct {
	PageSize int    `form:"page_size"`
	Page     int    `form:"page"`
	Name     string `form:"name"`
}
