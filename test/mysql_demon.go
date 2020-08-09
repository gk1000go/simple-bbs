package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)


type Xyz struct {
	X int `db:"X"`
}
func main()  {
	defer func() {
		// panicがおきたらロールバック
		if err := recover(); err != nil {
			log.Println("defer:")
			log.Println(err)
		}
	}()
	connStr := "root:" + "root" + "@tcp(" + "localhost" + ":" + "3306" + ")/" + "rechannel_db0" + "?charset=" + "utf8"
	hDB, err := sqlx.Open("mysql", connStr)
	if err != nil {
		ShowSqlError(err)
		return
	}
	defer hDB.Close()
	//hDB1, err := sqlx.Open("mysql", connStr)
	//defer hDB1.Close()
	//hDB2, err := sqlx.Open("mysql", connStr)
	//defer hDB2.Close()
	//hDB.SetMaxOpenConns(10)
	//c,_ :=hDB.Conn(context.Background())
	//smtp,_ := c.PrepareContext(context.Background(),"")
	//smtp.
	log.Printf("max:=%d,open:=%d,use=%d\n",hDB.Stats().MaxOpenConnections,hDB.Stats().OpenConnections,hDB.Stats().InUse)

	//rows := hDB.QueryRow("select name from bbs_group")
//	rows,err := hDB.Queryx("SET names utf8, sql_mode=''")
	rows,err := hDB.Queryx("select 1 from dual")
	fmt.Println(err)
	for rows.Next() {
		v,err := rows.SliceScan()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("1-1,%#v\n", v)
	}
	rows,err = hDB.Queryx("select 21 from dual union all select 22 from dual")
	//if rows.Next() {
	//	v,err := rows.SliceScan()
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	fmt.Printf("1-2,%#v\n", v)
	//}
	//rows.Close() //rows 取完毕或Close才释放Conn
	//hDB.Select()
	//result := make(map[string]interface{})
	//if err = rows.MapScan(result);err!=nil{
	//	log.Println(err.Error())
	//}
	//log.Println(result)

	dog := Xyz{}
	err = hDB.Get(&dog,"select 21 AS X from dual union all select 22  AS X from dual")
	if (err != nil) {
		fmt.Printf("dog-e-%#v\n", err.Error())
	}
	fmt.Printf("dog-%#v\n", dog)

	rows,err = hDB.Queryx("select 3 from dual")
	rows,err = hDB.Queryx("select 4 from dual")
	rows,err = hDB.Queryx("select 5 from dual")

	s := "create table `my_table`(xid smallint(6) unsigned NOT NULL,PRIMARY KEY (xid))"

	r,e := hDB.Exec(s)
	if(e != nil){
		fmt.Println(e)
	}
	fmt.Println(r)

	log.Printf("max:=%d,open:=%d,use=%d\n",hDB.Stats().MaxOpenConnections,hDB.Stats().OpenConnections,hDB.Stats().InUse)
	////if err != nil {
	////	ShowSqlError(err)
	////}
	//var name sql.NullString
	////rows.Next()
	//rows.Scan(&name)
	//log.Println(reflect.TypeOf(name))
	//log.Println(name.Valid)
	////defer rows.Close()

}

func ShowSqlError(err error)  {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		log.Println("ShowSqlError:")
		log.Println(mysqlErr.Number)
		log.Println(mysqlErr.Message)
		log.Println(mysqlErr.Error())
	}else{
		log.Println("ShowSqlError:convert error!!!!")
	}
}

/*
go get github.com/jmoiron/sqlx


id := 100
var name string

if err := cnn.QueryRow("SELECT name FROM person WHERE id = ?LIMIT 1", id).Scan(&name); err != nil {
    log.Fatal(err)
}


rows, err := cnn.Query("SELECT id, name FROM person")
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
    var id int
    var name string
    if err := rows.Scan(&id, &name); err != nil {
        log.Fatal(err)
    }
    fmt.Println(id, name)
}

if err := rows.Err(); err != nil {
    log.Fatal(err)
}

--update
type Result interface {
        LastInsertId() (int64, error)
        RowsAffected() (int64, error)
}

result, err := cnn.Exec("UPDATE person SET name = ? WHERE id = ?", "Hogera", 100)
if err != nil {
    log.Fatal(err)
}

--トランザクションを使ってみる
tx, err := cnn.Begin()
if err != nil {
    log.Fatal(err)
}
tx.Commit()


--ロールバック
func hoge(tx *sql.Tx) {
    defer func() {
        // panicがおきたらロールバック
        if err := recover(); err != nil {
            tx.Rollback()
        }
    }()
    // …
    tx.Exec(…)
}

--pool
var cnnPool chan *sql.DB

func main() {
    cnnPool = make(chan *sql.DB, 10)
    for i := 0; i < cap(cnnPool); i++ {
        cnnPool <- sql.Open("mysql", "user:password@tpc(host:port)/dbname")
    }
}

func getFromDB() {
    cnn := <-cnnPool
    defer func() {
        cnnPool <- cnn
    }()
    // cnnを使った処理
}
 */