package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Diary struct {
	Id int			//Id:自增，主键
	Title string	//Title：日记标题
	Catalog int		//Catalog_id：目录分类 外键 table5
	Content string	//Content：内容
	Time time.Time 	//Time：写作时间
	Description string `orm:"size(100)"` 	//Description：描述
	KeyWord string `orm:size(100)`	//关键字
	From int			//From:文章来源（原创/转载）
	UserId int		//Author_id：作者 外键 非空 table1
	UserName string //User:作者名
	catalog *Catalog `orm:"rel(fk)"`    //设置一对多关系
}


//查询user
func GetDiary(id int) *Diary {
	o := orm.NewOrm()
	o.Using("default") // 默认使
	diary := Diary{Id: id}
	err := o.Read(&diary)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println("OK, we find the user!")
		return &diary
	}
	return nil
}


//添加
func InsertDiary(diary *Diary) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Insert(diary)
	if err != nil {
		return false
	}
	return true
}

//删除(根据名称删除) 
func DeleteDiary(diary *Diary) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Delete(diary)
	if err != nil {
		return false
	}
	return true
}


//更新
func UpdateDiary(diary *Diary) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Update(diary)
	if err != nil {
		return false
	}
	return true
}