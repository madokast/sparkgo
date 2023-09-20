package main

/*
go build -buildmode=plugin -o export.so exports.go
*/

import (
	"github.com/madokast/sparkgo/internal/exports/export_helper"
)

var ExN2F = export_helper.NamedFuncMap
