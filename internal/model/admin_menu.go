package model

import (
	"encoding/json"
)

type AdminMenu struct {
	Model
	ParentID   uint64       `json:"parent_id" `  // 父级菜单
	Name       string       `json:"name"`        // 菜单名称
	Permission string       `json:"permission"`  // 权限标识
	Path       string       `json:"path"`        // 菜单路径
	Type       uint8        `json:"type"`        // 1菜单，2操作
	Icon       string       `json:"icon"`        // 图标
	Component  string       `json:"component"`   // 组件路径
	IsDisabled uint8        `json:"is_disabled"` // 是否可用，1是，0否
	Sort       uint64       `json:"sort"`
	Children   []*AdminMenu `json:"children" gorm:"foreignKey:ParentID;references:ID"`
}

func (model AdminMenu) TableName() string {
	return "admin_menu"
}

func (model AdminMenu) MarshalJSON() ([]byte, error) {
	type Alias AdminMenu

	a := struct {
		Alias
	}{
		Alias: (Alias)(model),
	}

	if a.Children == nil {
		a.Children = make([]*AdminMenu, 0)
	}

	return json.Marshal(a)
}

type AdminMenuList []*AdminMenu

type TreeList map[uint64]AdminMenuList

func (menuList AdminMenuList) ToTree() AdminMenuList {
	tree := map[uint64]AdminMenuList{}
	for _, menu := range menuList {
		tree[menu.ParentID] = append(tree[menu.ParentID], menu)
	}
	list := make([]*AdminMenu, len(tree[0]))
	for k, v := range tree[0] {
		list[k] = v
	}

	list = procss(list, tree, 0)
	return list
}

func procss(list AdminMenuList, tree TreeList, level int) AdminMenuList {
	if level >= 10 {
		return list
	}
	for _, l := range list {
		l.Children = procss(tree[l.ID], tree, level+1)

	}
	return list
}
