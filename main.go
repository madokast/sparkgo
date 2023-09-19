package main

import (
	"plugin"
	"reflect"

	"github.com/madokast/sparkgo/internal/logger"
	"github.com/madokast/sparkgo/internal/reflect_utils"
)

func main() {
	p, err := plugin.Open("exports/export.so")
	if err != nil {
		panic(err)
	}
	add, err := p.Lookup("AddInt64")
	if err != nil {
		panic(err)
	}
	v10 := reflect.ValueOf(int64(10))
	v22 := reflect.ValueOf(int64(22))

	res := reflect.ValueOf(add).Call([]reflect.Value{v10, v22})
	logger.Info("add 10 22 = ", res[0].Interface().(int64))

	logger.Info(reflect_utils.FuncString(add))
}
