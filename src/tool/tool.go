package tool

import "strings"

//高效拼接字符串
var sBuild strings.Builder

func Splicing(strs ...string) string {
	sBuild.Reset()
	len := len(strs)
	for i := 0; i < len; i++ {
		sBuild.WriteString(strs[i])
	}
	return sBuild.String()
}

//查询数组是否含有
func FindIndex[T string | int](originArray []T, flag T) int {

	if originArray == nil {
		return -1
	}

	len := len(originArray)
	for i := 0; i < len; i++ {
		if originArray[i] == flag {
			return i
		}
	}
	return -1
}

//三元运算
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
