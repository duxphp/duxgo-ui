package form

import (
	"encoding/json"
	"fmt"
	"github.com/duxphp/duxgo-ui/lib/node"
)

// Data 动态数据编辑器
type Data struct {
	column []map[string]any
	option bool
	max    uint
	min    uint
	wrap   bool
}

// NewData 创建日期
func NewData() *Data {
	return &Data{}
}

func (a *Data) AddText(name string, field string, width string) *Data {
	a.column = append(a.column, map[string]any{
		"name":  name,
		"key":   field,
		"type":  "text",
		"width": width,
	})
	return a
}

func (a *Data) AddSelect(name string, field string, data map[any]string, width string) *Data {
	a.column = append(a.column, map[string]any{
		"name":  name,
		"key":   field,
		"type":  "select",
		"data":  data,
		"width": width,
	})
	return a
}

func (a *Data) AddImage(name string, field string, url string, width string) *Data {
	a.column = append(a.column, map[string]any{
		"name":  name,
		"key":   field,
		"type":  "image",
		"width": width,
		"url":   url,
	})
	return a
}

func (a *Data) AddArea(name string, field string, url string, multi bool, width string) *Data {
	a.column = append(a.column, map[string]any{
		"name":  name,
		"key":   field,
		"type":  "area",
		"width": width,
		"url":   url,
		"multi": multi,
	})
	return a
}

func (a *Data) AddShow(name string, field string, width string) *Data {
	a.column = append(a.column, map[string]any{
		"name":  name,
		"key":   field,
		"type":  "show",
		"width": width,
	})
	return a
}

func (a *Data) AddHtml(field string, html any) *Data {
	a.column = append(a.column, map[string]any{
		"key":  field,
		"type": "html",
		"data": html,
	})
	return a
}

func (a *Data) AddHidden(name string, field string) *Data {
	a.column = append(a.column, map[string]any{
		"name": name,
		"key":  field,
		"type": "show",
	})
	return a
}

func (a *Data) SetOption(status bool) *Data {
	a.option = status
	return a
}

func (a *Data) SetMax(num uint) *Data {
	a.max = num
	return a
}

func (a *Data) SetMin(num uint) *Data {
	a.min = num
	return a
}

func (a *Data) SetWrap(status bool) *Data {
	a.wrap = status
	return a
}

// GetValue 格式化值
func (a *Data) GetValue(value any, info map[string]any) any {
	return value
}

// SaveValue 保存数据
func (a *Data) SaveValue(value any, data map[string]any) any {
	return value
}

// Render 渲染
func (a *Data) Render(element node.IField) *node.TNode {

	inner := []map[string]any{}
	data := map[string]any{}

	for _, column := range a.column {
		data[column["key"].(string)] = ""
		field := fmt.Sprintf("value['%s']", column["key"])
		innerNode := map[string]any{}

		if column["type"] == "text" {
			innerNode = map[string]any{
				"nodeName": "div",
				"class":    "flex-grow",
				"child": map[string]any{
					"nodeName":          "a-input",
					"placeholder":       "请输入" + column["name"].(string),
					"vModel:modelValue": field,
				},
			}
		}

		if column["type"] == "image" {
			innerNode = map[string]any{
				"nodeName": "div",
				"class":    "flex-none",
				"child": map[string]any{
					"nodeName":     "app-file",
					"image":        true,
					"mini":         true,
					"size":         8,
					"vModel:value": field,
				},
			}
		}

		if column["type"] == "select" {
			options := []map[string]any{}
			for k, v := range column["data"].(map[any]string) {
				options = append(options, map[string]any{
					"label": v,
					"value": k,
				})
			}
			innerNode = map[string]any{
				"nodeName": "div",
				"class":    "flex-grow",
				"child": map[string]any{
					"nodeName": "app-select",
					"nParams": map[string]any{
						"placeholder": "请输入" + column["name"].(string),
						"options":     options,
					},
					"vModel:modelValue": field,
				},
			}
		}

		if column["type"] == "area" {
			innerNode = map[string]any{
				"nodeName": "div",
				"class":    "flex-grow",
				"child": map[string]any{
					"nodeName": "app-cascader",
					"nParams": map[string]any{
						"placeholder":  "请输入" + column["name"].(string),
						"allow-search": true,
						"path-mode":    true,
						"clearable":    true,
						"multiple":     column["multi"].(bool),
					},
					"dataUrl":      column["url"],
					"vModel:value": field,
				},
			}
		}

		if column["type"] == "show" {
			innerNode = map[string]any{
				"nodeName": "div",
				"class":    "flex-grow",
				"child":    field,
			}
		}

		if column["width"] != "" {
			innerNode["style"] = map[string]any{
				"width": column["width"],
			}
		}

		inner = append(inner, innerNode)
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	ui := node.TNode{
		"nodeName":        "app-dynamic-data",
		"vBind:on-create": "() => { return " + string(marshal) + " }",
		"vModel:value":    element.GetUIField(),
		"placeholder":     "请输入" + element.GetName(),
		"renderRow: value, index": map[string]any{
			"nodeName": "div",
			"class":    "flex flex-grow gap-4 items-center",
			"child":    inner,
		},
	}
	if a.max > 0 {
		ui["max"] = a.max
	}
	if a.min > 0 {
		ui["min"] = a.min
	}
	return &ui
}
