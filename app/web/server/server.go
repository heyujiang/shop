package server

import (
	"fmt"
	"github.com/heyujiang/shop/app/web/server/routes"
)

func StartWebServer(port int) {
	g := routes.InitRoutes()
	if err := g.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
