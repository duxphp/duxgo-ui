package form

import (
	"encoding/json"
	"fmt"
	"github.com/duxphp/duxgo-ui/lib/node"
)

// Data 动态数据编辑器
type Data struct {
	uiField []*DataField
	option  bool
	max     uint
	min     uint
	wrap    bool
}

type DataField struct {
	Name string
	Key  string
	UI   IElement
}

// NewData 创建日期
func NewData() *Data {
	return &Data{}
}

func (a *Data) AddField(ui ...*DataField) *Data {
	a.uiField = ui
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

// 数据元素
type dataElement struct {
	field string
	name  string
}

func (a *dataElement) GetUIField(field ...string) string {
	return fmt.Sprintf("value.%s", a.field)
}

func (a *dataElement) GetName() string {
	return a.name
}

// Render 渲染
func (a *Data) Render(element node.IField) *node.TNode {
	inner := []map[string]any{}
	data := map[string]any{}
	for _, field := range a.uiField {
		inner = append(inner, map[string]any{
			"nodeName": "div",
			"class":    "flex-1",
			"child":    field.UI.Render(&dataElement{field: field.Key, name: field.Name}),
		})
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
			"class":    "flex flex-grow gap-4 items-center ",
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
