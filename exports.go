package main

/*
go build -buildmode=plugin -o export.so exports.go
*/

import (
	"github.com/madokast/sparkgo/internal/exports"
	"github.com/madokast/sparkgo/internal/exports/export_helper"
)

var ExN2F, _ = export_helper.Export(exports.F)
