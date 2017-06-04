package common

import (
	"github.com/astaxie/beego/config"
)

var SysdConfig config.Configer

func init() {
	conf, err := config.NewConfig("ini", "conf/sysd.conf")
	if err != nil {
		panic(err)
	}
	SysdConfig = conf
}
