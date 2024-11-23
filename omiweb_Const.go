package omiweb

import (
	"embed"
)

const target_path = "static"
const index_path = "/index.html"

//go:embed templateSource/static/*
var templateSource embed.FS
