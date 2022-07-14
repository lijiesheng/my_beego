package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Category struct {
	Id     int
	Pid    int    //分类id
	Title  string `orm:"size(30);unique"` // 唯一
	Intro  string //介绍
	Icon   string // 图标路径
	Cnt    int  //统计分类下图书
	Sort   int  //排序
	Status bool //状态，true 显示
}

/**
	表名
 */
func TNCategory() string{
	return "md_category"
}

// orm 回调
func (m *Category) TableName() string {
	return TNCategory()
}


// 获取所有分类
// -1 获取全部分类
func (m *Category) GetCates(pid , status int) (cates []Category, rows int64, err error) {
	sql := "select id, pid, title, intro, icon,  cnt,sort, status from md_category where 1 = 1 "
	if pid > -1 {
		sql += " and pid = " + strconv.Itoa(pid)
	}
	if status == 0 || status == 1 {
		sql += " and status = " + strconv.Itoa(status)
	}
	sql += " order by status desc , sort asc , title asc"
	newOrm := orm.NewOrm()
	rows, err = newOrm.Raw(sql).QueryRows(&cates)
	return
}
