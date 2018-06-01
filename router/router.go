package router

import (
	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	casbinmw "github.com/labstack/echo-contrib/casbin"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/wxkj001/bbeb/controller/Admin"
	"github.com/wxkj001/bbeb/controller/Home"
)

var enforcer *casbin.Enforcer
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
	//e.Pre(middleware.AddTrailingSlash())
	//t := &Template{
	//	templates: template.Must(template.ParseGlob("./view/*.html")),
	//}
	//e.Renderer=t
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{"Content-Type"},
	}))
	a := xormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/blog")
	enforcer=casbin.NewEnforcer("data/casbin_auth_model.conf", a)
	e.Use(casbinmw.MiddlewareWithConfig(casbinmw.Config{Skipper:middleware.Skipper(APIAuth),Enforcer: enforcer}))
	//e.Use(casbinmw.Middleware(enforcer))
	APIR := e.Group("/api")
	APIR.Use(APIHeader)
	{
		AdminRS := APIR.Group("/admin")
		{
			AdminC := admin.APIConstroller{}
			//登录
			AdminRS.POST("/login", AdminC.Login)
			//文章列表
			AdminRS.GET("/article/list", AdminC.ArticleList)
			//文章详情
			AdminRS.GET("/article/:id", AdminC.ArticleInfo)
			//添加文章
			AdminRS.POST("/article/", AdminC.ArticleADD)
			//更新文章
			AdminRS.PUT("/article/:id", AdminC.ArticleEDIT)
			//删除文章
			AdminRS.DELETE("/article/:id", AdminC.ArticleDEL)
		}
		HomeRS := APIR.Group("/home")
		{
			homeC := home.APIConstroller{}
			HomeRS.POST("/article/list", homeC.ArticleList)
		}
	}
	HomeR := e.Group("/Home")
	HomeR.Use(HomeHeader)
	{
		HomeC := home.PublicController{}
		HomeR.GET("/public/env", HomeC.Index)
		HomeR.POST("/public/setmysql", HomeC.SetForm)
		HomeR.POST("/public/setadmin", HomeC.SetAdmin)
	}

	e.Logger.Fatal(e.Start(":3001"))
}

func HomeHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Blog/1.0")
		return next(c)
	}
}
func APIAuth(c echo.Context) bool {
	user :="alice"
	method := c.Request().Method
	path := c.Request().URL.Path
	b:=enforcer.Enforce(user, path, method)
	return b
}
func APIHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Blog/1.0")
		return next(c)
	}
}
