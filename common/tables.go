package common

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(ShopUser), new(ShopAdmin))

}

type ShopUser struct {
	Userid     int `orm:"pk"`
	Username   string
	Userpass   string
	Useremail  string
	Openid     int
	Createtime int
}

/*
CREATE TABLE `shop_admin` (
  `adminid` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `adminuser` varchar(32) NOT NULL DEFAULT '' COMMENT '管理员账号',
  `adminpass` char(64) NOT NULL DEFAULT '',
  `adminemail` varchar(50) NOT NULL DEFAULT '' COMMENT '管理员电子邮箱',
  `logintime` varchar(32) NOT NULL DEFAULT ''  COMMENT '登录时间',
  `loginip` varchar(32) NOT NULL DEFAULT '' COMMENT '登录IP',
  `createtime` varchar(32) NOT NULL DEFAULT '' COMMENT '创建时间',
  PRIMARY KEY (`adminid`),
  UNIQUE KEY `shop_admin_adminuser_adminpass` (`adminuser`,`adminpass`),
  UNIQUE KEY `shop_admin_adminuser_adminemail` (`adminuser`,`adminemail`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
*/

type ShopAdmin struct {
	Adminid    int    `orm:"pk"`
	Adminuser  string `valid:"Required"`
	Adminpass  string `valid:"Required"`
	Adminemail string `valid:"Email; MaxSize(100)"`
	Logintime  string
	Loginip    string
	Createtime string
}
