package admin

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/wxkj001/bbeb/core"
	"github.com/wxkj001/bbeb/model"
)

type APIConstroller struct {
	*core.FrameWork
}

func (a *APIConstroller) Login(e echo.Context) error {
	type loginFrom struct {
		Name string `json:"name" form:"name"`
		Pass string `json:"pass" form:"pass"`
	}
	lf := new(loginFrom)
	//绑定用户数据
	if err := e.Bind(lf); err != nil {
		return e.JSON(http.StatusOK, core.H{"code": 1, "msg": "数据错误"})
	}
	um := model.UserModel{}
	//获取用户数据
	user, _ := um.GetUser(lf.Name).(*model.BlogUsers)
	if user.Id == 0 {
		return e.JSON(http.StatusOK, core.H{"code": 1, "msg": "账号不存在"})
	}
	//MD5加密
	m5 := md5.New()
	m5.Write([]byte(lf.Pass))
	cipherStr := m5.Sum(nil)
	if user.Passwd != hex.EncodeToString(cipherStr) {
		return e.JSON(http.StatusOK, core.H{"code": 1, "msg": "密码错误"})

	}
	time.LoadLocation("Local")
	t := time.Now()
	buser := new(model.BlogUsers)
	buser.LoginTime = t
	buser.Device = e.Request().Header.Get("Device")
	_, err := core.ORM.Where("id= ?", user.Id).Update(buser)
	log.Println(err)
	/*
	 *娱乐密码
	 */
	p := strings.Count(lf.Pass, "") - 1
	user.Passwd = ""
	for i := 0; i < p; i++ {
		user.Passwd += "▉"
	}
	//
	sess, _ := session.Get("token", e)
	log.Printf("%+v\n", sess)
	if sess.IsNew == false {
		sess.Values["id"] = user.Id
		sess.Save(e.Request(), e.Response())
	}
	return e.JSON(http.StatusOK, core.H{"code": 0, "msg": "ok", "data": user})
}
func (a *APIConstroller) ArticleList(e echo.Context) error {
	type ListP struct {
		Page int `json:"page" form:"page" query:"page"`
		Size int `json:"size" form:"size" query:"size"`
	}

	//绑定数据
	lp := new(ListP)
	if err := e.Bind(lp); err != nil {
		return e.JSON(http.StatusOK, core.H{"code": 1, "msg": "数据绑定错误"})
	}
	//return e.JSON(http.StatusOK, core.H{"code": 1, "msg": "e.Bind(lp)==nil","Validate":e.Bind(lp),"OK":e.Bind(lp)==nil})
	sess, _ := session.Get("token", e)
	log.Printf("%+v", sess.Values["id"])
	var ba []model.BlogArticle
	am := model.ArticleModel{}
	err, ba, p := am.List(lp.Page, lp.Size)
	log.Println(err)

	return e.JSON(http.StatusOK, core.H{"code": 0, "msg": "ok", "data": ba, "page": p})
}
func (a *APIConstroller) ArticleInfo(e echo.Context) error {
	id := e.Param("id")
	var art model.BlogArticle
	core.ORM.Where("id= ?", id).Get(&art)
	log.Println(time.Now().Format("2006-01-02 15:04:05"))
	return e.JSON(http.StatusOK, core.H{"code": 0, "msg": "ok", "data": art})
}
func (a *APIConstroller) ArticleADD(e echo.Context) error {
	type Form struct {
		Title   string `json:"title" form:"title"`
		Content string `json:"content" form:"content"`
		TagID   string `json:"tag_id" form:"tag_id"`
	}
	form := new(Form)
	if err := e.Bind(form); err != nil {
		log.Println(err)
		return e.JSON(http.StatusOK, core.H{"code": 1, "msg": "数据错误"})
	}
	var article model.BlogArticle
	article.Title = form.Title
	article.Content = form.Content
	article.TagId = form.TagID
	article.Status = 1
	article.Time = time.Now().Unix()
	Id, err := core.ORM.Insert(&article)
	if err != nil {
		return e.JSON(http.StatusOK, core.H{"code": 1, "msg": "添加失败"})
	}
	return e.JSON(http.StatusOK, core.H{"code": 0, "msg": "ok", "data": Id})
}
func (a *APIConstroller) ArticleEDIT(e echo.Context) error {
	id := e.Param("id")
	type Form struct {
		Title   string `json:"title" form:"title"`
		Content string `json:"content" form:"content"`
		TagID   string `json:"tag_id" form:"tag_id"`
	}
	form := new(Form)
	if err := e.Bind(form); err != nil {
		log.Println(err)
		return e.JSON(http.StatusOK, core.H{"code": 1, "msg": "数据错误"})
	}
	var article model.BlogArticle
	article.Title = form.Title
	article.Content = form.Content
	article.TagId = form.TagID
	article.ETime = time.Now().Unix()
	Id, err := core.ORM.Where("id= ?", id).Update(&article)
	if err != nil {
		return e.JSON(http.StatusOK, core.H{"code": 1, "msg": "修改失败"})
	}
	return e.JSON(http.StatusOK, core.H{"code": 0, "msg": "ok", "data": Id})
}
func (a *APIConstroller) ArticleDEL(e echo.Context) error {
	id := e.Param("id")
	var article model.BlogArticle
	article.Status = -1
	Id, err := core.ORM.Where("id= ?", id).Update(&article)
	if err != nil {
		return e.JSON(http.StatusOK, core.H{"code": 1, "msg": "修改失败"})
	}
	return e.JSON(http.StatusOK, core.H{"code": 0, "msg": "ok", "data": Id})
}
