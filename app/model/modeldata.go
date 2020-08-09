package model

import (
	"reflect"
	"unsafe"
)

type DataSliceType []interface{}

func (d DataSliceType)GetInt(i int) int {
	if d[i] == nil {return 0}
	return d[i].(int)
}
func (d DataSliceType)GetStr(i int) string {
	if d[i] == nil {return ""}
	x := d[i].([]byte)
	s := *(*string)(unsafe.Pointer(&x))
	return s
}
type DataMapType map[string]interface{}
func (d DataMapType)GetInt(k string) int {
	if v,ok := d[k];ok {
		if v == nil {return 0}
		return v.(int)
	}else{
		return 0
	}
}
func (d DataMapType)GetStr(k string) string {
	if v,ok := d[k];ok {
		if v == nil {return ""}
		x := v.([]byte)
		s := *(*string)(unsafe.Pointer(&x))
		return s
	}else{
		return ""
	}
}

//orcaman
// concurrent-map
type ModelData struct {
	hasData bool
	RawDataKind reflect.Kind
	LastEffectRows int
	DataSlice DataSliceType
	DataMap DataMapType
	DataStruct interface{}
	Cols []string
}

func (m *ModelData)Get(key string) interface{} {
	if m.RawDataKind == reflect.Map{
		return m.DataMap[key]
	}
	return nil
}
func (m *ModelData)GetStruct(pointer interface{}) interface{} {
	if m.RawDataKind == reflect.Struct{
		return m.DataStruct
	}
	if m.DataStruct != nil {
		return m.DataStruct
	}

	return nil
}
func (m *ModelData)GetSlice() DataSliceType {
	if m.RawDataKind == reflect.Struct{
		return m.DataSlice
	}
	return nil
}
func (m *ModelData)GetMap(pointer interface{}) DataMapType {
	if m.RawDataKind == reflect.Map{
		return m.DataMap
	}
	return nil
}




func Byte2Str(in interface{})(s string){
	x := in.([]byte)
	s = *(*string)(unsafe.Pointer(&x))
	return
}

func Byte2Str2(in interface{})(s string){
	x := in.([]byte)
	s = string(x)
	return
}