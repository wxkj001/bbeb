package model

import (
	"github.com/wxkj001/bbeb/core"
)

type ArticleModel struct {
	//*core.Model
}
type BlogArticle struct {
	Id      int64  `xorm:"int(11) pk autoincr unique index" json:"id"`
	UserId  int64  `xorm:"int(11)" json:"user_id"`
	Title   string `xorm:"varchar(255)" json:"title"`
	Content string `xorm:"Text" json:"content"`
	TagId   string `xorm:"varchar(255)" json:"tag_id"`
	Time    int64 `xorm:"int(11)" json:"time"`
	ETime   int64 `xorm:"int(11)" json:"e_time"`
	Status  int64  `xorm:"int(2)" json:"status"`
}

func (a *ArticleModel) List(page int, size int) (error, []BlogArticle, int) {
	var ba []BlogArticle
	//计算分页
	//pages := (page - 1) / (size + 1)
	err := core.ORM.Where("1=1").Find(&ba)

	return err, ba, page
}
