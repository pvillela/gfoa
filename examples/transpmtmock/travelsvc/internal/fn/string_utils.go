package fn

import (
	"fmt"
	"strings"
)

// TypeStr returns the simple type name of its argument as a string.
func TypeStr(x interface{}) string {
	qualifiedType := fmt.Sprintf("%T", x)
	dotIdx := strings.LastIndexByte(qualifiedType, '.')
	return qualifiedType[dotIdx+1:]
}

// StrApply returns a string that combines its arguments to help with the mock implementatios
// of the functions in the example application.
func StrApply(x interface{}, cfg string, args ...string) string {
	argStr := strings.Join(args, ", ")
	return fmt.Sprintf("%v[%v](%v)", TypeStr(x), cfg, argStr)
}
