package frpc

import (
	"embed"

	"github.com/shuffleman/frp/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
