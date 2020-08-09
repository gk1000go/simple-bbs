package svcfunc

import (
	"awesomeProject1/app-web/app/dao"
	"awesomeProject1/app-web/app/services"
	"awesomeProject1/app-web/framework"
	"fmt"
	"strings"
)

func user_token_get() uint {
	return 0
}

func User_read(uid uint) framework.BaseDataMap {
	user := dao.GetUserSingle(uid)
	user_format(user)
	//$g_static_users[$uid] = $user;
	return user
}

func user_format(u framework.BaseDataMap)  {
	if u == nil {return}

	u["create_ip_fmt"] = services.Long2ip(uint32(u.UInt("create_ip")))
	u["create_date_fmt"] = services.UIntToTime(u.UInt("create_date"))
	u["login_ip_fmt"] = services.Long2ip(uint32(u.UInt("login_ip")))
	u["login_date_fmt"] = services.UIntToTime(u.UInt("login_date"))

	u["groupname"] = get_group_name(u.UInt("gid"))

	dir := services.Substr(fmt.Sprintf("%09d",u.UInt("uid")),0,3) //$dir = substr(sprintf("%09d", $user['uid']), 0, 3);

	cfg := framework.GetApp().Cfg
	u["avatar_url"] = "view/img/avatar.png"
	u["avatar_path"] = ""
	if u.UInt("avatar") > 0{
		u["avatar_url"] = fmt.Sprintf("%savatar/%s/%d.png?%d",cfg.BaseSite.Upload_url,dir,u.UInt("uid"),u.UInt("avatar"))
		u["avatar_path"]  = fmt.Sprintf("%savatar/%s/%d.png?%d",cfg.BaseSite.Upload_path,dir,u.UInt("uid"),u.UInt("avatar"))
	}
	u["online_status"] = 1
}

func user_find_by_uids(uids string) framework.ResultMaps {
	if len(uids) == 0 {return nil}
	result := make(framework.ResultMaps)

	uidList := strings.Split(uids,",")
	for _,v := range uidList{
		user := user_read_cache(services.Str2Uint(v));
		if(user == nil) {
			continue
		}
		result[user.UInt("uid")] = user
	}
	return result
}

func user_read_cache(uid uint) framework.BaseDataMap {
	//if(isset($g_static_users[$uid])) return $g_static_users[$uid];
	if(uid == 0) {return user_guest()}
	//$r = cache_get("user-$uid");
	result := User_read(uid)
	if result == nil {return user_guest()}
	//cache_set("user-$uid", $r);
	return result
}

func user_safe_info(u framework.BaseDataMap) {
	delete(u,"password")
	delete(u,"email")
	delete(u,"salt")
	delete(u,"password_sms")
	delete(u,"idnumber")
	delete(u,"realname")
	delete(u,"qq")
	delete(u,"mobile")
	delete(u,"create_ip")
	delete(u,"create_ip_fmt")
	delete(u,"create_date")
	delete(u,"create_date_fmt")
	delete(u,"login_ip")
	delete(u,"login_date")
	delete(u,"login_ip_fmt")
	delete(u,"login_date_fmt")
	delete(u,"logins")
	delete(u,"password")
	delete(u,"password")
	delete(u,"password")
}

var user_uest_data framework.BaseDataMap
func user_guest() framework.BaseDataMap {
	if user_uest_data != nil {return user_uest_data
	}
	user_uest_data = framework.BaseDataMap{
		"uid" :0,
		"gid" : 0,
		"groupname" : "guest_group",//lang("guest_group"),
		"username" : "guest", //lang("guest"),
		"avatar_url" : "view/img/avatar.png",
		"create_ip_fmt" : "",
		"create_date_fmt" : "",
		"login_date_fmt" : "",
		"email" : "",

		"threads" : 0,
		"posts" : 0,
	}
	return user_uest_data
	//return &model.User{
	//	Uid : 0,
	//	Gid : 0,
	//	Groupname : "guest_group" ,	//lang("guest_group'),
	//	Username : "guest",			//lang('guest'),
	//	Avatar_url : "view/img/avatar.png",
	//	Create_ip_fmt : "",
	//	Create_date_fmt : nil,
	//	Login_date_fmt: nil,
	//	Email : "",
	//
	//	Threads : 0,
	//	Posts : 0,
	//}
}