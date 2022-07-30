package ui

import (
	"embed"
	"github.com/duxphp/duxgo/bootstrap"
	"github.com/labstack/echo/v4"
)

//go:embed public/*
var StaticFs embed.FS

func New(t *bootstrap.Bootstrap) {
	t.App.StaticFS("/", echo.MustSubFS(StaticFs, "public"))
}
