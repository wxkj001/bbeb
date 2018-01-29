package main

import (
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-xorm/xorm"
	"github.com/wxkj001/bbeb/router"
)

func main() {
	router.Router = gin.Default()
	// g.Delims("{{", "}}")
	// g.Static("/Public/", "./Public")
	// g.LoadHTMLGlob("view/**/*")
	// g.Use(Middleware)
	// {

	// }
	router.List()
	router.Router.Run(":3000")
}
