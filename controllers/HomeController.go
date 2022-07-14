package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"zs403_mbook_copy/models"
)

type HomeController struct {
	BaseController
}

// 获取所有分类
func (c *HomeController) Index() {
	if cates,row, err := new(models.Category).GetCates(-1, 1); err == nil {
		c.Data["Cates"] = cates
		c.Data["Row"] = row
		fmt.Println("row==>", row)
	} else {
		beego.Error(err.Error())
	}
	c.TplName = "home/list.html"
}

func (c *HomeController) Index2(){
	c.TplName = "home/list.html"
}