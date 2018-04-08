package model

import (
	"time"
	"github.com/wxkj001/bbeb/core"
)

type ArticleModel struct {
	//*core.Model
}
type BlogArticle struct {
	Id int64 `xorm:"int(11) pk autoincr unique index"`
	UserId int64 `xorm:"int(11)"`
	Title string `xorm:"varchar(255)"`
	Content string `xorm:"Text"`
	TagId string `xorm:"varchar(255)"`
	Time time.Time `xorm:"DateTime"`
	Status int64 `xorm:"int(2)"`
}

func (this *ArticleModel)List(page int64) []BlogArticle {
	var ba []BlogArticle
	core.ORM.Where("1=1").Find(&ba)
	return ba
}
