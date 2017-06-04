package main

import (
	"sysd/common"
	_ "sysd/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	setStaticPath()
	setMysql()
	setLog()
}

func main() {
	//beego.SetLogFuncCall(true)
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}

func setStaticPath() {
	beego.SetStaticPath("admin_images", "views/admin_page/assets/images")
	beego.SetStaticPath("admin_js", "views/admin_page/assets/js")
	beego.SetStaticPath("admin_css", "views/admin_page/assets/css")

	beego.SetStaticPath("front_images", "views/front_page/assets/images")
	beego.SetStaticPath("front_js", "views/front_page/assets/js")
	beego.SetStaticPath("front_css", "views/front_page/assets/css")
}

func setMysql() {
	usr := common.SysdConfig.DefaultString("mysqlusr", "root")
	pass := common.SysdConfig.DefaultString("mysqlpas", "root123")
	s := usr + ":" + pass + "@/sysd?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", s, 30, 30)
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Error(err)
	}
}
func setLog() {
	log := logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"log/sysd.log"}`)
	log.EnableFuncCallDepth(true)
	log.SetLevel(7)
	//beego.BConfig.Log.FileLineNum = true

}
