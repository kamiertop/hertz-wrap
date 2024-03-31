package web

import "embed"

//go:embed index.html
var IndexHtml embed.FS

//go:embed static/assets/*
var Dist embed.FS
