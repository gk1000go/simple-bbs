package trash

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
)
type DbCond interface {}
type DbOrderBy interface {}
type DbCols interface {}
type DbDataSet interface {}

type IDbCommon interface {
	//sql_find_one(sSql string) bool
	//sql_find(sSql string,key []string) DbDataSet
	find(table string,cond DbCond,orderBy DbOrderBy,page int,pageSize int,key string,col DbCols) DbDataSet
	find_one(table string,cond DbCond,orderBy DbOrderBy,col DbCols) bool
}

type baseDBConn struct {
	DbLink *sql.DB
	ErrNo int
	ErrStr string
}
//var baseDBConnection *baseDBConn
//func GetConn() *baseDBConn {
//	if (baseDBConnection ==nil){
//		baseDBConnection = &baseDBConn{}
//	}
//	return baseDBConnection
//}

func (this *baseDBConn) find(table string,cond DbCond,orderBy DbOrderBy,page int,pageSize int,key string,col DbCols) DbDataSet {
	return nil
}
func (this *baseDBConn) find_one(table string,cond DbCond,orderBy DbOrderBy,col DbCols) bool {
	return false
}

type IDbFunc interface {
	Db_connect()
	Db_close()
	Db_sql_find_one()
	Db_sql_find()
	Db_exec()
	Db_count()
	Db_maxid()
	Db_create()
	Db_insert()
	Db_replace()
	Db_update()
	Db_delete()
	Db_truncate()
	Db_read()
	Db_find()
	Db_find_one()
	//Db_errno_errstr()
	//Db_errstr_safe()
	//Db_cond_to_sqladd()
	//Db_orderby_to_sqladd()
	//Db_array_to_update_sqladd()
	//Db_array_to_insert_sqladd()
}




type DbData struct {
	Key string
	Value interface{}
}
type DbDataWithOp struct {
	Action string
	Ops []string
	Keys interface{}
	Values []interface{}
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
	dataFmt := ""
	switch t.Value.(type) {
	case int:
		dataFmt = "%d"
	case float64:
		dataFmt = "%f"
	case string:
		dataFmt = "'%s'"
	default:
		return ""
	}
	return fmt.Sprintf(dataFmt, t.Value)
}

type DbDataWithOpEx struct {
	Action string
	Data map[string]interface{}
}

