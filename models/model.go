package models

import "github.com/astaxie/beego/orm"

func init() {
	// 需要在init中注册定义的model
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/MYBLOG?charset=utf8")

	orm.RegisterModel(new(User), new(Blog), new(Catalog), new(Comment), new(Diary), new(Friend), new(Message))

	// create table
	orm.RunSyncdb("default", false, true)

}