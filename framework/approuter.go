package framework

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IAppRouter interface {
	InitRouter()
}

type AppGroupRouter struct {
	Name       string
	RootRouter string
	SubRouter  string
	HtmlTmpl   string
	ICObject   IController
	GinGroup   *gin.RouterGroup
}

func (ap *AppBase)InitRouter() {
	ap.GinEngine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"result":"hello world ..."})
	})
}