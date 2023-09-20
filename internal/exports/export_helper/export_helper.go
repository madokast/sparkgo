package export_helper

import (
	"fmt"
	"reflect"

	"github.com/madokast/sparkgo/internal/logger"
	"github.com/madokast/sparkgo/internal/reflect_utils"
)

type iFunctionGroup any
type FunctionGroup struct{}

type funcName = string
type function = any

// 导出所有的函数 functions
// name2func 方法名->方法。远程使用
// func2name 方法地址->方法名。本地使用
func Export(functions iFunctionGroup) (name2func map[funcName]function, func2name map[uintptr]funcName) {
	name2func = map[funcName]function{}
	func2name = map[uintptr]funcName{}

	fsVal := reflect.ValueOf(functions)
	fsType := reflect.TypeOf(functions)
	for i := 0; i < fsType.NumMethod(); i++ {
		m := fsType.Method(i)   // 仅仅用于获取函数名
		mVal := fsVal.Method(i) // 方法，用于远程调用
		mValType := mVal.Type() // 方法的类型，从中可以获取函数地址
		logger.Debug("register func", m.Name, fmt.Sprintf("%v", reflect_utils.InterfacePointer(mValType)))
		name2func[m.Name] = reflect.MakeFunc(mValType, func(args []reflect.Value) (results []reflect.Value) {
			logger.Debug("export func", m.Name, "called")
			return mVal.Call(args)
		}).Interface()
		func2name[reflect_utils.InterfacePointer(mValType) /*方法地址*/] = m.Name
	}
	return
}
