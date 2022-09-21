package form

import (
	"github.com/duxphp/duxgo-ui/lib/node"
)

type RadioOptions struct {
	Key  any
	Name any
	Desc any
}

// Radio 文本输入框
type Radio struct {
	options []RadioOptions
	box     bool
}

// NewRadio 创建文本
func NewRadio() *Radio {
	return &Radio{}
}

// SetOptions 设置选项
func (a *Radio) SetOptions(options []RadioOptions) *Radio {
	a.options = options
	return a
}

// SetBox 设置盒模式
func (a *Radio) SetBox(status bool) *Radio {
	a.box = status
	return a
}

// GetValue 格式化值
func (a *Radio) GetValue(value any, info map[string]any) any {
	if value == nil {
		value = a.options[0].Key
	}
	return value
}

// SaveValue 保存数据
func (a *Radio) SaveValue(value any, data map[string]any) any {
	return value
}

// Render 渲染
func (a *Radio) Render(element node.IField) *node.TNode {
	var options []map[string]any
	for _, item := range a.options {
		if a.box {
			options = append(options, map[string]any{
				"nodeName": "a-radio",
				"child": map[string]any{
					"vSlot:radio": "{checked: true}",
					"nodeName":    "a-space",
					"align":       "start",
					"class":       "custom-radio-card",
					"vBind:class": "{ 'custom-radio-card-checked': checked }",
					"child": []map[string]any{
						{
							"nodeName": "div",
							"class":    "custom-radio-card-mask",
							"child": map[string]any{
								"nodeName": "div",
								"class":    "custom-radio-card-mask-dot",
							},
						},
						{
							"nodeName": "div",
							"child": []map[string]any{
								{
									"nodeName": "div",
									"class":    "custom-radio-card-title",
									"child":    item.Name,
								},
								{
									"nodeName": "a-typography-text",
									"type":     "secondary",
									"child":    item.Desc,
								},
							},
						},
					},
				},
				"value": item.Key,
			})
		} else {
			options = append(options, map[string]any{
				"nodeName": "a-radio",
				"child":    item.Name,
				"value":    item.Key,
			})
		}

	}
	ui := node.TNode{
		"nodeName":          "a-radio-group",
		"child":             options,
		"vModel:modelValue": element.GetUIField(),
		"placeholder":       "请输入" + element.GetName(),
	}
	return &ui
}
