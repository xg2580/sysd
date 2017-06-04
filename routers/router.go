package routers

import (
	"sysd/controllers"

	"github.com/astaxie/beego"
)

func init() {

	auth := beego.NewNamespace("/auth",
		//beego.NSRouter("/*", &controllers.AuthController{}, "get:Auth"),
		beego.NSRouter("/login", &controllers.AuthController{}, "get,post:Login"),
		beego.NSRouter("/register", &controllers.AuthController{}, "get:Register"),
		beego.NSRouter("/Logout", &controllers.AuthController{}, "get:Logout"),
	)
	beego.AddNamespace(auth)

	admin := beego.NewNamespace("/admin",
		beego.NSRouter("/index", &controllers.AdminController{}, "get,post:Get"),
		beego.NSRouter("/listmanagers", &controllers.AdminController{}, "get,post:ListManagers"),
		beego.NSRouter("/tomanage", &controllers.AdminController{}, "get,post:ToManagerPage"),
		beego.NSRouter("/addmanager", &controllers.AdminController{}, "get,post:AddManager"),
		beego.NSRouter("/delete", &controllers.AdminController{}, "get,post:DeleteManager"),
	)
	beego.AddNamespace(admin)
}
