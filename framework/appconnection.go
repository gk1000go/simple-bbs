package framework

import (
	"awesomeProject1/app-web/framework/dbdriver"
	"github.com/jmoiron/sqlx"
)

const (
	DBTypeMysql = iota
)

type dbConn struct {
	WConn *sqlx.DB
	RConn []*sqlx.DB
}
var conn *dbConn



func GetConn() *sqlx.DB {
	return GetConnW()
}
func GetConnW() *sqlx.DB {
	if(conn.WConn == nil) {
		panic("No Database connection(W).")
	}
	return conn.WConn
}
func GetConnR() *sqlx.DB {
	if(len(conn.RConn) == 0) {
		conn.RConn = append(conn.RConn, GetConnW())
	}
	return conn.RConn[0] // todo ????
}

func connectToDB(dbType int,connStr string) {
	var db *sqlx.DB
	switch dbType {
	case DBTypeMysql:
		db = dbdriver.ConnectToMysql(connStr)
	default:

	}
	conn = &dbConn{WConn: db}
}
func closeAllConn() {
	if conn.WConn != nil{
		_ = conn.WConn.Close()
	}
	for _,db := range conn.RConn{
		_ = db.Close()
	}
}