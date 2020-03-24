package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

func AddCategory(categoryName string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title:     categoryName,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}
	// 先在表中看是否已经有categoryName了
	cateTable := o.QueryTable("category")
	err := cateTable.Filter("title", categoryName).One(cate)
	if err == nil {
		return nil
	}

	_, err = o.Insert(cate)
	if err != nil {
		beego.Error("insert new category to orm error, "+
			"categoryName=", categoryName, "err=", err)
		return err
	}
	return nil
}

func GetAllCategory() (categoryList []*Category, err error) {
	o := orm.NewOrm()

	cateTable := o.QueryTable("category")
	_, err = cateTable.All(&categoryList)
	if err != nil {
		beego.Error("get all category from orm failed, err=", err)
	}
	return categoryList, err
}

func DelCategory(categoryId string) error {
	o := orm.NewOrm()
	categoryId64, err := strconv.ParseInt(categoryId, 10, 64)
	if err != nil {
		beego.Error("categoryId is not legal, categoryId=", categoryId)
		return err
	}
	cate := &Category{
		Id: categoryId64,
	}

	cateTable := o.QueryTable("category")
	err = cateTable.Filter("id", categoryId).One(cate)
	if err != nil {
		beego.Error("categoryId is not exists, categoryId=", categoryId)
		return err
	}

	_, err = cateTable.Filter("id", categoryId).Delete()
	if err != nil {
		beego.Error("delete category by Id failed, categoryId=", categoryId)
		return err
	}
	return nil
}
