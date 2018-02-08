package Admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wxkj001/bbeb/core"
	"log"
	"github.com/wxkj001/bbeb/model"
	"crypto/md5"
	"encoding/hex"
)

type Index struct {
	*core.FrameWork
}

func (this *Index)Index(g *gin.Context) {
	log.Println(core.SESSION.Get("AdminId"))
	if core.SESSION.Get("AdminId")==nil {
		g.Redirect(http.StatusMovedPermanently,"/Admin/login")
		return
	}
	g.HTML(http.StatusOK,"index.html",gin.H{})
}
func (this *Index)IndexInfo(g *gin.Context)  {
	g.HTML(http.StatusOK,"IndexInfo.html",gin.H{})
}
func (this *Index)Login(g *gin.Context)  {
	g.HTML(http.StatusOK,"login.html",gin.H{})
}
func (this *Index)LoginPost(g *gin.Context)  {
	models:=model.UserModel{}
	name:=g.PostForm("username")
	passwd:=g.PostForm("passwd")
	if name==""{
		g.JSON(http.StatusOK,gin.H{"status":1,"statusOk":"用户名错误"})
		return
	}
	if passwd==""{
		g.JSON(http.StatusOK,gin.H{"status":1,"statusOk":"密码未填写"})
		return
	}
	//md5加密
	m5 := md5.New()
	m5.Write([]byte(passwd))
	pass := m5.Sum(nil)
	passwd = hex.EncodeToString(pass)
	user:=models.GetUser(name,passwd)
	users:=user.(*model.BlogUsers)
	if users.Id==0{
		g.JSON(http.StatusOK,gin.H{"status":1,"statusOk":"密码错误"})
		return
	}
	this.Session().Set("AdminId",users.Id)
	this.Session().Save()
	g.JSON(http.StatusOK,gin.H{"status":0,"rows":users})
}
