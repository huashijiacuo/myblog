package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Catalog struct {
	Id int	`orm:"pk;auto"`		//Id：自增，主键
	//UserId int 		//Author_id：目录创建者id 外键 table1
	CatalogName string		//Name:目录名称
	Blogs []*Blog `orm:"reverse(many)"` // 设置一对多的反向关系
	Diarys []*Diary `orm:"reverse(many)"` // 设置一对多的反向关系
	User *User `orm:"rel(fk)"`    //设置一对多关系
}


//查询user
func GetCatalog(id int) *Catalog {
	o := orm.NewOrm()
	o.Using("default") // 默认使
	catalog := Catalog{Id: id}
	err := o.Read(&catalog)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println("OK, we find the user!")
		return &catalog
	}
	return nil
}

func GetCatalogByName(cataName string) *Catalog {
	o := orm.NewOrm()
	o.Using("default") // 默认使
	catalog := Catalog{CatalogName: cataName}
	err := o.Read(&catalog)

	if err == orm.ErrNoRows {
		beego.Informational("查询不到")
	} else if err == orm.ErrMissPK {
		beego.Informational("找不到主键")
	} else {
		beego.Informational("OK, we find the user!")
		return &catalog
	}
	return nil
}

//添加
func InsertCatalog(catalog *Catalog) error {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Insert(catalog)
	if err != nil {
		beego.Informational("models.Catalog 插入数据失败！")
		return err
	}
	return err
}

//删除(根据名称删除) 
func DeleteCatalog(catalog *Catalog) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Delete(catalog)
	if err != nil {
		return false
	}
	return true
}


//更新
func UpdateCatalog(catalog *Catalog) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	_, err := o.Update(catalog)
	if err != nil {
		return false
	}
	return true
}
