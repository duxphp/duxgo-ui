package table

import (
	"github.com/duxphp/duxgo-ui/lib/node"
	"gorm.io/gorm"
)

// Tab 筛选结构
type Tab struct {
	Name  string
	Icon  string
	Where func(*gorm.DB)
}

// SetIcon 设置图标
func (a *Tab) SetIcon(icon string) *Tab {
	a.Icon = icon
	return a
}

// SetWhere 设置条件
func (a *Tab) SetWhere(where func(*gorm.DB)) *Tab {
	a.Where = where
	return a
}

// Render 渲染UI
func (a *Tab) Render(index int) node.TNode {
	icon := node.TNode{}
	if a.Icon != "" {
		icon = node.TNode{
			"nodeName": a.Icon,
		}
	}
	return node.TNode{
		"nodeName": "a-radio",
		"value":    index,
		"child": []node.TNode{
			icon,
			{
				"nodeName": "span",
				"child":    " " + a.Name,
			},
		},
	}
}
