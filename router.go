package main

import (
	ctrl "awesomeProject1/app-web/app/controllers"
	fr "awesomeProject1/app-web/framework"
)

type MyRouter struct {
	Name string
}

func makeRouter(name string,root string,sub string,tmpl string,new bool)  {
	if c := ctrl.Ctrl(name,new);c != nil {
		ap := fr.GetApp()
		router := &fr.AppGroupRouter{
			name,
			root,
			sub,
			tmpl,
			c,
			ap.GinEngine.Group(root),
		}
		ap.SetRouter(router)
	}
}
func (r *MyRouter)InitRouter() {
	// localhost:8080/?abc=b
	makeRouter("HomeController","/","/index/*action","index.htm",false)
	//makeRouter("HomeController","/test","/*action1","index.htm",true)
	makeRouter("UserController","/user","/*action","",false)
}
/*
if(DEBUG > 1) {
	include XIUNOPHP_PATH.'xiunophp.php';
} else {
	include XIUNOPHP_PATH.'xiunophp.min.php';
}

// 测试数据库连接 / try to connect database
//db_connect() OR exit($errstr);

include APP_PATH.'model/plugin.func.php';
include _include(APP_PATH.'model.inc.php');
include _include(APP_PATH.'index.inc.php');

*/


