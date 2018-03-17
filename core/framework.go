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
func (f *FrameWork)Config() (file *ini.File) {
	cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, "data/app.ini")
	if err!=nil{
		log.Fatalln("配置文件读取失败")
	}
	return cfg
}
//模型
type Model struct {

}