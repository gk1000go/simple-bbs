package svcfunc

import (
	"awesomeProject1/app-web/app/dao"
	"awesomeProject1/app-web/framework"
)

var GroupList framework.ResultMaps

func Group_list_cache() framework.ResultMaps {
	// cache_get
	if GroupList ==nil {
		GroupList = dao.GetGroupList()
	}
	// cache_set
	return GroupList
}
func get_group_name(gid uint) string {
	g := Group_list_cache()
	if v,ok := g[gid];ok{
		return v.Str("name")
	}
	return ""
}