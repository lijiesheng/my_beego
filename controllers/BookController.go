package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"html/template"
	"zs403_mbook_copy/common"
	"zs403_mbook_copy/models"
	"zs403_mbook_copy/utils"
)

type BookController struct {
	BaseController
}

// 我的图书页面
func (c *BookController) Index() {
	pageIndex, _:= c.GetInt("page", 1)
	private, _ := c.GetInt("private", 1)  //默认私有
	books, totalCount, err := (&models.Book{}).SelectPage(pageIndex, common.PageSize, c.Member.MemberId, private)
	if err != nil {
		logs.Error("BookController.Index => ", err)
		c.Abort("404")
	}
	if totalCount > 0 {
		// 调用 BookController 方法中的 Index 方法
		c.Data["PageHtml"] = utils.NewPaginations(common.RollPage, totalCount, common.PageSize, pageIndex, beego.URLFor("BookController.Index"), fmt.Sprintf("&private=%v", private))
	} else {
		c.Data["PageHtml"] = ""
	}
	//封面图片
	for idx, book := range books {
		book.Cover = utils.ShowImg(book.Cover, "cover")
		books[idx] = book
	}
	b, err := json.Marshal(books)
	if err != nil || len(books) <= 0 {
		c.Data["Result"] = template.JS("[]")
	} else {
		c.Data["Result"] = template.JS(string(b))
	}
	c.Data["Private"] = private
	c.TplName = "book/index.html"
}

// 设置图书页面
func (c *BookController) Setting() {

	key := c.Ctx.Input.Param(":key")

	if key == "" {
		c.Abort("404")
	}

	book, err := (&models.BookData{}).SelectByIdentify(key, c.Member.MemberId)
	if err != nil && err != orm.ErrNoRows {
		c.Abort("404")
	}

	//需管理员以上权限
	if book.RoleId != common.BookFounder && book.RoleId != common.BookAdmin {
		c.Abort("404")
	}

	if book.PrivateToken != "" {
		book.PrivateToken = c.BaseUrl() + beego.URLFor("DocumentController.Index", ":key", book.Identify, "token", book.PrivateToken)
	}

	//查询图书分类
	if selectedCates, rows, _ := new(models.BookCategory).SelectByBookId(book.BookId); rows > 0 {
		var maps = make(map[int]bool)
		for _, cate := range selectedCates {
			maps[cate.Id] = true
		}
		c.Data["Maps"] = maps
	}

	c.Data["Cates"], _, _ = (&models.Category{}).GetCates(-1, 1)
	c.Data["Model"] = book
	c.TplName = "book/setting.html"
}