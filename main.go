package main

import (
	"errors"
	"plugin"
	"reflect"

	"github.com/madokast/sparkgo/internal/exports"
	"github.com/madokast/sparkgo/internal/exports/export_helper"
	"github.com/madokast/sparkgo/internal/logger"
	"github.com/madokast/sparkgo/internal/reflect_utils"
)

func dynamicCall(fun any, params ...any) (any, error) {
	funVal := reflect.ValueOf(fun)
	name, ok := export_helper.FuncNamedMap[reflect_utils.ValueType(funVal)]
	if !ok {
		return nil, errors.New("un-registered function")
	}
	p, err := plugin.Open("export.so")
	if err != nil {
		return nil, err
	}
	n2f, err := p.Lookup("ExN2F")
	if err != nil {
		return nil, err
	}
	exFunc, ok := (*n2f.(*map[string]any))[name]
	if !ok {
		return nil, errors.New("un-registered function " + name)
	}
	logger.Info(reflect_utils.FuncString(exFunc))

	var values []reflect.Value
	// values = append(values, reflect.ValueOf(exports.F))
	for _, param := range params {
		values = append(values, reflect.ValueOf(param))
	}
	res := reflect.ValueOf(exFunc).Call(values)

	switch len(res) {
	case 0:
		return nil, nil
	case 1:
		return res[0].Interface(), nil
	}
	if err, ok := res[1].Interface().(error); ok {
		return res[0].Interface(), err
	} else {
		return res[0].Interface(), nil
	}
}

func main() {
	logger.Info(dynamicCall(exports.F.AddInt64, int64(123), int64(500)))
	logger.Info(dynamicCall(exports.F.AddInt64, int64(100), int64(55)))
	logger.Info(dynamicCall(exports.F.AddInt64, int64(5), int64(7)))
}
