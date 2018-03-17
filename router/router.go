package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wxkj001/bbeb/controller/Home"
	"io"
	"html/template"
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
	e.Static("/public", "public")
	//t := &Template{
	//	templates: template.Must(template.ParseGlob("./view/*.html")),
	//}
	//e.Renderer=t

	HomeR:=e.Group("/Home")
	HomeR.Use(HomeHeader)
	{
		HomrC:=Home.PublicController{}
		HomeR.GET("/public/env",HomrC.Index)
	}

	e.Logger.Fatal(e.Start(":3001"))
}
func HomeHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}
