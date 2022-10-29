package form

import (
	"github.com/duxphp/duxgo-ui/lib/node"
	"github.com/duxphp/duxgo/util/function"
)

// Cascader 级联选择器
type Cascader struct {
	options     []map[string]any
	placeholder string
	multi       bool
	maxCount    uint
	url         string
	urlParams   map[string]any
	image       string
	desc        string
	icon        string
	leaf        bool
	tree        bool
	params      map[string]any
}

// NewCascader 创建选择器
func NewCascader() *Cascader {
	return &Cascader{}

}

// SetMulti 设置多选
func (a *Cascader) SetMulti(num ...uint) *Cascader {
	a.multi = true
	if len(num) > 0 {
		a.maxCount = num[0]
	}
	return a
}

// SetOptions 设置选项
func (a *Cascader) SetOptions(options []map[string]any) *Cascader {
	a.options = append(a.options, options...)
	return a
}

// SetPlaceholder 提示信息
func (a *Cascader) SetPlaceholder(content string) *Cascader {
	a.placeholder = content
	return a
}

// SetUrl 远程搜索
func (a *Cascader) SetUrl(url string, params map[string]any) *Cascader {
	a.url = url
	a.urlParams = params
	return a
}

// SetTree 设置树形格式化
func (a *Cascader) SetTree() *Cascader {
	a.tree = true
	return a
}

// SetLeaf 设置叶级选择
func (a *Cascader) SetLeaf() *Cascader {
	a.leaf = true
	return a
}

func (a *Cascader) SetParams(key string, value any) *Cascader {
	a.params[key] = value
	return a
}

// GetValue 格式化值
func (a *Cascader) GetValue(value any, info map[string]any) any {
	return value
}

// SaveValue 保存数据
func (a *Cascader) SaveValue(value any, data map[string]any) any {
	return value
}

// Render 渲染
func (a *Cascader) Render(element node.IField) *node.TNode {

	placeholder := a.placeholder

	if placeholder == "" {
		placeholder = "请输入" + element.GetName()
	}

	options := []map[string]any{}

	if a.tree {
		options = *cascaderLoop(a.options)
	} else {
		options = a.options
	}

	nParams := node.TNode{
		"placeholder":   placeholder,
		"options":       options,
		"allow-search":  true,
		"allow-clear":   true,
		"multiple":      a.multi,
		"max-tag-count": a.maxCount,
	}

	ui := node.TNode{
		"nodeName":     "app-cascader",
		"nParams":      nParams,
		"vModel:value": element.GetUIField(),
		"placeholder":  "请输入" + element.GetName(),
	}

	if a.url != "" {
		ui["vBind:dataUrl"] = function.BuildUrl(a.url, a.urlParams, false)
	}
	return &ui
}

func cascaderLoop(data []map[string]any) *[]map[string]any {
	var dataArr []map[string]any
	for _, item := range data {
		tmpData := map[string]any{
			"label": item["name"],
			"value": item["id"],
		}
		if item["children"] != nil {
			tmpData["children"] = cascaderLoop(item["children"].([]map[string]any))
		}
		dataArr = append(dataArr, tmpData)
	}
	return &dataArr
}
