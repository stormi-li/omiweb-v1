package omiweb

import (
	"embed"
)

var StaticPath = "static"

var IndexPath = "/index.html"

//go:embed templateSource/static/*
var templateSource embed.FS
