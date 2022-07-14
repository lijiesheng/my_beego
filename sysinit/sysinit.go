package sysinit

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"zs403_mbook_copy/models"
	"zs403_mbook_copy/utils"
)


func sysinit()  {

	gob.Register(models.Member{})    // 这里要先注册，不然 encode 和 decode 会出错


	// 注册前端使用的函数 比如
	registerFunctions()
}

func registerFunctions()  {
	// 给前端提供的方法
	beego.AddFuncMap("showImg", utils.ShowImg)
}