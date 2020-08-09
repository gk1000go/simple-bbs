package dao

import (
	"awesomeProject1/app-web/app/model"
	"awesomeProject1/app-web/framework"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type HomeDAO struct {
	framework.BaseDAO
}

func NewHomeDao() *HomeDAO {
	h := &HomeDAO{}
	h.Init()
	return h
}

func (d *HomeDAO)GetUserOne(uid uint) (result *model.User) {
	sql := fmt.Sprintf("select * from %suser where uid = ?",d.Prefix)
 	Fetch(func(rows *sqlx.Rows) error {
		data := MkFetchResultSt(&model.User{})
		if err := FetchData(rows,data);err==nil{
			result = (data.Data).(*model.User)
			return EndOfFetch
		}else{
			return err
		}
	},sql,uid)
	return
}
func (d *HomeDAO)GetUserOneSafe(uid uint) (result *model.User) {
	sql := fmt.Sprintf("select uid,gid,username,Threads,Posts,Credits,Golds,Rmbs,Avatar from %suser where uid = ?",d.Prefix)
	Fetch(func(rows *sqlx.Rows) error {
		data := MkFetchResultSt(&model.User{})
		if err := FetchData(rows,data);err==nil{
			result = (data.Data).(*model.User)
			return EndOfFetch
		}else{
			return err
		}
	},sql,uid)
	return
}

func (d *HomeDAO)GetForumList(page,pageSize int) (result []*model.Forum) {
	page = Max(1, page)
	offset := (page - 1) * pageSize
	sql := fmt.Sprintf("select * from %sforum order by `rank` desc LIMIT %d,%d",d.Prefix,page,offset)
	Fetch(func(rows *sqlx.Rows) error {
		data := MkFetchResultSt(&model.Forum{})
		if err :=FetchData(rows,data);err==nil{
			result = append(result,(data.Data).(*model.Forum))
			return nil
		}else{
			return err
		}
	},sql,page,offset)
	return
}
func (d *HomeDAO)Forum_access_find_by_fid(fid uint) (result []*model.Forum_access) {
	sql := fmt.Sprintf("select * from %sforum_access where fid = ? order by gid asc LIMIT 1,100",d.Prefix)
	Fetch(func(rows *sqlx.Rows) error {
		data := MkFetchResultSt(&model.Forum_access{})
		if err :=FetchData(rows,data);err==nil{
			result = append(result,(data.Data).(*model.Forum_access))
			return nil
		}else{
			return err
		}
	},sql,fid)
	return
}

func (d *HomeDAO)User_find_by_uids(uid ...uint) (result []*model.Forum_access) {
	//sql := fmt.Sprintf("select * from %sforum_access where fid = ? order by gid asc LIMIT 1,100",d.Prefix)
	//sql := `# SELECT * FROM users WHERE id IN (?)`
	//sql, params, err := sqlx.In(sql, []int{1,2,3,4,5})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var users []User
	//if err := sqlx.Select(&users, sql, params...); err != nil {
	//	log.Fatal(err)
	//}
	//Fetch(func(rows *sqlx.Rows) error {
	//	data := MkFetchResultSt(&model.Forum_access{})
	//	if err :=FetchData(rows,data);err==nil{
	//		result = append(result,(data.Data).(*model.Forum_access))
	//		return nil
	//	}else{
	//		return err
	//	}
	//},sql,fid)
	return
}

func (d *HomeDAO)User_read(uid uint) (result map[string]interface{}) {
	sql := fmt.Sprintf("select * from %suser where uid = ?",d.Prefix)
	Fetch(func(rows *sqlx.Rows) error {
		data := MkFetchResultMp()
		if err := FetchData(rows,data);err==nil{
			result = (data.Data).(map[string]interface{})
			return EndOfFetch
		}else{
			return err
		}
	},sql,uid)
	return
}


/* for test
func (d *HomeDAO)GetGroupList() (result map[uint]*model.Group) {
	sql := fmt.Sprintf("select * from %sgroup order by gid",d.Prefix)
	var data *model.Group
	_ = reflect.TypeOf(data)
	Fetch(func(rows *sqlx.Rows) error {
		data := MkFetchResultSt(&model.Group{})
		if err := FetchData(rows,data);err==nil{
			//log.Println(newData)
			//data = reflect.ValueOf(newData).Interface().(*model.Group)
			//data = unsafe.Pointer(&newData).(*model.Group)//newData.(*model.Group)
			log.Println(data.Data)
			//result[data.Gid] = data
			return nil
		}else{
			return err
		}
	},sql)

	log.Println("slice")
	Fetch(func(rows *sqlx.Rows) error {
		//var data []interface{}
		data := MkFetchResultSl(nil)
		if err := FetchData(rows,data);err==nil{
			//log.Println(newData)
			//data = reflect.ValueOf(newData).Interface().(*model.Group)
			//data = unsafe.Pointer(&newData).(*model.Group)//newData.(*model.Group)
			log.Println(data.Data)
			//result[data.Gid] = data
			return nil
		}else{
			return err
		}
	},sql)

	log.Println("map")
	Fetch(func(rows *sqlx.Rows) error {
		data := MkFetchResultMp(nil)
		if err := FetchData(rows,data);err==nil{
			//log.Println(newData)
			//data = reflect.ValueOf(newData).Interface().(*model.Group)
			//data = unsafe.Pointer(&newData).(*model.Group)//newData.(*model.Group)
			log.Println(data.Data)
			//result[data.Gid] = data
			return nil
		}else{
			return err
		}
	},sql)
	return
}
 */