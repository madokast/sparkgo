package reflect_utils

import "unsafe"

type interfaceType struct {
	typ  uintptr
	word uintptr
}

func InterfacePointer(val any) uintptr {
	return ((*interfaceType)(unsafe.Pointer(&val))).word
}

func InterfaceType(val any) uintptr {
	return ((*interfaceType)(unsafe.Pointer(&val))).typ
}
