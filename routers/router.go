package routers

import (
	"code.bytedance.com/beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	// 自动路由功能，函数名和路径保持一致，例如/topic就会默认调topicController
	// /topic/add自动调TopicController下的add方法
	beego.AutoRouter(&controllers.TopicController{})
	beego.AutoRouter(&controllers.CommentController{})
}
