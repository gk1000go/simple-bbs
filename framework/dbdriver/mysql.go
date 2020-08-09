package dbdriver

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func getMysqlError(err error) (*SqlErrInfo) {
	if err != nil{
		my := err.(*mysql.MySQLError)
		return &SqlErrInfo{my.Number,my.Message,my.Error()}
	}
	return nil
}

func ConnectToMysql(connStr string) *sqlx.DB {

	//db, err := gorm.Open("mysql", connStr) //"root:root@tcp(localhost:3406)/next9?parseTime=True"
	db := sqlx.MustConnect("mysql", connStr)


	log.Debug("connStr:",connStr)
	//if err != nil {
	//	panic("failed to connect database\n" + connStr)
	//}

	//db.SingularTable(true)
	//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	//	return "bbs_" + defaultTableName;
	//}
	return db
}