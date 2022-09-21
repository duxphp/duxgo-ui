package form

import (
	"github.com/duxphp/duxgo-ui/lib/node"
)

// Images 组图上传
type Images struct {
	mode      FileType
	url       string
	manageUrl string
}

// NewImages 创建
func NewImages() *Images {
	return &Images{}
}

// Type 上传模式
func (a *Images) Type(mode FileType) *Images {
	a.mode = mode
	return a
}

// Url 上传地址
func (a *Images) Url(url string) *Images {
	a.url = url
	return a
}

// ManageUrl 管理地址
func (a *Images) ManageUrl(manageUrl string) *Images {
	a.manageUrl = manageUrl
	return a
}

// GetValue 格式化值
func (a *Images) GetValue(value any, info map[string]any) any {
	return value
}

// SaveValue 保存数据
func (a *Images) SaveValue(value any, data map[string]any) any {
	return value
}

// Render 渲染
func (a *Images) Render(element node.IField) *node.TNode {
	ui := node.TNode{
		"nodeName":     "app-images",
		"upload":       a.url,
		"fileUrl":      a.manageUrl,
		"type":         a.mode,
		"vModel:value": element.GetUIField(),
	}
	return &ui
}
