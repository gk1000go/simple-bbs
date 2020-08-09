package aputils

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

///https://github.com/orcaman/concurrent-map/blob/master/concurrent_map.go

// check interface is nil
func IsNil(x interface{}) bool {
	return ( (x == nil) || reflect.ValueOf(x).IsNil() )
}

func File_line(skp int) string {
	_, fileName, fileLine, ok := runtime.Caller(skp) // skip File_line, and chkError
	var s string
	if ok {
		s = fmt.Sprintf("%s:%d", fileName, fileLine)
	} else {
		s = ""
	}
	return s
}

func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		result[field] = value
	}

	return result
}
func MapToStruct(mapVal map[string]interface{}, val interface{}) (ok bool) {
	structVal:= reflect.Indirect(reflect.ValueOf(val))
	for name, elem := range mapVal {
		structVal.FieldByName(name).Set(reflect.ValueOf(elem))
	}
	return
}
func GetStr(x []byte) string {
	if x == nil || len(x) == 0 {return ""}
	s := *(*string)(unsafe.Pointer(&x))
	return s
}

func dotime()  {
	//startTime := time.Now().UnixNano()
	//
	///* 程序主体 */
	//
	//endTime := time.Now().UnixNano()

	//seconds:= float64((endTime - startTime) / 1e9)
	//Milliseconds:= float64((endTime - startTime) / 1e6)
	//nanoSeconds:= float64(endTime - startTime)

}

func strip(s_ string, chars_ string) string {
	s, chars := []rune(s_), []rune(chars_)
	length := len(s)
	max := len(s) - 1
	l, r := true, true //标记当左端或者右端找到正常字符后就停止继续寻找
	start, end := 0, max
	tmpEnd := 0
	charset := make(map[rune]bool) //创建字符集，也就是唯一的字符，方便后面判断是否存在
	for i := 0; i < len(chars); i++ {
		charset[chars[i]] = true
	}
	for i := 0; i < length; i++ {
		if _, exist := charset[s[i]]; l && !exist {
			start = i
			l = false
		}
		tmpEnd = max - i
		if _, exist := charset[s[tmpEnd]]; r && !exist {
			end = tmpEnd
			r = false
		}
		if !l && !r {
			break
		}
	}
	if l && r { // 如果左端和右端都没找到正常字符，那么表示该字符串没有正常字符
		return ""
	}
	return string(s[start : end+1])

}