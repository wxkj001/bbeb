package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

//路由列表
func List() {
	Router.Use(Middleware)
	{
		Router.GET("/", func(g *gin.Context) {
			log.Println("Index")
		})
	}
}

//路由中间件
func Middleware(g *gin.Context) {

}
