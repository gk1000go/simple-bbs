package framework

import (
	"awesomeProject1/app-web/framework/dbdriver"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sync"
)

type BaseDAO struct {
	Db *sqlx.DB
	ErrInfo *dbdriver.SqlErrInfo
	Prefix string
}

const Prefix = "bbs_"

func (d *BaseDAO)Init() {
	d.Db = GetConn()
	d.Prefix = "bbs_"
}
func (d *BaseDAO)T(t string) string {
	return fmt.Sprintf("%s%s",d.Prefix,t)
}

//func (d *BaseDAO)QueryOne(query string, args ...interface{}) *model.ModelData {
//	md := model.New()
//
//}

func (d *BaseDAO) GetError(e error) bool {
	d.ErrInfo = dbdriver.GetSqlError(e)

	return e!=nil
}

type FuncNewSql func (prefix string) string
var sqlStorage sync.Map
func GetSql(key string,f FuncNewSql) string {
	str,ok := sqlStorage.Load(key);
	if !ok{
		str := f(Prefix)
		sqlStorage.Store(key,str)
	}
	return str.(string)
}











type abcd map[string]interface{}

func (m abcd)do() {
	m["a1"] = 345
	fmt.Println("no key:",m["b1"])
}

func test()  {
	a := map[string]interface{}{
		"a1":123,
		"a2":"asd",
	}
	b := abcd(a)
	b.do()
	fmt.Println(b)
}