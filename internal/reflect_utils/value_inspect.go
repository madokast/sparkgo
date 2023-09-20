package reflect_utils

import (
	"reflect"
	"unsafe"
)

type valueType struct {
	typ_ uintptr
	ptr  uintptr
	flag uintptr
}

func ValueType(v reflect.Value) uintptr {
	return (*valueType)(unsafe.Pointer(&v)).typ_
}

func ValuePtr(v reflect.Value) uintptr {
	return (*valueType)(unsafe.Pointer(&v)).ptr
}
