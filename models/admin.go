package models

import (
	"errors"
	"sysd/common"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

func AddManager(admin *common.ShopAdmin) (err error) {
	var b bool
	v := validation.Validation{}
	b, err = v.Valid(admin)
	if err != nil {
		logs.Info(err)
		return
	}
	if !b {
		for _, e := range v.Errors {
			logs.Error(e.Key, e.Message)
			err = e
		}
		return
	}
	if a := common.GetAdminByInfo("adminuser", admin.Adminuser, 0); a != nil {
		return errors.New(admin.Adminuser + " is have register")
	}
	if a := common.GetAdminByInfo("adminemail", admin.Adminemail, 0); a != nil {
		return errors.New(admin.Adminemail + " is have register")
	}

	o := orm.NewOrm()
	admin.Createtime = time.Now().Format("2006-01-02 15:04:05")
	_, err = o.Insert(admin)
	return err
}

func GrepAdmins() (admins []*common.ShopAdmin, err error) {
	o := orm.NewOrm()
	sql := "select * from shop_admin"
	_, err = o.Raw(sql).QueryRows(&admins)
	if err != nil {
		return
	}
	return
}
func DeleteManager(id int) (err error) {
	o := orm.NewOrm()
	_, err = o.Delete(&common.ShopAdmin{Adminid: id})
	return
}
