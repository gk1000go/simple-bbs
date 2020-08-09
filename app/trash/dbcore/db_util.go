package dbcore

import (
	"bytes"
	"fmt"
	"strings"
	"unsafe"
)

// 在不进行内存分配的情况下，将 []byte 转换为 string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
//
//// 在不进行内存分配的情况下，将 string 转换为 []byte
//func StringToBytes(str string) []byte {
//	s := (*[2]uintptr)(unsafe.Pointer(&str))
//	h := [3]uintptr{s[0], s[1], s[1]}
//	return *(*[]byte)(unsafe.Pointer(&h))
//}


type DbData struct {
	Key string
	Value interface{}
}
func (this DbData) getValue() string {
	a := bytes.Buffer{}
a.Bytes()
	return getValue(this.Value)
}

type OpData struct {
	Op string
	Value interface{}
}
func (t OpData)getSqlStr() string {
	if(t.Op == ""){ t.Op = "="}
	v1 := t.getValue()
	sqlStr := fmt.Sprintf("%s %s", t.Op, v1)
	switch t.Op {
	case "%": sqlStr = fmt.Sprintf("LIKE '%%%s%%'", v1[1:len(v1)-1])
	case "%-": sqlStr = fmt.Sprintf("LIKE '%%%s'", v1[1:len(v1)-1])
	case "-%": sqlStr = fmt.Sprintf("LIKE '%s%%'", v1[1:len(v1)-1])
	}
	return sqlStr
}
func (t OpData)getValue() string {
	return getValue(t.Value)
}

type DbDataWithOp struct {
	Action string
	Data map[string]interface{}
}

func getValue(data interface{}) string {
	dataFmt := ""
	switch data.(type) {
	case int:
		dataFmt = "%d"
	case float64:
		dataFmt = "%f"
	case string:
		dataFmt = "'%s'"
	default:
		return ""
	}
	return fmt.Sprintf(dataFmt, data)
}

func db_cond_to_sqladd(data []DbDataWithOp,sqlStr *bytes.Buffer)  {
	if(len(data) == 0) {return }
	//sqlStr := bytes.NewBufferString(" WHERE ")
	sqlStr.WriteString(" WHERE ")

	for _,val := range data{
		act := strings.ToUpper(val.Action)
		switch act {
		case "IN":
			sqlStr.WriteString("`")
			for k,v := range val.Data{
				sqlStr.WriteString(k)
				sqlStr.WriteString("` IN (")
				for _,subV := range v.([]interface{}){
					sqlStr.WriteString(getValue(subV))
					sqlStr.WriteString(",")
				}
				sqlStr.Truncate(sqlStr.Len()-1)
				sqlStr.WriteString(") AND `")
			}
			sqlStr.Truncate(sqlStr.Len()-6)
		case "OR":
			sqlStr.WriteString("(`")
			for k,v := range val.Data{
				vv := v.(OpData)
				sqlStr.WriteString(k)
				sqlStr.WriteString("` ")
				sqlStr.WriteString(vv.getSqlStr())
				sqlStr.WriteString(" OR `")
			}
			sqlStr.Truncate(sqlStr.Len() - 5)
			sqlStr.WriteString(")")
		case "AND":	fallthrough
		default:
			sqlStr.WriteString("`")
			for k,v := range val.Data{
				vv := v.(OpData)
				sqlStr.WriteString(k)
				sqlStr.WriteString("` ")
				sqlStr.WriteString(vv.getSqlStr())
				sqlStr.WriteString(" AND `")
			}
			sqlStr.Truncate(sqlStr.Len() - 6)
		}
		sqlStr.WriteString(" AND ")
	}
	sqlStr.Truncate(sqlStr.Len() - 5)
	return
}

func db_orderby_to_sqladd(data []DbData,sqlStr *bytes.Buffer)  {
	if(len(data) == 0) {return }
	//sqlStr := bytes.NewBufferString(" ORDER BY ")
	sqlStr.WriteString(" ORDER BY ")
	for _,val := range data{
		sqlStr.WriteString("`")
		sqlStr.WriteString(val.Key)
		sqlStr.WriteString("`")
		if(val.getValue() == "1"){
			sqlStr.WriteString(" ASC ")
		}else{
			sqlStr.WriteString(" DESC ")
		}
		sqlStr.WriteString(",")
	}
	sqlStr.Truncate(sqlStr.Len()-1)

	return
}
/*
	$arr = array(
		'name'=>'abc',
		'stocks+'=>1,
		'date'=>12345678900,
	)
	db_array_to_update_sqladd($arr);
*/
func db_array_to_update_sqladd(data []DbData,sqlStr *bytes.Buffer)  {
	if(len(data) == 0) {return }
	//sqlStr := bytes.NewBufferString("")

	for _,val := range data{
		sqlStr.WriteString("`")
		op := val.Key[len(val.Key)-1:]
		if(op == "+" || op == "-") {
			sqlStr.WriteString(val.Key[:len(val.Key)-1])
			sqlStr.WriteString("`=")
			sqlStr.WriteString(val.Key)
			sqlStr.WriteString(val.getValue())
		}else{
			sqlStr.WriteString(val.Key)
			sqlStr.WriteString("`=")
			sqlStr.WriteString(val.getValue())
		}
		sqlStr.WriteString(",")
	}
	sqlStr.Truncate(sqlStr.Len()-1)
	return
}

/*
	$arr = array(
		'name'=>'abc',
		'date'=>12345678900,
	)
	db_array_to_insert_sqladd($arr);
*/
func db_array_to_insert_sqladd(data []DbData,sqlStr *bytes.Buffer) {
	if(len(data) == 0) {return }
	//sqlStr := bytes.NewBufferString("(`")
	sqlStr.WriteString("(`")
	sqlVal := bytes.NewBufferString("")

	for _,val := range data{
		sqlStr.WriteString(val.Key)
		sqlStr.WriteString("`,")

		sqlVal.WriteString(val.getValue())
		sqlVal.WriteString(",")
	}
	sqlStr.Truncate(sqlStr.Len()-1)
	sqlStr.WriteString(") VALUES (")

	sqlVal.Truncate(sqlVal.Len()-1)
	sqlStr.Write(sqlVal.Bytes())
	sqlStr.WriteString(")")
	return
}
