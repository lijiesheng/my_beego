package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
	"zs403_mbook_copy/models"
)

type DocumentController struct {
	BaseController
}

//图书目录&详情页
// url : /books/:key  请求方式 任意
func (c *DocumentController) Index() {
	token := c.GetString("token")
	identify := c.Ctx.Input.Param(":key")   // 获取路径中的参数
	if "" == identify {
		c.Abort("404")
	}
	tab := strings.ToLower(c.GetString("tab"))

	bookResult := c.getBookData(identify, token)

}

// 获取图书内容并判断权限
func (c *DocumentController) getBookData (identify ,token string) *models.BookData {
	book, err := (&models.Book{}).Select("identify", identify)
	if err != nil {
		beego.Error(err)
		c.Abort("404")
	}
	// 私有文档
	if book.PrivatelyOwned == 1 && !c.Member.IsAdministrator() {
		isOk := false
		if c.Member != nil {
			(&models.Relationship{}).SelectRoleId(book.BookId, c.Member.MemberId)
			if err == nil {
				isOk = true
			}
		}


	}
}

