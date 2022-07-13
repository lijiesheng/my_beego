package routers

import (
	"github.com/astaxie/beego"
	"zs403_mbook_copy/controllers"

)

func init() {
    beego.Router("/", &controllers.MainController{})
}
