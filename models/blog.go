package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Blog struct {
	Id int				//Id:自增，主键
	Title string `orm:"size(1000)"`		//Title:文章标题
	//CatalogId int			//Catalog_id:文章所属目录id 外键 not null table5
	Content string		//Content：文章内容
	Time time.Time `orm:"auto_now;type(datetime)"`	//Time：完成时间
	Description string `orm:"size(100)"` 	//Description：描述
	KeyWord string `orm:size(100)`	//关键字
	UserId int  		//Author:文章作者id 外键 not null table1
	UserName string 	//UserName作者名
	Count int 			//Count:浏览次数
	From int			//From:文章来源（原创/转载）
	Good int			//Good_id：点赞数
	Bad int				//Bad_id: 踩数
	Catalog *Catalog `orm:"rel(fk)"`    //设置一对多关系
	Comment []*Comment `orm:"reverse(many)"` // 设置一对多的反向关系
}


//查询user
func GetBlog(id int) *Blog {
	o := orm.NewOrm()
	o.Using("default") // 默认使
	blog := Blog{Id: id}
	err := o.Read(&blog)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println("OK, we find the user!")
		return &blog
	}
	return nil
}


//添加
func InsertBlog(blog *Blog) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Insert(blog)
	if err != nil {
		return false
	}
	return true
}

//删除(根据名称删除) 
func DeleteBlog(blog *Blog) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Delete(blog)
	if err != nil {
		return false
	}
	return true
}


//更新
func UpdateBlog(blog *Blog) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Update(blog)
	if err != nil {
		return false
	}
	return true
}