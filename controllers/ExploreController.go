package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"math"
	"strconv"
	"zs403_mbook_copy/models"
	"zs403_mbook_copy/utils"
)

type ExploreController struct {
	BaseController
}

func (c *ExploreController) Index() {
	var (
		cid   int // 分类id
		cate  models.Category
		urlPrefix = beego.URLFor("ExploreController.Index")  // 链接前缀  给前端使用
	)
	if cid, _ = c.GetInt("cid"); cid > 0 {
		var err error
		cateModel :=  &models.Category{}
		cate, err = cateModel.Find(cid)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Data["Cate"] = cate
	}
	c.Data["Cid"] = cid
	c.TplName = "explore/index.html"

	pageIndex, _ := c.GetInt("page", 1)  // 获取 page 的值，默认是1
	pageSize := 24

	books, totalCount , err :=
		(&models.Book{}).HomeData(pageIndex, pageSize, cid)
	if err != nil {
		beego.Error(err)
		c.Abort("404")
	}
	if totalCount > 0 {
		urlSuffix := ""
		if cid > 0 {
			urlPrefix = urlPrefix + "&cid=" + strconv.Itoa(cid)
		}
		html := utils.NewPaginations(4,totalCount, pageSize, pageIndex, urlPrefix, urlSuffix)
		c.Data["PageHtml"] = html
	} else {
		c.Data["PageHtml"] = ""
	}
	c.Data["TotalPages"] = int(math.Ceil(float64(totalCount) / float64(pageSize)))
	c.Data["Lists"] = books
}

