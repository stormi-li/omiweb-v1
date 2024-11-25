package omiweb

import (
	"embed"
)

var sourcePath = "static"

var indexPath = "/index.html"

//go:embed templateSource/static/*
var templateSource embed.FS
