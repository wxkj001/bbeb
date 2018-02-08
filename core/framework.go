package core

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"io/ioutil"
	"regexp"
	"strings"
	"go/token"
	"go/parser"
	"reflect"
	"github.com/go-ini/ini"
	"github.com/gin-contrib/sessions"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)
var RegStruct map[string]interface{}
var SESSION sessions.Session
var ORM *xorm.Engine
//框架
type FrameWork struct {
}

func (f *FrameWork)ORM() *xorm.Engine {
	return ORM
}
func (f *FrameWork)Session() sessions.Session {
	return SESSION
}
func (f *FrameWork)Config() (file *ini.File) {
	cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, "data/app.ini")
	if err!=nil{
		log.Fatalln("配置文件读取失败")
	}
	return cfg
}
func Annotation(Router *gin.Engine)  {
	t:=token.NewFileSet()
	f, err :=os.OpenFile("./controller/Admin/Index.go", os.O_RDWR|os.O_CREATE, 0755)
	if err!=nil{
		log.Fatal("File error")
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	p,e:=parser.ParseFile(t,"Index.go",data,parser.ParseComments)
	if e!=nil{
		log.Fatal("error")
		log.Fatal(e.Error())
	}
	for _,z:=range p.Comments {
		RouterReg := regexp.MustCompile(`@Router`)
		RR:=RouterReg.FindString(z.Text())
		if len(RR)>0{
			reg := regexp.MustCompile(`(/\w+)`)
			rz:=reg.FindString(z.Text())
			reg2 := regexp.MustCompile(`(\w+)(,)(\w+)`)
			rz2:=reg2.FindString(z.Text())
			data:=strings.Split(rz2,",")
			//log.Println(data[:1][0])
			//log.Println(data[1:2][0])
			switch data[:1][0] {
			case "POST":
				log.Println("POST")
				for _, s := range p.Scope.Objects {
					if s.Kind.String()=="type"{
						log.Println(s.Name)
						if RegStruct[s.Name] != nil {
							t := reflect.ValueOf(RegStruct[s.Name]).Type()
							v := reflect.New(t).Elem()
							//log.Println(t.)
							log.Println(v.NumMethod())
						}
					}
				}
				//str := PackName
				Router.GET(rz, func(context *gin.Context) {

				})
				break
			case "get":

				break
			}

		}
	}
}
//模型
type Model struct {

}

//func (m *Model) ORM() (*xorm.Engine) {
//	F:=FrameWork{}
//	c:=F.Config()
//	host:=c.Section("database").Key("host").String()
//	user:=c.Section("database").Key("username").String()
//	password:=c.Section("database").Key("password").String()
//	database:=c.Section("database").Key("database").String()
//	port:=c.Section("database").Key("port").String()
//	mapping := func(s string) string {
//		m := map[string]string{
//			"host": host,
//			"user": user,
//			"password":password,
//			"dbname":database,
//			"port":port,
//		}
//		return m[s]
//	}
//	ormStr:=os.Expand("$user:$password@tcp($host:$port)/$dbname?charset=utf8&parseTime=True&loc=Local",mapping)
//	orm, err := xorm.NewEngine("mysql", ormStr)
//	if err!=nil{
//		log.Println("数据库连接失败")
//	}
//	return orm
//}