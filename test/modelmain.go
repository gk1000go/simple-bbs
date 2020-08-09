package main

import (
	"awesomeProject1/app-web/app/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type Group struct {
	Gid             uint
	Name            string // 用户组名称
	Creditsfrom     int    // 积分从
	Creditsto       int    // 积分到
	Allowread       int    // 允许访问
	Allowthread     int    // 允许发主题
	Allowpost       int    // 允许回帖
	Allowattach     int    // 允许上传文件
	Allowdown       int    // 允许下载文件
	Allowtop        int    // 允许置顶
	Allowupdate     int    // 允许编辑
	Allowdelete     int    // 允许删除
	Allowmove       int    // 允许移动
	Allowbanuser    int    // 允许禁止用户
	Allowdeleteuser int    // 允许删除用户
	Allowviewip     uint   // 允许查看用户敏感信息
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt *time.Time `sql:"index"`
}


func main()  {
	//charset=utf8&loc=Asia%2FShanghai
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3406)/next9?parseTime=True")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "bbs_" + defaultTableName;
	}

	//gorm.Model{}

	testuser11(db)

}
func testuser11(db *gorm.DB)  {
	s := &Group{}
	log.Println(s)
	db.Where("gid=?",104).First(s)
	log.Println(s)
	//var user []model.Group
	//user := make(model.Group)

}

func testuser(db *gorm.DB)  {
	for _,s := range model.All {
		log.Println(s)
		db.First(s)
		log.Println(s)
	}
	//var user []model.Group
	//user := make(model.Group)

}

//func group(db *gorm.DB)  {
//	log.Println("start group.")
//	db.AutoMigrate(&model.Group{})
//	for _, tt := range model.InitGroup() {
//		db.Create(&tt)
//		if db.Error != nil {
//			log.Println(db.Error)
//		}
//		//db.Update(&tt)
//	}
//}
//func user(db *gorm.DB)  {
//	log.Println("start User.")
//	db.AutoMigrate(&model.User{})
//	if db.Error != nil {
//		log.Println(db.Error)
//	}
//	for _, tt := range model.InitUser() {
//		db.Create(&tt)
//		if db.Error != nil {
//			log.Println(db.Error)
//		}
//		//db.Update(&tt)
//	}
//}
//
//func forum(db *gorm.DB)  {
//	log.Println("start forum.")
//	db.AutoMigrate(&model.Forum{})
//	if db.Error != nil {
//		log.Println(db.Error)
//	}
//	for _, tt := range model.InitForum() {
//		db.Create(&tt)
//		if db.Error != nil {
//			log.Println(db.Error)
//		}
//		//db.Update(&tt)
//	}
//}