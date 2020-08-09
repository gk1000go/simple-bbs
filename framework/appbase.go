package framework

import (
	"awesomeProject1/app-web/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//type IAppBase interface {
//	InitRouter()
//	Run()
//}

type AppBase struct {
	Server    *http.Server
	GinEngine *gin.Engine
	Routers   *AppGroupRouter
	Cfg       *conf.AppConfig
}

type SiteData map[string]interface{}

var siteApp *AppBase
func NewApp(c *conf.AppConfig) *AppBase {
	cfg := c
	if (cfg == nil) {
		cfg = &conf.DefaultConfig
	}

	connStr := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?parseTime=True","root","localhost","3406","next9")
	connectToDB(DBTypeMysql,connStr)

	siteApp = &AppBase{
		//Routers:make(map[string][]*AppGroupRouter),
		Cfg: cfg,
	}
	siteApp.initEngine()

	return siteApp
}
func GetApp() *AppBase {
	if(siteApp == nil) {
		panic("No AppBase.")
	}
	return siteApp
}

func (ap *AppBase)initEngine() {
	ap.GinEngine = gin.Default()
	if (ap.GinEngine == nil) {
		panic("No gin.Engine.")
	}

	// ap.GinEngine.Static("/assets", "./static")
	// LoadHTMLFiles or LoadHTMLGlob is only one can be work
	// -- ap.GinEngine.LoadHTMLGlob("templates/**/*")
	// instead to use github.com/gin-contrib/multitemplate
	ap.GinEngine.Static(ap.Cfg.BaseSite.View_url, siteApp.Cfg.BaseSite.View_path) //"/unmoved","static/"
	ap.GinEngine.LoadHTMLGlob(ap.Cfg.BaseSite.Tmpl_path)                          //"view/htm/*.tmpl"

	ap.GinEngine.Use(ap.DefaultSession(),ap.DefaultMiddleware())
}

func (app *AppBase) AddRouter(name string,root string,sub string,tmpl string,ic IController)  {
	router := &AppGroupRouter{
		name,
		root,
		sub,
		tmpl,
		ic,
		app.GinEngine.Group(root),
	}
	app.SetRouter(router)
}
func (app *AppBase) SetRouter(router *AppGroupRouter)  {
	if router == nil { return }
	//app.Routers[router.Name] = router
	ic := router.ICObject.(IController)

	ic.InitBaseController(router)
	ic.InitMiddleware()
}

type IController interface {
	InitBaseController(r *AppGroupRouter)
	InitMiddleware()
	InitEntry()
	BeforeRequest()
	DoRequest()
	AfterRequest()
	Render()
	EndOfRequest()
}

//func (this *AppGroupRouter)DispatchController(){
//	this.GinGroup = this.App.GinEngine.Group(this.RootRouter)
//
//	ic:=this.ICObject
//
//	ic.InitBaseController()
//	ic.InitMiddleware()
//	ic.InitEntry()
//
//	this.App.controllers[this.Name] = this
//}
