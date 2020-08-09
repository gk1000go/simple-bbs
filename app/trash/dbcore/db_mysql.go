package dbcore

import (
	"awesomeProject1/app-web/aputils"
	"awesomeProject1/app-web/conf"
	"bytes"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type DBMysql struct {
	ConfMaster   *conf.DbConf   // 配置，可以支持主从
	ConfSlaves   []*conf.DbConf // 配置，可以支持主从
	WLink        *sqlx.DB        // 写连接
	RLinks       []*sqlx.DB        // 读连接
	Link         *sqlx.DB        // 最后一次使用的连接
	Errno        int
	ErrStr       string
	Sqls         interface{}
	Tablepre     string
	Innodb_first bool					// 优先 InnoDB
	MyErr 		 *mysql.MySQLError
}

var conn *DBMysql

func NewMysql() IDbCommon {
	if(conn != nil){
		return conn
	}
	cfg := conf.DefaultConfig
	conn = &DBMysql{}
	conn.ConfMaster = &cfg.BaseDb.Master
	conn.Tablepre = cfg.BaseDb.Master.TablePre
	if (len(cfg.BaseDb.Slaves)>0){
		for i,_ := range cfg.BaseDb.Slaves {
			conn.ConfSlaves = append(conn.ConfSlaves,&cfg.BaseDb.Slaves[i])
		}
	}
	return conn
}
func (this *DBMysql) version() string {
	return "8.0.21"
}
// 根据配置文件连接
func (this *DBMysql) connect() bool {
	this.connect_master()
	this.connect_slaves()
	return (this.WLink != nil && len(this.RLinks) == 0)
}

func (this *DBMysql) close() {
	_ = this.WLink.Close()
	for _,d := range this.RLinks {
		_ = d.Close()
	}
}

// 连接写服务器
func (this *DBMysql) connect_master() {
	log.Debug("mysql connect_master")
	if (this.WLink != nil) {return}
	cfg := this.ConfMaster
	if(this.WLink == nil) {
		this.WLink = this.real_connect(cfg.Host,cfg.Port,cfg.User,cfg.Password,cfg.Name,cfg.Charset,cfg.Engine)
		log.WithFields(aputils.StructToMap(cfg)).Info("mysql master is connected.")
	}
}

// 连接从服务器，如果有多台，则随机挑选一台，如果为空，则与主服务器一致。
func (this *DBMysql) connect_slaves() {
	log.Debug("mysql connect_slaves")
	if (this.RLinks != nil && len(this.RLinks)>0) {return}
	if(len(this.ConfSlaves) == 0) {
		if (this.WLink == nil) {
			this.connect_master()
		}
		this.RLinks = append(this.RLinks,this.WLink)
		this.ConfSlaves = append(this.ConfSlaves,this.ConfMaster)
	}else{

		for i, cfg := range this.ConfSlaves{
			if d := this.real_connect(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.Charset, cfg.Engine);d!=nil{
				this.RLinks = append(this.RLinks,d)
				log.WithFields(aputils.StructToMap(cfg)).Infof("mysql [%d]slave is connected.",i)
			}
		}
	}
}

/*
Open no pint
Connect with ping
 */
func (this *DBMysql) real_connect(host string,port string,user string,password string,dbname string,charset string,engine string) *sqlx.DB {
	log.Debug("real_connect")
	//$link = @mysql_connect($host, $user, $password); // 如果用户名相同，则返回同一个连接。 fastcgi 持久连接更省资源
	connStr := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=" + charset
	link, err := sqlx.Connect("mysql", connStr)

	if (this.hasError(err)) {//panic(err.Error())
		log.WithFields(this.printError()).Warn("mysql connect fail...")
		return nil
	}
	buf := bytes.Buffer{}
	buf.WriteString("SET names "+charset+", sql_mode=''")
	this.do_query(&buf, link);
	return link
}

func (this *DBMysql) selectRLink(){
	rand.Seed(time.Now().UnixNano())
	this.Link = this.RLinks[rand.Intn(len(this.RLinks))]
}

func (this *DBMysql) do_query(buf *bytes.Buffer,link *sqlx.DB) (*sqlx.Rows,bool){
	if(link==nil) {this.selectRLink()}
	if(this.Link == nil){return nil,false}

	// make SQL
	sqlStr := BytesToString(buf.Bytes())

	startTime := time.Now().UnixNano()
	rows,err := link.Queryx(sqlStr)
	endTime := time.Now().UnixNano()
	if(this.hasError(err)){
		log.WithFields(this.printError()).Warn("mysql do_query fail...")
		return nil,false
	}
	log.WithFields(log.Fields{
		"sql":sqlStr,
		"time":float64((endTime - startTime) / 1e6),
	}).Debug("mysql do_query time.")

	return rows,true
}

func (this *DBMysql) do_exec(buf *bytes.Buffer,link *sqlx.DB) (*sql.Result,bool){
	if(link == nil) {
		link = this.WLink
	}
	sqlStr := BytesToString(buf.Bytes())
	startTime := time.Now().UnixNano()
	result,err := link.Exec(sqlStr)
	endTime := time.Now().UnixNano()
	if(this.hasError(err)){
		log.WithFields(this.printError()).Warn("mysql do_exec fail...")
		return nil,false
	}
	log.WithFields(log.Fields{
		"sql":sqlStr,
		"time":float64((endTime - startTime) / 1e6),
	}).Debug("mysql do_exec time.")

	return &result,true
}

func (this *DBMysql)sql_find_one(param *DbParam) (ok bool) {
	if(param.SqlStr.Len() == 0) {return false}
	buf := param.SqlStr
	buf.WriteString(" LIMIT 1")
	return
}

func (this *DBMysql)sql_find(param *DbParam) (ok bool) {
	if(param.SqlStr.Len() == 0) {return false}
	param.ResultRows,ok = this.do_query(param.SqlStr,nil)
	return
}

func (this *DBMysql)count(param *DbParam)  {
	buf := param.SqlStr
	buf.Reset()
	buf.WriteString("SELECT COUNT(*) AS NUM FROM `")
	buf.WriteString(param.Table)
	buf.WriteString("`")
	db_cond_to_sqladd(param.Cond,buf)

}

func (this *DBMysql)max_id(param *DbParam)  {
	buf := param.SqlStr
	buf.Reset()
	buf.WriteString("SELECT MAX(`")
	for _,col := range param.Cols{
		buf.WriteString(col)
		buf.WriteString("`) AS ")
		buf.WriteString(col)
	}
	buf.WriteString(" FROM `")
	buf.WriteString(param.Table)
	buf.WriteString("`")
	db_cond_to_sqladd(param.Cond,buf)

}
func (this *DBMysql)find_one(param *DbParam) bool {
	buf := param.SqlStr
	buf.Reset()
	buf.WriteString("SELECT COUNT(*) AS NUM FROM `")
	buf.WriteString(param.Table)
	buf.WriteString("`")
	db_cond_to_sqladd(param.Cond,buf)
	buf.WriteString(" LIMIT 1")
	//	if rows,ok := this.query(param.SqlStr,nil);ok{
	//
	//		param.ResultRows
	//	}
	return false
}
//func (this *DBMysql)truncate(table string) bool {
//	return false
//}

func (this *DBMysql)find(param *DbParam) bool {
	return false
}

func (this *DBMysql) exec(param *DbParam)  {

}

func (this *DBMysql) get_tablepre() string {
	return this.Tablepre
}


func (this *DBMysql)hasError(err error) (ok bool) {
	this.MyErr,ok = err.(*mysql.MySQLError)
	return
}
func (this *DBMysql)printError() (e map[string]interface{}) {
	return map[string]interface{}{
		"Errno":this.MyErr.Number,
		"ErrMsg":this.MyErr.Message,
		"Error":this.MyErr.Error(),
	}
}

func ConvMyErr(err error)(int,string,string,bool){

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			return int(mysqlErr.Number),mysqlErr.Message,mysqlErr.Error(),true
		}
	}
	return 0,"","",false
}