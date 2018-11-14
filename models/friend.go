package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Friend struct {
	Id int 			//Id:自增，主键
	UserId	int		//Master_id：主人id 外键 table1 not null
	FriendId	int 	//Friend _id：朋友id 外键 table1 not null
	MarkName string	//Mark:备注（默认为用户名）
	Time time.Time 	//Time:申请时间
	Agree bool		//Agree:同意否
}



//查询user
func GetFriend(id int) *Friend {
	o := orm.NewOrm()
	o.Using("default") // 默认使
	friend := Friend{Id: id}
	err := o.Read(&friend)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println("OK, we find the user!")
		return &friend
	}
	return nil
}


//添加
func InsertFriend(friend *Friend) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Insert(friend)
	if err != nil {
		return false
	}
	return true
}

//删除(根据名称删除) 
func DeleteFriend(friend *Friend) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Delete(friend)
	if err != nil {
		return false
	}
	return true
}


//更新
func UpdateFriend(friend *Friend) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Update(friend)
	if err != nil {
		return false
	}
	return true
}