package dao

import (
	"awesomeProject1/app-web/framework"
	"fmt"
)

type GroupDao struct {
	framework.BaseDAO
}

func NewGroupDAO() *GroupDao {
	d := &GroupDao{}
	d.Init()
	return d
}

//func (d *GroupDao)GetGroupList() (result map[uint]*model.Group) {
//	sql := fmt.Sprintf("select * from %sgroup order by gid",d.Prefix)
//	result = make(map[uint]*model.Group)
//	Fetch(func(rows *sqlx.Rows) error {
//		data := MkFetchResultSt(&model.Group{})
//		if err := FetchData(rows,data);err==nil{
//			gdata := (data.Data).(*model.Group)
//			result[gdata.Gid] = gdata
//			//log.Println(data.Data)
//			return nil
//		}else{
//			return err
//		}
//	},sql)
//
//	return
//}

func GetGroupList() (result framework.ResultMaps) {
	sql := framework.GetSql("GetGroupList", func(prefix string) string {
		return fmt.Sprintf("select * from %sgroup order by gid",prefix)
	})
	result = FetchResultDataMultiMap(sql,"gid")
	return
}