package controllers

import (
	"awesomeProject1/app-web/framework"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserController struct {
	framework.BaseController
	page int
	order string
	pagesize int
	active string
}

func (this *UserController)InitEntry() {
	this.RenderType = framework.RenderTypeJson
}

func (this *UserController) DoRequest() {
	// data := base.GetStaticData()
	c := this.Ctx
	log.Println("UserController GetAction is executed! ")

	//action := c.Param("action")
	log.Println(c.Params)
	log.Println(c.Query("a"))
	//c.JSON(http.StatusOK,"GetAction_OK")
	this.JsonData = framework.SiteData{"result": "UserController : GetAction_OK"}
}

//var forumlist = []map[string]string{
//	{
//		"fid":"idd1",
//		"name":"name1...",
//		"forum_fid":"aaurl1",
//	},
//	{
//		"fid":"idd2",
//		"name":"name2...",
//		"forum_fid":"aaurl2",
//	},
//}
//var siteurl = map[string]string{
//	"route":"localhost:8080",
//	"thread_create_fid":"localhost:8080/thread_create_fid_url",
//	"user_login":"localhost:8080/user-login",
//	"my":"mymy",
//	"user_logout":"user-logout",
//}
//var siteruntime = map[string]interface{}{
//	"threads":"threads...",
//	"posts":"posts...",
//	"users":"users...",
//	"onlines":123,
//}

func (this *UserController)home2() gin.HandlerFunc {
	return func	(c * gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"title": "hello my home2!",
		})
	}
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