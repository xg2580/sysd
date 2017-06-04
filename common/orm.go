package common

import (
	//"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func GetAdminByInfo(key, u string, id int) (admin *ShopAdmin) {
	var err error
	o := orm.NewOrm()
	switch key {
	case "adminid":
		admin = &ShopAdmin{Adminid: id}
	case "adminuser":
		admin = &ShopAdmin{Adminuser: u}
		err = o.Read(admin, "adminuser")
	case "adminemail":
		admin = &ShopAdmin{Adminemail: u}
		err = o.Read(admin, "adminemail")
	}
	if err = o.Read(admin); err != nil {
		return nil
	}
	return admin
}
