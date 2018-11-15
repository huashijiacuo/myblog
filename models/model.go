package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDriver("mysql", orm.DRSqlite)

	orm.RegisterDataBase("default", "mysql", "root:Love_zhi0928@tcp(127.0.0.1:3306)/MYBLOG?charset=utf8")

	orm.RegisterModel(new(User), new(Blog), new(Catalog), new(Comment), new(Diary), new(Friend), new(Message))

	// create table
	orm.RunSyncdb("default", false, true)

}