package home

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/wxkj001/bbeb/core"
	"github.com/wxkj001/bbeb/model"
)

type APIConstroller struct {
	*core.FrameWork
}

func (a *APIConstroller) ArticleList(e echo.Context) error {
	type ListP struct {
		Page int `json:"page" form:"page"`
		Size int `json:"size" form:"size"`
	}
	//绑定数据
	lp := new(ListP)
	if err := e.Bind(lp); err != nil {
		return e.JSON(http.StatusOK, core.H{"code": 1, "msg": "数据错误"})
	}
	sess, _ := session.Get("token", e)
	log.Printf("%+v", sess.Values["id"])
	var ba []model.BlogArticle
	am := model.ArticleModel{}
	err, ba, p := am.List(lp.Page, lp.Size)
	log.Println(err)

	return e.JSON(http.StatusOK, core.H{"code": 0, "msg": "ok", "data": ba, "page": p})
}
