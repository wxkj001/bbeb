package model

import (
	"github.com/wxkj001/bbeb/core"
	"time"
)

type ArticleModel struct {
	*core.Model
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
