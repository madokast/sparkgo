package exports

import "github.com/madokast/sparkgo/internal/exports/export_helper"

func (__remote__) PrintHello() {
	println("Hwllo, world!")
}

func (__remote__) AddInt64(a, b int64) int64 {
	return a + b
}

// --------------- reflect helpers -----------------------

type __remote__ export_helper.FunctionGroup

var F = __remote__{}
