package controllers

import (
	. "sysd/common"
	"sysd/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
)

type AdminController struct {
	beego.Controller
	i18n.Locale
}

func (this *AdminController) Get() {
	this.TplName = "admin_page/views/default/index.html"
	this.Layout = "admin_page/views/layouts/layout.html"
}

func (this *AdminController) ToManagerPage() {
	this.TplName = "admin_page/views/manage/reg.html"
	this.Layout = "admin_page/views/layouts/layout.html"

}
func (this *AdminController) AddManager() {
	var msg Message
	var admin ShopAdmin
	this.Data["json"] = &msg
	defer this.ServeJSON()
	if err := this.ParseForm(&admin); err != nil {
		msg = Message{false, admin.Adminuser + " 添加失败 " + err.Error()}
		return
	}
	if err := models.AddManager(&admin); err != nil {
		msg = Message{false, admin.Adminuser + " 添加失败 " + err.Error()}
		return
	}
	msg = Message{true, admin.Adminuser + " 添加成功"}
}

func (this *AdminController) ListManagers() {
	var msg Message
	this.Data["adminsInfo"] = &msg
	adminsInfo, err := models.GrepAdmins()
	if err != nil {
		msg = Message{false, " 查询管理员列表失败 "}
		return
	}
	logs.Info(adminsInfo)
	msg = Message{true, adminsInfo}
	this.Layout = "admin_page/views/layouts/layout.html"
	this.TplName = "admin_page/views/manage/managers.html"
}

func (this *AdminController) DeleteManager() {
	var msg Message
	this.Data["deleteInfo"] = &msg
	adminid, err := this.GetInt("adminid")
	if err != nil {
		msg = Message{false, " 删除失败 "}
		logs.Error(err)
		return
	}
	if err = models.DeleteManager(adminid); err != nil {
		msg = Message{false, " 删除失败 "}
		logs.Error(err)
		return
	}
	msg = Message{false, " 删除成功 "}
	//this.Layout = "admin_page/views/layouts/layout.html"
	this.TplName = "admin_page/views/manage/managers.html"
}
