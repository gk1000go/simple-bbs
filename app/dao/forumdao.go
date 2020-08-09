package dao

import (
	"awesomeProject1/app-web/framework"
	"fmt"
)

type ForumDao struct {
	framework.BaseDAO
}

func NewForumDao() *ForumDao {
	d := &ForumDao{}
	d.Init()
	return d
}

func GetForumList(page,pageSize int) (result framework.ResultMaps) {
	sql := framework.GetSql("GetForumList", func(prefix string) string {
		return fmt.Sprintf("select * from %sforum order by `rank` desc LIMIT ?,?", prefix)
	})
	a,b := makePageOffset(page,pageSize)
	result = FetchResultDataMultiMap(sql,"fid",a,b)
	return
}

func GetForumAccessList(fid uint,page,pageSize int) (result framework.ResultMaps) {
	sql := framework.GetSql("GetForumAccessList", func(prefix string) string {
		return fmt.Sprintf("select * from %sforum_access where fid=? order by gid LIMIT ?,?", prefix)
	})
	a,b := makePageOffset(page,pageSize)
	result = FetchResultDataMultiMap(sql,"gid",fid,a,b)
	return
}
