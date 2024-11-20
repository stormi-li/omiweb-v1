package omiweb

import (
	"embed"
)

const target_path = "static"
const index_path = "/index.html"

//go:embed TemplateSource/*
var templateSource embed.FS
