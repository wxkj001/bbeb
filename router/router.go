package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wxkj001/bbeb/controller/Home"
	"io"
	"html/template"
	"github.com/labstack/echo-contrib/session"
	"github.com/gorilla/sessions"
	"github.com/wxkj001/bbeb/controller/Admin"
)
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
//路由列表
func List() {
	//F:=core.FrameWork{}
	//c:=F.Config()
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Static("/public", "public")
	//t := &Template{
	//	templates: template.Must(template.ParseGlob("./view/*.html")),
	//}
	//e.Renderer=t
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{"Content-Type"},
	}))
	ApiR:=e.Group("/Api")
	{
		AdminRS:=ApiR.Group("/Admin")
		{
			AdminC:=Admin.AdminApiConstroller{}
			AdminRS.POST("/login",AdminC.Login)
		}
	}
	HomeR:=e.Group("/Home")
	HomeR.Use(HomeHeader)

	{
		HomeC:=Home.PublicController{}
		HomeR.GET("/public/env",HomeC.Index)
		HomeR.POST("/public/setmysql",HomeC.SetForm)
		HomeR.POST("/public/setadmin",HomeC.SetAdmin)
	}

	e.Logger.Fatal(e.Start(":3001"))
}
func HomeHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		//c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
		//c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Content-Type")
		//c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, "POST, GET, OPTIONS, PUT, DELETE")
		return next(c)
	}
}
