package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


type UserTest struct {
	Id      int	`json:"uid"`
	Name    string	`json:"uname"`
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	UserTest *UserTest `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	UserTest  *UserTest  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}


func init_1() {
	// 需要在init中注册定义的model
	orm.RegisterDriver("mysql", orm.DRMySQL)
	appName := beego.AppConfig.String("appname")
	fmt.Println(appName)
	mysql := beego.AppConfig.String("mysql")
	fmt.Println(mysql)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/MYBLOG?charset=utf8")

	orm.RegisterModel(new(UserTest), new(Post), new(Profile), new(Tag))

	// create table
	orm.RunSyncdb("default", false, true)
}


func Example_models() {

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	profile := new(Profile)
	profile.Age = 30

	user := new(UserTest)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))

}

func ReadUser() {

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	user := UserTest{Id: 1}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println("OK, we find the user!")
		fmt.Println(user.Id, user.Name)
	}
}


