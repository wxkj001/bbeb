package model

import (
	"github.com/wxkj001/bbeb/core"
	"log"
	"time"
)

type UserModel struct {
	*core.Model
}
type BlogUsers struct {
	Id int64 `xorm:"int(11) pk autoincr unique index"`
	UserName string `xorm:"varchar(255)"`
	UserEmail string `xorm:"varchar(255)"`
	Passwd string `xorm:"varchar(255) ->"`
	Headimgurl string `xorm:"TEXT"`
	Time time.Time `xorm:"DateTime"`
	LoginTime time.Time `xorm:"DateTime"`
	Status int64 `xorm:"int(2)"`
}
func (m *UserModel)GetUser(username string,passwd string) interface{} {
	users:=new(BlogUsers)
	orm:=core.ORM
	_,err:=orm.Where("user_name=? and passwd=?", username,passwd).Get(users)
	if err!=nil{
		log.Println(err)
	}
	return users
}