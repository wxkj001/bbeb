package main

import (
	"flag"
	"github.com/wxkj001/bbeb/core"
	"github.com/wxkj001/bbeb/router"
	"os"
	"github.com/go-xorm/xorm"
	"log"
	"github.com/wxkj001/bbeb/model"
)

func init() {
	F := core.FrameWork{}
	c := F.Config()
	host := c.Section("database").Key("host").String()
	user := c.Section("database").Key("username").String()
	password := c.Section("database").Key("password").String()
	database := c.Section("database").Key("database").String()
	port := c.Section("database").Key("port").String()
	if password == "" && database == "" {

	}
	mapping := func(s string) string {
		m := map[string]string{
			"host":     host,
			"user":     user,
			"password": password,
			"dbname":   database,
			"port":     port,
		}
		return m[s]
	}
	ormStr := os.Expand("$user:$password@tcp($host:$port)/$dbname?charset=utf8&parseTime=True&loc=Local", mapping)
	orm, err := xorm.NewEngine("mysql", ormStr)
	if err != nil {
		log.Println("数据库连接失败")
	}
	orm.Sync2(new(model.BlogUsers), new(model.BlogTags), new(model.BlogArticle))
	core.ORM = orm
}

var s = flag.String("s", "start", "stop|start|reload")
var b = flag.String("b", "build", "build")

func main() {
	router.List()
	//flag.Parse()
	//out := make(chan int)
	//if len(os.Args) <= 1 {
	//	flag.PrintDefaults()
	//	log.Fatalln("ERROR")
	//}
	//input := os.Args[1]
	//if input == "-s" {
	//	switch *s {
	//	case "start":
	//		go runs(out)
	//		<-out
	//		break
	//	case "stop":
	//		break
	//	case "reload":
	//		break
	//	}
	//} else if input == "-b" {
	//	log.Println("-b")
	//} else {
	//	flag.PrintDefaults()
	//}
}
func runs(in chan int) {
	_, err := core.ORMLoad()
	if err != nil {

	}
	router.List()
}
