package formLayout

import (
	form2 "github.com/duxphp/duxgo-ui/lib/form"
	"github.com/duxphp/duxgo-ui/lib/node"
)

type Block struct {
	form   *form2.Form
	data   map[string]any
	dialog bool
	title  string
}

// NewBlock 块布局
func NewBlock(title string) *Block {
	return &Block{
		title: title,
	}
}

// SetData 设置数据
func (a *Block) SetData(data map[string]any) {
	a.data = data
}

// SetDialog 设置弹窗
func (a *Block) SetDialog(dialog bool) {
	a.dialog = dialog
}

// Column 列元素
func (a *Block) Column(callback func(form *form2.Form), opt ...any) {
	formUI := form2.NewForm()
	formUI.SetData(a.data)
	formUI.SetDialog(a.dialog)
	a.form = formUI
	callback(a.form)
}

// Form 获取表单
func (a *Block) Form(index ...int) *form2.Form {
	return a.form
}

// Expand 展开元素
func (a *Block) Expand() []*form2.Element {
	return a.form.ExpandElement()
}

// Render 渲染
func (a *Block) Render() *node.TNode {
	element := a.form.RenderElement()
	ui := node.TNode{
		"nodeName": "div",
		"class":    "pt-1",
		"child": []node.TNode{
			{
				"nodeName":    "a-divider",
				"orientation": "left",
				"child":       a.title,
			},
			{
				"nodeName": "div",
				"child":    element,
			},
		},
	}
	return &ui
}
