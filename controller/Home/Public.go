package Home

import (
	"github.com/wxkj001/bbeb/core"
	"net/http"
	"os"
	"log"
	"github.com/labstack/echo"
)

type PublicController struct{
	*core.FrameWork
}

func (this *PublicController)Index(e echo.Context) error {
	env:=map[string]int{
		"install":0,
		"data":0,
	}
	_, err := os.Stat("./data/install.lock")
	if err ==nil {
		env["install"]=1
	}
	file_info, err := os.Stat("./data")
	if err != nil {
		log.Panicln("data文件夹不存在", err)
		return err
	}
	file_mode:=file_info.Mode()
	perm := file_mode.Perm()
	log.Println("file_mode:", file_mode)
	log.Println("perm:", perm)
	return e.JSON(http.StatusOK,core.H{"code":0,"msg":"","data":env})
}

func (this *PublicController)SetForm(e echo.Context) error {
	conf:=this.Config()
	//conf.Section("database").NewKey("database",datebase)
	//conf.Section("database").NewKey("host",dateurl)
	//conf.Section("database").NewKey("username",user)
	//conf.Section("database").NewKey("password",pass)
	err:=conf.SaveTo("./data/app.ini")
	if err!=nil{
		log.Println("文件写入失败:",err)
		return err
	}
	return e.JSON(http.StatusOK,core.H{"code":0,"msg":"修改成功！"})
}
