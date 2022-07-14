package routers

import (
	"github.com/astaxie/beego"
	"zs403_mbook_copy/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.HomeController{}, "get:Index")  // 获取 Category 所有信息
	beego.Router("/2", &controllers.HomeController{}, "get:Index2") // 跳转到 "home/list.html"

	//beego.Router("/explore", )
}
