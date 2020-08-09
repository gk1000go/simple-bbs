package dao

import (
	"awesomeProject1/app-web/framework"
	"fmt"
)

type UserDao struct {
	framework.BaseDAO
}

func NewUserDao() *UserDao {
	d := &UserDao{}
	d.Init()
	return d
}

func GetUserSingle(uid uint) (result framework.BaseDataMap) {
	sql := framework.GetSql("GetUserSingle", func(prefix string) string {
		return fmt.Sprintf("select * from %suser where uid = ?",prefix)
	})
	result = FetchResultDataSingleMap(sql,uid)
	return
}
//func (d *UserDao)GetUserOneSafe(uid uint) (result *model.User) {
//	sql := fmt.Sprintf("select uid,gid,username,Threads,Posts,Credits,Golds,Rmbs,Avatar from %suser where uid = ?",d.Prefix)
//	Fetch(func(rows *sqlx.Rows) error {
//		data := MkFetchResultSt(&model.User{})
//		if err := FetchData(rows,data);err==nil{
//			result = (data.Data).(*model.User)
//			return EndOfFetch
//		}else{
//			return err
//		}
//	},sql,uid)
//	return
//}


func (this *UserDao)user__create(data map[string]interface{}) int {
	// hook model_user__create_start.php
	// hook model_user__create_end.php

	return 0
}

func (this *UserDao)user__update(uid string,data map[string]interface{},d interface{}) int {
	// hook model_user__create_start.php
	// hook model_user__create_end.php

	return 0
}