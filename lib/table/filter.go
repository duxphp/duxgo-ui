package table

import (
	"github.com/duxphp/duxgo-ui/lib/node"
	"gorm.io/gorm"
)

// IFilter 筛选UI接口
type IFilter interface {
	Render(element node.IField) *node.TNode
}

// Filter 筛选结构
type Filter struct {
	Name    string
	Field   string
	Where   func(string, *gorm.DB)
	Quick   bool
	Default any
	UI      IFilter
	Format  func(value any) any
}

// SetQuick 设置快速筛选
func (a *Filter) SetQuick(quick bool) *Filter {
	a.Quick = quick
	return a
}

// SetWhere 设置条件
func (a *Filter) SetWhere(where func(string, *gorm.DB)) *Filter {
	a.Where = where
	return a
}

// SetDefault 设置默认值
func (a *Filter) SetDefault(value any) *Filter {
	a.Default = value
	return a
}

// SetFormat 格式化字段
func (a *Filter) SetFormat(callback func(value any) any) *Filter {
	a.Format = callback
	return a
}

// SetUI 设置UI
func (a *Filter) SetUI(ui IFilter) *Filter {
	a.UI = ui
	return a
}

// GetName 获取名称
func (a *Filter) GetName() string {
	return a.Name
}

// GetUIField 获取模型字段名
func (a *Filter) GetUIField(field ...string) string {
	content := a.Field
	if len(field) > 0 {
		content = field[0]
	}
	return "data.filter['" + content + "']"
}

// Render 渲染UI
func (a *Filter) Render() node.TNode {

	if a.Quick {
		return node.TNode{
			"nodeName": "div",
			"class":    "lg:w-40",
			"child":    a.UI.Render(a),
		}
	}
	return node.TNode{
		"nodeName": "div",
		"class":    "my-2",
		"child": []node.TNode{
			{
				"nodeName": "div",
				"child":    a.Name,
			},
			{
				"nodeName": "div",
				"class":    "mt-2",
				"child":    a.UI.Render(a),
			},
		},
	}
}
