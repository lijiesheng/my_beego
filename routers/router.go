package routers

import (
	"github.com/astaxie/beego"
	"zs403_mbook_copy/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.HomeController{}, "get:Index")  // 获取 Category 所有信息
	beego.Router("/2", &controllers.HomeController{}, "get:Index2") // 跳转到 "home/list.html"
	beego.Router("/explore", &controllers.ExploreController{}, "get:Index") // 获取某个分类下的所有读书

	// 读书
	beego.Router("/books/:key", &controllers.DocumentController{}, "*:Index")



	//login
	beego.Router("/login", &controllers.AcountController{}, "*:Login")   // 登录
	beego.Router("/regist", &controllers.AcountController{}, "*:Regist")  // 注册
	beego.Router("/logout", &controllers.AcountController{}, "*:Logout")  // 注册
	beego.Router("/doregist", &controllers.AcountController{}, "post:DoRegist")  // 查表
}
