package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id int 			//Id:留言Id
	FromUserId int 	//From_author_id:留言者 外键 table1 非空
	ToUserId int	//To_author_id:留言对象 外键 table1 非空
	Content string	//Content：留言内容  not  null
	Time time.Time  //Time:留言时间
	Read bool		//Read:是否已读
}


//查询user
func GetMessage(id int) *Message {
	o := orm.NewOrm()
	o.Using("default") // 默认使
	message := Message{Id: id}
	err := o.Read(&message)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println("OK, we find the user!")
		return &message
	}
	return nil
}


//添加
func InsertMessage(message *Message) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Insert(message)
	if err != nil {
		return false
	}
	return true
}

//删除(根据名称删除) 
func DeleteMessage(message *Message) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Delete(message)
	if err != nil {
		return false
	}
	return true
}


//更新
func UpdateMessage(message *Message) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Update(message)
	if err != nil {
		return false
	}
	return true
}