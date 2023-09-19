package reflect_utils

import (
	"fmt"
	"reflect"
	"strings"
)

type function = any

func FuncString(f function) string {
	t := reflect.TypeOf(f)
	k := t.Kind()
	if k != reflect.Func {
		return fmt.Sprintf("Non-Func(%s)", k.String())
	}
	sb := strings.Builder{}
	sb.WriteString("func(")
	numIn := t.NumIn()
	for i := 0; i < numIn; i++ {
		sb.WriteString(t.In(i).Kind().String())
		if i < numIn-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(") ")
	numOut := t.NumOut()
	for i := 0; i < numOut; i++ {
		sb.WriteString(t.Out(i).Kind().String())
		if i < numOut-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}

func Call(f function, params []any) []any {
	v := reflect.ValueOf(f)
	
	v.Call()
}