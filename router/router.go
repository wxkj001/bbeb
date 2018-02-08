package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wxkj001/bbeb/core"
	"github.com/wxkj001/bbeb/controller/Admin"
	"github.com/gin-contrib/sessions"
)

var Router *gin.Engine

//路由列表
func List() {
	F:=core.FrameWork{}
	c:=F.Config()
	Router.Delims(c.Section("TMPL").Key("L_DELIM").String(),c.Section("TMPL").Key("R_DELIM").String())
	Router.Static("/public/", "./public")
	Router.LoadHTMLGlob("view/**/**/*")
	Admins:=Router.Group("Admin")
	store := sessions.NewCookieStore([]byte("session"))
	Admins.Use(sessions.Sessions("admin", store))
	Admins.Use(Middleware)
	{
		Index:=Admin.Index{}

		Admins.GET("/", Index.Index)
		Admins.GET("/login", Index.Login)
		Admins.POST("/login",Index.LoginPost)
		Admins.GET("/IndexInfo",Index.IndexInfo)
		{
			Article:=Admin.Article{}
			Admins.GET("/Article",Article.ArticleGet)
		}
	}
}

//路由中间件
func Middleware(g *gin.Context) {
	 core.SESSION = sessions.Default(g)
	 //core.SESSION.Set("time",time.Now())
	 //core.SESSION.Save()

}
