package controllers

import (
	. "sysd/common"
	//	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"github.com/astaxie/beego/orm"
)

type AuthController struct {
	beego.Controller
}

func (this *AuthController) Login() {

}

func (this *AuthController) Register() {

	user := this.GetString("email-user")
	logs.Info("============================", user)
	//	o := orm.NewOrm()
	//	u := new(common.ShopUser)
	//	u.Userid = 89
	//	u.Username = "xg"
	//	u.Userpass = "123"
	//	u.Useremail = "teaaast@qq.com"
	//	u.Openid = 23
	//	u.Createtime = int(time.Now().Unix())
	//	a, b := o.Insert(u)
	//	logs.Info(a, b)
	//	fmt.Println("register")
	var msg Message
	//defer this.ServeJSON()
	this.Data["test"] = &msg
	msg = Message{true, "scnsdbcnijdsnvindsivnoisdnvciondsivncsdonvodsndsvsi 添加成功"}
	this.TplName = "front_page/auth.html"
}

func (this *AuthController) Logout() {
	//fmt.Println("logout")

	this.TplName = "front_page/auth.html"
}
