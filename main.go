package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wxkj001/bbeb/router"
	"github.com/wxkj001/bbeb/core"
	"os"
	"github.com/go-xorm/xorm"
	"log"
	"github.com/wxkj001/bbeb/model"
)
func init()  {
	F:=core.FrameWork{}
	c:=F.Config()
	host:=c.Section("database").Key("host").String()
	user:=c.Section("database").Key("username").String()
	password:=c.Section("database").Key("password").String()
	database:=c.Section("database").Key("database").String()
	port:=c.Section("database").Key("port").String()
	mapping := func(s string) string {
		m := map[string]string{
			"host": host,
			"user": user,
			"password":password,
			"dbname":database,
			"port":port,
		}
		return m[s]
	}
	ormStr:=os.Expand("$user:$password@tcp($host:$port)/$dbname?charset=utf8&parseTime=True&loc=Local",mapping)
	orm, err := xorm.NewEngine("mysql", ormStr)
	if err!=nil{
		log.Println("数据库连接失败")
	}
	orm.Sync2(new(model.BlogUsers))
	core.ORM=orm
}
func main() {
	router.Router = gin.Default()
	 //g.Delims("{{", "}}")
	// g.Static("/Public/", "./Public")
	// g.LoadHTMLGlob("view/**/*")
	// g.Use(Middleware)
	// {

	// }
	router.List()
	router.Router.Run(":3000")
}
