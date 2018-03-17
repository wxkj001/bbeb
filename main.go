package main

import (
	"github.com/wxkj001/bbeb/core"
	"github.com/labstack/echo"
	"os"
	"github.com/go-xorm/xorm"
	"log"
	"github.com/wxkj001/bbeb/model"
	"net/http"
	"github.com/wxkj001/bbeb/router"
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
	orm.Sync2(new(model.BlogUsers),new(model.BlogTags),new(model.BlogArticle))
	core.ORM=orm
}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func main() {
	router.List()
}
