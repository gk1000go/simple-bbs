package dbcore

import (
	"awesomeProject1/app-web/conf"
	"awesomeProject1/app-web/framework"
	"bytes"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DataParamOp struct {
	Op1 string // AND OR IN
	Op2 string // = <= ... LIKE
	Data map[string]interface{} // key=1 or key=[1,2,3,4]
}
type DbDataParam struct {
	Table string
	UpdData map[string]interface{}
	CondData []*DataParamOp
	OrderBy []string // col1,asc,col2,desc
}

type IDataAccess interface {
	Insert()
}

type DbCond interface {}
type DbOrderBy interface {}
type DbCols []string
type DbDataSet interface {}

type IDbCommon interface {
	version() string
	connect() bool
	close()
	//sql_find_one(param *DbParam) bool
	sql_find(param *DbParam) bool
	exec(param *DbParam)
	//count(param *DbParam)
	//max_id(param *DbParam)
	//truncate(table string) bool
	find(param *DbParam) bool
	//find_one(param *DbParam) bool
	get_tablepre() string
}

type DbParam struct {
	Table         string
	Cond          []DbDataWithOp
	OrderBy       *DbOrderBy
	Page          int
	PageSize      int
	Key           string
	Cols          []string
	ResultDataSet *DbDataSet
	SqlStr        *bytes.Buffer//bytes.NewBufferString
	ResultRows    *sqlx.Rows
	ResultOne     map[string]interface{}
	DbConn        IDbCommon
}
func (this *DbParam)Reset(){
	if(this.SqlStr == nil) {this.SqlStr = new(bytes.Buffer)}
	this.SqlStr.Reset()
}
func (this *DbParam)Clear(){
	if(this.ResultRows != nil){
		this.ResultRows.Close()
	}
}

func (this *DbParam)GetConn() IDbCommon {
	if(this.DbConn==nil){
		this.DbConn = NewMysql()
		this.DbConn.connect()
	}
	return this.DbConn
}

func db_connect(param *DbParam)  {
	d := param.GetConn()
	d.connect()
	//db_errno_errstr($r, $d);
}

func db_close(param *DbParam)  {
	d := param.GetConn()
	d.close()
	//db_errno_errstr($r, $d);
}

func db_sql_find_one(sqlStr string,param *DbParam)  {
	d := param.GetConn()
	d.find(param)
	//db_errno_errstr($arr, $d, $sql);
}

func db_sql_find(param *DbParam)  {
	d := param.GetConn()
	//sqlStr := ""
	d.sql_find(param)
	//db_errno_errstr($arr, $d, $sql);
}

// 如果为 INSERT 或者 REPLACE，则返回 mysql_insert_id();
// 如果为 UPDATE 或者 DELETE，则返回 mysql_affected_rows();
// 对于非自增的表，INSERT 后，返回的一直是 0
// 判断是否执行成功: mysql_exec() === FALSE
func db_exec(param *DbParam)  {
	d := param.GetConn()
	//xn_log(sqlStr, "db_exec")
	d.exec(param)
	//db_errno_errstr($n, $d, $sql);
}

func db_count(param *DbParam)  {
	//d := param.GetConn()
	//d.count(param)
}

func db_maxid(param *DbParam)  {
	//d := param.GetConn()
	//d.maxid(param)
	//db_errno_errstr
}

func db_create(param *DbParam) {
	db_insert(param);
}

func db_insert(param *DbParam) int {
	//d := param.GetConn()
	//sqlInsert := ""//db_array_to_insert_sqladd($arr);
	//sqlStr := fmt.Sprintf("INSERT INTO %s%s %s",d.get_tablepre(),param.Table,sqlInsert)
	//db_exec(sqlStr,param)
	return 0
}

func db_replace(param *DbParam){
	//d := param.GetConn()
	//sqlReplace := ""//db_array_to_insert_sqladd($arr);
	//sqlStr := fmt.Sprintf("REPLACE INTO %s%s %s",d.get_tablepre(),param.Table,sqlReplace)
	//db_exec(sqlStr,param)
}

func Db_update(param *DbParam) int {
	//d := param.GetConn()
	//sqlWhere := ""//db_cond_to_sqladd($cond);
	//sqlUpdate := ""//db_array_to_update_sqladd($update);
	//sqlStr := fmt.Sprintf("UPDATE %s%s SET %s %s",d.get_tablepre(),param.Table,sqlUpdate,sqlWhere)
	//db_exec(sqlStr,param)
	return 0
}

func db_delete(param *DbParam) {
	//d := param.GetConn()
	//sqlWhere := ""//db_cond_to_sqladd($cond);
	//sqlStr := fmt.Sprintf("DELETE FROM %s%s %s",d.get_tablepre(),param.Table,sqlWhere)
	//db_exec(sqlStr, param);
}

func Db_truncate(param *DbParam)  {
	//d := param.GetConn()
	//return d.truncate(fmt.Sprintf("%s%s",d.get_tablepre(),param.Table))
}

func Db_read(param *DbParam)  {
	//d := param.GetConn()
	//sqlWhere := ""//db_cond_to_sqladd($cond);
	//sqlStr := fmt.Sprintf("SELECT * FROM %s%s %s",d.get_tablepre(),param.Table,sqlWhere)
	// db_sql_find_one(sqlStr,param)
}

func db_find(param *DbParam)  {
	//d := param.GetConn()
	//d.find(param)
	//return true
}

func db_find_one(param *DbParam)  {
	//d := param.GetConn()
	//return d.find_one(param)
}

func db_errno_errstr(r bool,d *DBMysql,sSql string){
	if(!r) { //  && $d->errno != 0
		c := framework.GetConn()
		c.ErrNo = d.Errno
		c.ErrStr = db_errstr_safe(d.Errno, d.ErrStr);

		str := fmt.Sprintf("SQL:%s\r\nerrno: %d, errstr: %s",sSql,c.ErrNo,c.ErrStr)
		xn_log(str,"db_error")
	}
}
func db_errstr_safe(errno int,errStr string) string {
	if(conf.DefaultConfig.DEBUG>0) {return errStr}
	if(errno == 1049) {
		return "数据库名不存在，请手工创建"
	} else if(errno == 2003 ) {
		return "连接数据库服务器失败，请检查IP是否正确，或者防火墙设置"
	} else if(errno == 1024) {
		return "连接数据库失败"
	} else if(errno == 1045) {
		return "数据库账户密码错误"
	}
	return errStr
}