func (this DbData)GetValue() string {
	return getValue(this.Value)
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

func db_cond_to_sqladd_ex(data []DbDataWithOpEx) string {
	if(len(data) == 0) {return ""}
	sqlStr := bytes.NewBufferString(" WHERE ")

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
	return sqlStr.String()
}

/*
$cond = array('id'=>123, 'groupid'=>array('>'=>100, 'LIKE'=>'\'jack'));
$s = db_cond_to_sqladd($cond);
echo $s;

WHERE id=123 AND groupid>100 AND groupid LIKE '%\'jack%'

// 格式：
array('id'=>123, 'groupid'=>123)
array('id'=>array(1,2,3,4,5))                    -- IN
array('id'=>array('>' => 100, '<' => 200))       -- AND
array('||'=>array('aa>' => 100, 'bb<=' => 200))  -- OR
array('username'=>array('LIKE' => 'jack'))
*/
func db_cond_to_sqladd(data []DbDataWithOp) string {
	if(len(data) == 0) {return ""}
	sqlStr := bytes.NewBufferString(" WHERE ")

	for _,val := range data{
		act := strings.ToUpper(val.Action)
		switch act {
		case "IN":
			sqlStr.WriteString("`")
			sqlStr.WriteString(val.Keys.(string))
			sqlStr.WriteString("` IN (")
			//sqlStr.WriteString(getValues(val.Values))
			for _,inVal := range val.Values{
				sqlStr.WriteString(getValue(inVal))
				sqlStr.WriteString(",")
			}
			sqlStr.Truncate(sqlStr.Len()-1)
			sqlStr.WriteString(")")
		case "AND":
			sqlStr.WriteString("`")
			keys := val.Keys.([]string)
			for idx,key := range keys{ //val.Ops
				sqlStr.WriteString(key)
				sqlStr.WriteString("` ")
				if(val.Ops == nil || len(val.Ops)==0){
					sqlStr.WriteString("=")
				}else{
					sqlStr.WriteString(val.Ops[idx])
				}
				sqlStr.WriteString(" ")
				sqlStr.WriteString(getValue(val.Values[idx]))
				sqlStr.WriteString(" AND `")
			}
			sqlStr.Truncate(sqlStr.Len() - 6)
		case "OR":
			sqlStr.WriteString("(`")
			keys := val.Keys.([]string)
			for idx,key := range keys{
				sqlStr.WriteString(key)
				sqlStr.WriteString("` ")
				if(val.Ops == nil || len(val.Ops)==0){
					sqlStr.WriteString("=")
				}else{
					sqlStr.WriteString(val.Ops[idx])
				}
				sqlStr.WriteString(" ")
				sqlStr.WriteString(getValue(val.Values[idx]))
				sqlStr.WriteString(" OR `")
			}
			sqlStr.Truncate(sqlStr.Len() - 5)
			sqlStr.WriteString(")")
		default:
			sqlStr.WriteString("`")
			keys := val.Keys.([]string)
			for idx,key := range keys{
				sqlStr.WriteString(key)
				sqlStr.WriteString("`")
				sqlStr.WriteString(" = ")
				sqlStr.WriteString(getValue(val.Values[idx]))
				sqlStr.WriteString(" AND `")
			}
			sqlStr.Truncate(sqlStr.Len() - 6)
		}
		sqlStr.WriteString(" AND ")
	}

	sqlStr.Truncate(sqlStr.Len()-4)
	return sqlStr.String()
}

func Db_orderby_to_sqladd(data []DbData) string {
	if(len(data) == 0) {return ""}
	sqlStr := bytes.NewBufferString(" ORDER BY ")
	for _,val := range data{
		sqlStr.WriteString("`")
		sqlStr.WriteString(val.Key)
		sqlStr.WriteString("`")
		if(val.GetValue() == "1"){
			sqlStr.WriteString(" ASC ")
		}else{
			sqlStr.WriteString(" DESC ")
		}
		sqlStr.WriteString(",")
	}
	sqlStr.Truncate(sqlStr.Len()-1)

	return sqlStr.String()
}
/*
	$arr = array(
		'name'=>'abc',
		'stocks+'=>1,
		'date'=>12345678900,
	)
	db_array_to_update_sqladd($arr);
*/
func Db_array_to_update_sqladd(data []DbData) string {
	if(len(data) == 0) {return ""}
	sqlStr := bytes.NewBufferString("")

	for _,val := range data{
		sqlStr.WriteString("`")
		op := val.Key[len(val.Key)-1:]
		if(op == "+" || op == "-") {
			sqlStr.WriteString(val.Key[:len(val.Key)-1])
			sqlStr.WriteString("`=")
			sqlStr.WriteString(val.Key)
			sqlStr.WriteString(val.GetValue())
		}else{
			sqlStr.WriteString(val.Key)
			sqlStr.WriteString("`=")
			sqlStr.WriteString(val.GetValue())
		}
		sqlStr.WriteString(",")
	}
	sqlStr.Truncate(sqlStr.Len()-1)

	return sqlStr.String()
}

/*
	$arr = array(
		'name'=>'abc',
		'date'=>12345678900,
	)
	db_array_to_insert_sqladd($arr);
*/
func Db_array_to_insert_sqladd(data []DbData) string {
	if(len(data) == 0) {return ""}
	sqlStr := bytes.NewBufferString("(`")
	sqlVal := bytes.NewBufferString("")

	for _,val := range data{
		sqlStr.WriteString(val.Key)
		sqlStr.WriteString("`,")

		sqlVal.WriteString(val.GetValue())
		sqlVal.WriteString(",")
	}
	sqlStr.Truncate(sqlStr.Len()-1)
	sqlStr.WriteString(") VALUES (")

	sqlVal.Truncate(sqlVal.Len()-1)
	sqlStr.Write(sqlVal.Bytes())
	sqlStr.WriteString(")")

	return sqlStr.String()
}





