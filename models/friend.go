package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Friend struct {
	Id int 			//Id:自增，主键
	UserId	int		//Master_id：主人id 外键 table1 not null
	FriendId	int 	//Friend _id：朋友id 外键 table1 not null
	Message string  //打招呼
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

//查询user
func GetFriendByName(name string) *Friend {
	o := orm.NewOrm()
	o.Using("default") // 默认使
	user := User{UserName:name}
	o.Read(&user)
	beego.Informational("当前查询申请列表的用户为：" + name)
	friend := Friend{UserId: user.Id}
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


func GetAppliesByName(name string) ([]*Friend, error) {
	o := orm.NewOrm()
	o.Using("default") // 默认使
	user := User{UserName:name}
	o.Read(&user, "user_name")
	var friends []*Friend
	var friendList []*Friend

	beego.Informational("查询申请列表的用户名：" + name + "; ID:")
	beego.Informational(user.Id)

	var maps []orm.Params
	//num, err := o.QueryTable("Friend").Filter("friend_id",user.Id).Filter("agree", false).GroupBy("user_id").OrderBy("-time").Values(&maps,"UserId", "MarkName", "Time", "Agree")
	num, err := o.QueryTable("Friend").Filter("FriendId", user.Id).Distinct().Values(&maps,"UserId", "FriendId", "Agree")


	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		beego.Informational("OK, we find the Applying user! Total number is " + string(num))
		beego.Informational(num)
		for _, value := range maps {
			//beego.Informational(value)
			//beego.Informational(value["FriendId"])
			//beego.Informational(value["UserId"])
			num1, err1 := o.QueryTable("Friend").Filter("friend_id", value["FriendId"]).Filter("agree", false).Filter("user_id", value["UserId"]).OrderBy("-time").Limit(1).All(&friends)

			//beego.Informational(num1)

			if err1 != nil {
				beego.Informational(err1)
				return nil, err1
			}
			if num1 > 0 {
				beego.Informational(friends[0].UserId, friends[0].FriendId, friends[0].MarkName, friends[0].Time, friends[0].Agree)
				friendList = append(friendList, friends[0])
			}

		}

		return friendList, err
	}
	return nil, err
}

func GetApplyByFriendAndUser(friendName , userName string) ([]*Friend, error) {
	o := orm.NewOrm()
	o.Using("default") // 默认使
	user := GetUserByName(userName)
	friend := GetUserByName(friendName)

	var friendApply []*Friend
	num, err := o.QueryTable("Friend").Filter("friend_id", user.Id).Filter("agree", false).Filter("user_id", friend.Id).OrderBy("-time").All(&friendApply)

	if err == nil {
		beego.Informational(num)
		return friendApply, err
	}

	return nil, err
}

//添加
func ApplyFriend(friend *Friend) error {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	friend.Time = time.Now()
	beego.Informational(time.Now())
	_, err := o.Insert(friend)
	if err != nil {
		beego.Informational("插入好友申请失败！")
		return err
	}
	return err
}

//删除(根据名称删除) 
func DeleteFriend(friend *Friend) error {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Delete(friend)
	if err != nil {
		beego.Informational(err)
	}

	return err
}


//更新
func UpdateFriend(friend *Friend) error {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Update(friend)
	if err != nil {
		beego.Informational(err)
	}
	return err
}