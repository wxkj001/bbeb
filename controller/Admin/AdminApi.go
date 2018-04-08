package Admin

import (
	"github.com/wxkj001/bbeb/core"
	"github.com/labstack/echo"
	"net/http"
	"github.com/wxkj001/bbeb/model"
	"crypto/md5"
	"encoding/hex"
)

type AdminApiConstroller struct {
	*core.FrameWork
}

func (this *AdminApiConstroller)Login(e echo.Context) error {
	type loginFrom struct {
		Name string `json:"name" form:"name"`
		Pass string `json:"pass" form:"pass"`
	}
	lf:=new(loginFrom)
	//绑定用户数据
	if err := e.Bind(lf); err != nil {
		return e.JSON(http.StatusOK,core.H{"code":1,"msg":""})
	}
	um:=model.UserModel{}
	//获取用户数据
	user,_:=um.GetUser(lf.Name).(*model.BlogUsers)
	if user.Id ==0{
		return e.JSON(http.StatusOK,core.H{"code":1,"msg":"账号不存在"})
	}
	//MD5加密
	m5:=md5.New()
	m5.Write([]byte(lf.Pass))
	cipherStr := m5.Sum(nil)
	if user.Passwd!=hex.EncodeToString(cipherStr){
		return e.JSON(http.StatusOK,core.H{"code":1,"msg":"密码错误"})

	}
	return e.JSON(http.StatusOK,core.H{"code":0,"msg":"ok","data":user})
}
func (this *AdminApiConstroller)ArticleList (e echo.Context) error {
	var ba []model.BlogArticle
	am:=model.ArticleModel{}
	ba=am.List(1)
	return e.JSON(http.StatusOK,core.H{"code":0,"msg":"ok","data":ba})
}

