package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Comment struct {
	Id int 			//Id:自增，主键
	Blog	int		//Blog_id:评论的博客id 外键 非空table2
	Content	string	//Content：评论内容
	Time time.Time	//Time：评论时间
	UserId int		//Author_id：发表评论的作者id 外键 table1 非空
	UserName int	//user_name:评论者的用户名
	ToUserId int	//To_author_id:评论的文章的作者 外键 非空 table1
	blog *Blog `orm:"rel(fk)"`    //设置一对多关系
}

//查询user
func GetComment(id int) *Comment {
	o := orm.NewOrm()
	o.Using("default") // 默认使
	comment := Comment{Id: id}
	err := o.Read(&comment)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println("OK, we find the user!")
		return &comment
	}
	return nil
}


//添加
func InsertComment(comment *Comment) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Insert(comment)
	if err != nil {
		return false
	}
	return true
}

//删除(根据名称删除) 
func DeleteComment(comment *Comment) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Delete(comment)
	if err != nil {
		return false
	}
	return true
}


//更新
func UpdateComment(comment *Comment) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Update(comment)
	if err != nil {
		return false
	}
	return true
}