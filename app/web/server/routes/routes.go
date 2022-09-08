package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/heyujiang/shop/app/web/server/routes/sys"
)

func InitRoutes() *gin.Engine {
	g := gin.Default()
	_ = g.SetTrustedProxies(nil)

	g.GET("/login", sys.Login)

	g.GET("/users/:id", sys.GetUserInfo)

	return g
}
