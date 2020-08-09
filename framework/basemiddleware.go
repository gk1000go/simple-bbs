package framework

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type BaseMiddleware struct {

}



// group middleware
func (this *BaseMiddleware) baseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("baseMiddleware is detected.")
		//log.Println("before baseMiddleware logic.")
		c.Next()
		//log.Println("after baseMiddleware logic.")
	}
}