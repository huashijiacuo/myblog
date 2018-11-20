package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
	//设置主键且为自动增长（可以不设置，默认就是这样）
	Id int `orm:"pk;auto" form:"-"`
	UserName string `orm:"size(100)" form:"username"`
	Sex int `form:"gender"`
	Email string `orm:"size(100)" form:"Email"`
	PassWord string `form:"pwd"`
	Catalog []*Catalog `orm:"reverse(many)"` // 设置一对多的反向关系

}

//查询user
func GetUserById(id int) *User {
	o := orm.NewOrm()
	o.Using("default") // 默认使
	user := User{Id: id}
	err := o.Read(&user)

	if err == orm.ErrNoRows {
		beego.Informational("查询不到")
	} else if err == orm.ErrMissPK {
		beego.Informational("找不到主键")
	} else {
		beego.Informational("OK, we find the user!")
		return &user
	}
	return nil
}

//查询user
func GetUserByName(userName string) *User {
	o := orm.NewOrm()
	o.Using("default") // 默认使
	user := User{UserName: userName}
	err := o.Read(&user, "user_name")

	if err == orm.ErrNoRows {
		beego.Informational("查询不到")
	} else if err == orm.ErrMissPK {
		beego.Informational("找不到主键")
	} else {
		beego.Informational("OK, we find the user!")
		return &user
	}
	return nil
}

//添加
func (u *User) InsertUser(user *User) (message string, err error)  {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	beego.Informational("开始创建用户，并尝试将其存入数据库！")
	qs := o.QueryTable("user")
	err = qs.Filter("user_name", user.UserName).One(user)
	if err == nil {
		beego.Informational("创建失败，用户名已存在！！")
		beego.Informational(err)
		message = "the username is exist"
		err = errors.New("the username is exist")
		return message, err
	}
	_, err = o.Insert(user)
	if err != nil {
		beego.Informational("创建失败！")
		message = "insert failed!"
		return message, err
	}
	beego.Informational("恭喜用户名创建成功，已存入数据库！")
	message = "success! Create a new user!"
	return message,nil
}

//删除(根据名称删除) 
func (u *User) DeleteUser(user *User) error {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Delete(user)
	return err
}


//更新
func (u *User) UpdateUser(user *User) error {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Update(user)
	return err
}

func (u *User) String() string {
	return "用户名：" + u.UserName
}


/*
func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/MYBLOG?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}


func user_example() {
	orm.Debug = true
	o := orm.NewOrm()

	user := User{Name: "slene"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}*/


