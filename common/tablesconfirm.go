package common

import (
	"strings"

	"github.com/astaxie/beego/validation"
)

func (ad *ShopAdmin) Valid(v *validation.Validation) {
	if strings.Index(ad.Adminuser, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}
