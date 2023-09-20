package export_helper

import (
	"fmt"
	"reflect"

	"github.com/madokast/sparkgo/internal/exports"
	"github.com/madokast/sparkgo/internal/logger"
	"github.com/madokast/sparkgo/internal/reflect_utils"
)

var NamedFuncMap = map[string]any{}     // 方法名->方法
var FuncNamedMap = map[uintptr]string{} // 方法地址->方法名。本地使用

func init() {
	funVal := reflect.ValueOf(exports.F)
	funType := reflect.TypeOf(exports.F)
	for i := 0; i < funType.NumMethod(); i++ {
		m := funType.Method(i)   // 仅仅用于获取函数名
		mVal := funVal.Method(i) // 方法，用于远程调用
		mValType := mVal.Type()  // 方法的类型，从中可以获取函数地址
		logger.Debug("register func", m.Name, fmt.Sprintf("%v", reflect_utils.InterfacePointer(mValType)))
		NamedFuncMap[m.Name] = reflect.MakeFunc(mValType, func(args []reflect.Value) (results []reflect.Value) {
			logger.Debug("export func", m.Name, "called")
			return mVal.Call(args)
		}).Interface()
		FuncNamedMap[reflect_utils.InterfacePointer(mValType) /*方法地址*/] = m.Name
	}
}
