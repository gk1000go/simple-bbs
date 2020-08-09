package svcfunc

import (
	"awesomeProject1/app-web/app/dao"
	"awesomeProject1/app-web/app/services"
	"awesomeProject1/app-web/framework"
	"fmt"
)



func forum_list_cache() framework.ResultMaps {
	//$forumlist = cache_get('forumlist');
	list := dao.GetForumList(1,1000)
	//cache_set('forumlist', $forumlist, 60);
	return list
}

func forum_format(forum framework.BaseDataMap) {
	if forum == nil {return}

	forum["create_date_fmt"] = services.UIntToTime(forum.UInt("create_date"))
	site := framework.GetApp().Cfg.BaseSite
	if forum.UInt("icon") > 0{
		forum["icon_url"] = fmt.Sprintf("%sforum/%d.png",site.Upload_url,forum.UInt("fid"))
	}else{
		forum["icon_url"] = site.View_url + "img/forum.png"
	}
	forum["accesslist"] = nil
	if forum.Int("accesson") > 0{
		forum["accesslist"] = forum_access_find_by_fid(forum.UInt("fid"))
	}
	forum["modlist"] = nil
	if len(forum.Str("moduids"))>0 {
		modlist := user_find_by_uids(forum.Str("moduids"))
		for _,v := range modlist{
			user_safe_info(v)
		}
	}
}
func forum_access_find_by_fid(fid uint) framework.ResultMaps {
/*	$cond = array('fid'=>$fid);
	$orderby = array('gid'=>1);
	$accesslist = db_find('forum_access', $cond, $orderby, 1, 100, 'gid');
 */
	return dao.GetForumAccessList(fid,1,100)
}
// 对 $forumlist 权限过滤，查看权限没有，则隐藏 -- allow:allowread
func forum_list_access_filter(forumList map[uint]framework.BaseDataMap,gid uint,allow string) {
	if forumList == nil || gid==1 {return }

	gList := Group_list_cache()
	group := gList[gid]


	for key, forum := range forumList {
		chk := forum.UInt("accesson") == 0 && group.Int(allow) == 0
		accesslist := forum["accesslist"].(map[uint]interface{})
		chk = chk || forum.UInt("accesson")>0 && (accesslist[gid].(framework.BaseDataMap)).Int(allow) == 0
		if chk {
			delete(forumList,key)
		}else{
			forum["accesslist"] = nil
		}

	}
}

/*
// 第二种
func BenchmarkReuse(t *testing.B) {
    t.ResetTimer()

    origin := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    for i := 0; i < t.N; i++ {
        target := origin[:0]
        for _, item := range origin {
            if item != 6 {
                target = append(target, item)
            }
        }
    }
}
 */