package form

import (
	"github.com/duxphp/duxgo-ui/lib/node"
)

// Image 图片上传
type Image struct {
	mode      FileType
	url       string
	manageUrl string
}

// NewImage 创建
func NewImage() *Image {
	return &Image{}
}

// Type 上传模式
func (a *Image) Type(mode FileType) *Image {
	a.mode = mode
	return a
}

// Url 上传地址
func (a *Image) Url(url string) *Image {
	a.url = url
	return a
}

// ManageUrl 管理地址
func (a *Image) ManageUrl(manageUrl string) *Image {
	a.manageUrl = manageUrl
	return a
}

// GetValue 格式化值
func (a *Image) GetValue(value any, info map[string]any) any {
	return value
}

// SaveValue 保存数据
func (a *Image) SaveValue(value any, data map[string]any) any {
	return value
}

// Render 渲染
func (a *Image) Render(element node.IField) *node.TNode {
	ui := node.TNode{
		"nodeName":     "app-file",
		"format":       "image",
		"image":        true,
		"size":         125,
		"upload":       a.url,
		"fileUrl":      a.manageUrl,
		"type":         a.mode,
		"vModel:value": element.GetUIField(),
	}
	return &ui
}
