package model

import (
	"github.com/wxkj001/bbeb/core"
	"time"
)

type TagsModel struct {
	*core.Model
}

type BlogTags struct {
	Id int64 `xorm:"int(11) pk autoincr unique index"`
	TagName string `xorm:"varchar(255)"`
	Time time.Time `xorm:"DateTime"`
	Status int64 `xorm:"int(2)"`
}