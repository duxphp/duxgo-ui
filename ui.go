package ui

import (
	"embed"
	"github.com/duxphp/duxgo/bootstrap"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

//go:embed public/*
var StaticFs embed.FS

// ConfigManifest 前端配置
var ConfigManifest map[string]any

// New UI库
func New(t *bootstrap.Bootstrap) {
	// 解析媒体文件
	jsonPath, err := StaticFs.Open("public/manifest.json")
	if err != nil {
		panic(err.Error())
	}
	config := viper.New()
	config.SetConfigType("json")
	err = config.ReadConfig(jsonPath)
	if err != nil {
		panic(err.Error())
	}
	jsonPath.Close()
	ConfigManifest = config.GetStringMap("src/main.js")

	// 前端静态文件
	t.App.StaticFS("/", echo.MustSubFS(StaticFs, "public"))
}
