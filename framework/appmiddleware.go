package framework

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
)

func (ap *AppBase)DefaultSession() gin.HandlerFunc {
	// セッションの設定
	store := cookie.NewStore([]byte("secret123456"))
	return sessions.Sessions("ReChannel", store)
}

// site entry middleware
func (ap *AppBase)DefaultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//log.Println("DefaultMiddleware is detected.")

		session := sessions.Default(c)
		uid := session.Get("UserId")
		log.Println("DefaultMiddleware:uid:",uid)
		c.Next()
		//c.AbortWithStatus(489)
		//log.Println("after DefaultMiddleware logic.")
	}
}