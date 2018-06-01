package home

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/wxkj001/bbeb/core"
)

type PublicController struct {
	*core.FrameWork
}

func (this *PublicController) Index(e echo.Context) error {
	env := map[string]int{
		"install": 0,
		"data":    0,
	}
	_, err := os.Stat("./data/install.lock")
	if err == nil {
		env["install"] = 1
	}
	file_info, err := os.Stat("./data")
	if err != nil {
		log.Panicln("data文件夹不存在", err)
		return err
	}
	file_mode := file_info.Mode()
	perm := file_mode.Perm()
	log.Println("file_mode:", file_mode)
	log.Println("perm:", perm)
	return e.JSON(http.StatusOK, core.H{"code": 0, "msg": "", "data": env})
}

type form struct {
	Datebase string `json:"datebase" form:"datebase" query:"datebase"`
	Host     string `json:"host" form:"host" query:"host"`
	User     string `json:"user" form:"user" query:"user"`
	Pass     string `json:"pass" form:"pass" query:"pass"`
	Port     string `json:"port" form:"port" query:"port"`
}

/*
	修改app.ini配置文件
*/
func (this *PublicController) SetForm(e echo.Context) (err error) {
	conf := this.Config()
	f := new(form)
	if err = e.Bind(f); err != nil {
		return
	}
	//接收数据
	conf.Section("database").NewKey("database", f.Datebase)
	conf.Section("database").NewKey("host", f.Host)
	conf.Section("database").NewKey("username", f.User)
	conf.Section("database").NewKey("password", f.Pass)
	conf.Section("database").NewKey("port", f.Port)
	//修改ini文件内容
	err = conf.SaveTo("./data/app.ini")
	if err != nil {
		log.Println("文件写入失败:", err)
		return
	}
	conf.Reload()
	core.ORM.Close()
	core.ORMLoad()
	return e.JSON(http.StatusOK, core.H{"code": 0, "msg": "修改成功！"})
}

type adminConf struct {
	AdminUser string `json:"adminuser" form:"adminuser" query:"adminuser"`
	AdminPass string `json:"adminpass" form:"adminpass" query:"adminpass"`
}

func (this *PublicController) SetAdmin(e echo.Context) (err error) {

	return e.JSON(http.StatusOK, core.H{"code": 0, "msg": "添加成功！"})
}
