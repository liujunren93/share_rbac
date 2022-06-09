package entity

import (
	"github.com/liujunren93/share_rbac/intenal/model"
	"github.com/liujunren93/share_utils/common/list"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 15:51
 */

type PathListReq struct {
	PageSize int    `form:"page_size"`
	Page     int    `form:"page"`
	Name     string `form:"name"`
}

type PathListRes struct {
	List  []model.RbacPath `json:"list"`
	Total int64            `json:"total"`
}

type MenuTree struct {
	Menu
	Childrens []list.TreeNoder `json:"childrens"`
}

func (m *MenuTree) GetID() uint {
	return m.ID
}

func (m *MenuTree) GetPid() uint {
	return m.ParentID
}

func (m *MenuTree) GetChilds() interface{} {
	return m.Childrens
}

func (m *MenuTree) AddChild(i interface{}) {
	m.Childrens = append(m.Childrens, i.(*MenuTree))
}

type Menu struct {
	ID        uint   `json:"id"`
	ParentID  uint   `json:"parent_id"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Redirect  string `json:"redirect"`
	Component string `json:"component"`
	Meta      string `json:"meta"`
	Target    string `json:"target"`
}
