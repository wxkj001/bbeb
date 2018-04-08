package core

import (
	"log"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)
var RegStruct map[string]interface{}
var ORM *xorm.Engine
type H map[string]interface{}
//框架
type FrameWork struct {
}

func (f *FrameWork)ORM() *xorm.Engine {
	return ORM
}
func ORMLoad() (*xorm.Engine,error) {
	//F:=FrameWork{}
	//c:=F.Config()
	//host:=c.Section("database").Key("host").String()
	//user:=c.Section("database").Key("username").String()
	//password:=c.Section("database").Key("password").String()
	//database:=c.Section("database").Key("database").String()
	//port:=c.Section("database").Key("port").String()
	//if password==""&&database==""{
	//	return ORM,errors.New("配置文件错误")
	//}
	//mapping := func(s string) string {
	//	m := map[string]string{
	//		"host": host,
	//		"user": user,
	//		"password":password,
	//		"dbname":database,
	//		"port":port,
	//	}
	//	return m[s]
	//}
	//ormStr:=os.Expand("$user:$password@tcp($host:$port)/$dbname?charset=utf8&parseTime=True&loc=Local",mapping)
	//orm, err := xorm.NewEngine("mysql", ormStr)
	//if err!=nil{
	//	log.Println("数据库连接失败")
	//	return ORM,err
	//}
	//orm.Sync2(new(model.BlogUsers),new(model.BlogTags),new(model.BlogArticle))
	//ORM=orm
	return ORM,nil
}
func (f *FrameWork)Config() (file *ini.File) {
	cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, "data/app.ini")
	if err!=nil{
		log.Fatalln("配置文件读取失败")
	}
	return cfg
}

type Model struct {

} 