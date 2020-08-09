package framework

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"

	//log "github.com/sirupsen/logrus"
)

const (
	RenderTypeHTML = iota
	RenderTypeJson
	RenderTypeOther
)
type BaseController struct {
	RouterInfo *AppGroupRouter
	Middleware *BaseMiddleware
	PageData   interface{}
	JsonData   interface{}
	PagePool   *sync.Pool
	Ctx        *gin.Context
	Sessions   sessions.Session
	RenderType int
}

func (this *BaseController)GetSessionUserId() uint {
	if uid := this.Sessions.Get("UserId");uid != nil {
		return uint(uid.(int))
	}
	return 0
}
func (this *BaseController)SetSessionUserId(uid uint) {
	this.Sessions.Set("UserId",uid)
	// to save?????
}
func (this *BaseController)GetSessionToken() string {
	if token := this.Sessions.Get("bbs_token");token != nil {
		return token.(string)
	}
	return ""
}
/*
pool 专用。没啥用的话，删除了也没关系
 */
//type IBasePoolFunc interface {
//	Reset()
//	Clear()
//}
//type PageDataSt map[string]interface{}
//func (pd PageDataSt)Reset() {} // 出库的时候需要重置的函数
//func (pd PageDataSt)Clear() {} // 入库的时候需要清空的函数
//func (this *BaseController)GetPoolPageData() PageDataSt {
//	data := this.PagePool.Get()
//	if(data != nil) {
//		this.PageData = data.(PageDataSt)
//		this.PageData.Reset()
//		return this.PageData
//	}
//	return nil
//}
//func (this *BaseController)ReleasePoolPageData() {
//	if(this.PageData != nil) {
//		this.PageData.Clear()
//		this.PagePool.Put(this.PageData)
//	}
//}
//func (this *BaseController)GetBasePageData() PageDataSt {
//	return PageDataSt{"siteconf": conf.SiteConf,}
//}

/*
get method
*/
//func (this *BaseController)do_get() gin.HandlerFunc {
//	return this.do_post()
//}
/*
post method
*/
func (this *BaseController)do_post() gin.HandlerFunc {
	return func (c *gin.Context){
		this.Ctx = c
		this.Sessions = sessions.Default(c)
		ctrlObj := this.RouterInfo.ICObject // controller chain???
		ctrlObj.InitEntry() // page  init
		ctrlObj.BeforeRequest()
		ctrlObj.DoRequest() // 各自实现
		ctrlObj.AfterRequest()
		ctrlObj.Render() // use chain????
		ctrlObj.EndOfRequest()
	}
}

/*
frame work
*/
func (this *BaseController)InitBaseController(r *AppGroupRouter) { // 需要的话，重写，先调用base.InitBaseController
	this.RouterInfo = r
	this.Middleware = &BaseMiddleware{} // group middle ware
	//this.PagePool = &sync.Pool{New: func() interface{} { return this.GetBasePageData() },}

	this.RenderType = RenderTypeHTML // default HTML output
	if(r.RootRouter == "/") {
		r.GinGroup.GET("/",this.do_post())
		r.GinGroup.POST("/",this.do_post())
	}
	r.GinGroup.GET(r.SubRouter,this.do_post()) // 8080/user or 8080/user/a/b/c => action=/ or action=/a/b/c
	r.GinGroup.POST(r.SubRouter,this.do_post())
}
func (this *BaseController)InitMiddleware() { // 需要的话，重写，先调用base.InitMiddleware
	this.RouterInfo.GinGroup.Use(this.Middleware.baseMiddleware())
}
func (this *BaseController) InitEntry()     {} // 需要的话，重写，覆盖掉（用于初始化）
func (this *BaseController) BeforeRequest() {} // 需要的话，重写，覆盖掉
func (this *BaseController) DoRequest()     {} // 重写，覆盖掉 Get & Post
func (this *BaseController) AfterRequest()  {} // 需要的话，重写，覆盖掉
func (this *BaseController) Render() { // 需要的话，重写，覆盖掉
	switch(this.RenderType){
	case RenderTypeHTML:
		this.Ctx.HTML(http.StatusOK,this.RouterInfo.HtmlTmpl,this.PageData)
	case RenderTypeJson:
		this.Ctx.JSON(http.StatusOK,this.JsonData)
	case RenderTypeOther:
	default:
	}
}
func (this *BaseController) EndOfRequest() {
}


