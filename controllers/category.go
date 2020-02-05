package controllers

import (
	"code.bytedance.com/beeblog/models"
	"code.bytedance.com/beeblog/utils"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "add":
		categoryName := c.Input().Get("categoryname")
		if len(categoryName) == 0 {
			break
		}
		err := models.AddCategory(categoryName)
		if err != nil {
			beego.Error("add category_name error, category_name=", categoryName, "err=", err)
		}
		c.Redirect("/category", 302)
		return

	case "del":
		categoryId := c.Input().Get("id")
		if len(categoryId) == 0 {
			break
		}
		err := models.DelCategory(categoryId)
		if err != nil {
			beego.Error("delete category by Id failed, id=", categoryId)
		}
		c.Redirect("/category", 302)
		return
	}

	c.TplName = "category.html"
	c.Data["IsCategory"] = true
	c.Data["IsLogin"] = utils.CheckAccountCookie(c.Ctx)
	categoryList, err := models.GetAllCategory()
	if err != nil {
		beego.Error("get category failed, err=", err)
	}
	c.Data["Categories"] = categoryList
}
