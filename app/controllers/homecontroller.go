package controllers

import (
	"awesomeProject1/app-web/app/model"
	"awesomeProject1/app-web/app/services"
	"awesomeProject1/app-web/conf"
	"awesomeProject1/app-web/framework"
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"html/template"
	"strings"
)

type HomeController struct {
	framework.BaseController
	page int
	order string
	pageSize int
	active string
	uid uint
	gid uint
	mGroup *model.Group
	service *services.HomeService
}

// 初始化函数
func (this *HomeController)InitEntry() {
	this.RenderType = framework.RenderTypeJson
	this.page = 1
	this.order = conf.DefaultConfig.BaseSite.Order_default
	this.pageSize = conf.DefaultConfig.BaseSite.Pgesize
	this.active = "default"
	//if token := this.Sessions.Get("UserId");token != nil { //string:bbs_token223456:bbs_token
	//	log.Println("token type:",reflect.TypeOf(token))
	//	log.Println("token type:",reflect.ValueOf(token))
	//	//this.uid = token.(uint)
	//}

	this.service = services.NewHomeService()
}

func (this *HomeController) DoRequest() {
	c := this.Ctx

	// 语言 / Language
	// lang = include _include(APP_PATH."lang/$conf[lang]/bbs.php");

	// 用户组 / Group
	groupList := this.service.GetGroupList()

	// 支持 Token 接口（token 与 session 双重登陆机制，方便 REST 接口设计，也方便 $_SESSION 使用）
	// Support Token interface (token and session dual match, to facilitate the design of the REST interface, but also to facilitate the use of $_SESSION)
	uid := this.GetSessionUserId()
	if uid == 0{
		uid = GetUserIdFromToken(c)
	}
	if uid != 0{
		this.SetSessionUserId(uid)
	}
	user1 := this.service.GetUser(uid)

	//gid := uint(0)
	//group1 := groupList[0]
	//if user1 != nil{
	//	gid = user1.Gid
	//	if v,ok := groupList[gid];ok{
	//		group1 = v
	//	}
	//}

	// 版块 / Forum
	fid := uint(0)
	forumlist := this.service.GetForumList()
	//forumList := this.service.GetForumList()
	//for k,v := range forumList{
	//
	//}


	log.Println(groupList)
	log.Println(user1)

	//this.gid = group.Gid

	//$forumlist = forum_find();
					//forum__find($cond, $orderby:array('rank'=>-1)[desc], $page=1, $pagesize=1000);

	//if($forumlist) foreach ($forumlist as &$forum) forum_format($forum);


	//$group = isset($grouplist[$gid]) ? $grouplist[$gid] : $grouplist[0];

	//svr := services.NewHomeService()
	// 从默认的地方读取主题列表
	// forumlist_show 二维数组
	//fids := svr.Get_fids()
	//threads := svr.Get_threads()
	//pagination := svr.Get_pagination()
	//threadlist := svr.Get_threadlist()


	action := ""
	if(len(c.Params) > 0) {
		action = c.Param("action")
		if action != "" {
			log.Println("action:",action," called.")
		}

		_ = strings.Split(action,"/") // "" -> len:1

	}
	this.Sessions.Set("bbs_token","bbs_token223456")
	this.Sessions.Set("UserId",1012345)
	_=this.Sessions.Save()
	log.Println(c.Params) // index/a/b/c/d => /a/b/c/d
	log.Println(c.Query("a")) // ?a=b
	//c.JSON(http.StatusOK,"GetAction_OK")
	this.JsonData = groupList//framework.SiteData{"result": "HomeController GetAction_OK"}

	fmt.Println("")

	//tmplData["uid"] = ""
	//tmplData["forumlist_show"] = forumlist
	//tmplData["user"] = map[string]string{
	//	"avatar_url":"avatar_url111",
	//	"username":"username12",
	//}
	//tmplData["gid"] = 1
	//tmplData["fid"] = 12
	//tmplData["active"] = "active" //'default' ? 'active'
	//tmplData["siteurl"] = siteurl
	//tmplData["sitelang"] = lang.Lang
	//tmplData["pagination"] = "pagination..." //分页
	//tmplData["runtime"] = siteruntime
	//
	//c.HTML(http.StatusOK, "index.htm", tmplData)

	//c.QueryMap()
}
func (this *HomeController) EndOfRequest() {
	log.Println("EndOfRequest executed.")
	this.mGroup = nil
}
//func (this *HomeController) BeforeRequest() {
//	//log.Info("BeforeRequest")
//}
//func (this *HomeController) AfterRequest() {
//	//log.Info("AfterRequest")
//	//client := &http.Client{};
//	//req,err := http.NewRequest("","http://localhost:8080/user/1/2/3?a=bcdef",nil)
//	//if(err!=nil){log.Info(err)}
//	//response, err := client.Do(req)
//	//defer response.Body.Close()
//	//log.Info(ioutil.ReadAll(response.Body))
//	////https://github.com/kirinlabs/HttpRequest
//}
//func (this *HomeController) Render(){
//	this.BaseController.Render()
//}

var forumlist = []map[string]string{
	{
		"fid":"idd1",
		"name":"name1...",
		"forum_fid":"aaurl1",
	},
	{
		"fid":"idd2",
		"name":"name2...",
		"forum_fid":"aaurl2",
	},
}
var siteurl = map[string]string{
	"route":"localhost:8080",
	"thread_create_fid":"localhost:8080/thread_create_fid_url",
	"user_login":"localhost:8080/user-login",
	"my":"mymy",
	"user_logout":"user-logout",
}
var siteruntime = map[string]interface{}{
	"threads":"threads...",
	"posts":"posts...",
	"users":"users...",
	"onlines":123,
}

func createRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/layout.tmpl", "templates/index.tmpl")
	return r
}

func home2_test(c *gin.Context) {
	log.Println("HomeController home2 is executed! ")
	t, err := template.ParseFiles("templates/first.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	t.Execute(c.Writer, struct {

	}{})
}

//func (this *HomeController)InitMiddleware()  {
//	this.BaseController.InitMiddleware()
//	this.Group.Use(this.homeMiddleware())
//}
//func (this *HomeController) homeMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		log.Println("before homeMiddleware logic.")
//		c.Next()
//		log.Println("after homeMiddleware logic.")
//	}
//}