package form

import (
	"github.com/duxphp/duxgo-ui/lib/node"
)

type CheckboxOptions struct {
	Key  any
	Name any
	Desc any
}

// Checkbox 多选框
type Checkbox struct {
	options []CheckboxOptions
	card    bool
}

// NewCheckbox 创建多选
func NewCheckbox() *Checkbox {
	return &Checkbox{}
}

// SetOptions 设置选项
func (a *Checkbox) SetOptions(options []CheckboxOptions) *Checkbox {
	a.options = options
	return a
}

// SetCard 设置盒模式
func (a *Checkbox) SetCard(status bool) *Checkbox {
	a.card = status
	return a
}

// GetValue 格式化值
func (a *Checkbox) GetValue(value any, info map[string]any) any {
	return value
}

// SaveValue 保存数据
func (a *Checkbox) SaveValue(value any, data map[string]any) any {
	return value
}

// Render 渲染
func (a *Checkbox) Render(element node.IField) *node.TNode {
	var options []map[string]any
	for _, item := range a.options {
		if a.card {
			options = append(options, map[string]any{
				"nodeName": "a-checkbox	",
				"child": map[string]any{
					"vSlot:checkbox": "{checked}",
					"nodeName":       "a-space",
					"align":          "start",
					"vBind:class":    "checked ? 'custom-radio-card custom-radio-card-checked' : 'custom-radio-card'",
					"child": []map[string]any{
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
				"nodeName": "a-checkbox",
				"child":    item.Name,
				"value":    item.Key,
			})
		}

	}
	ui := node.TNode{
		"nodeName":          "a-checkbox-group",
		"child":             options,
		"vModel:modelValue": element.GetUIField(),
		"placeholder":       "请输入" + element.GetName(),
	}
	return &ui
}
