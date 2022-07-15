package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
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


// 查询分类
func (m *Category) Find(id int) (cate *Category, err error) {

	// todo 这里没有写 status
	sql := "select id, pid, title, intro, icon,  cnt,sort, status from md_category where id = " + strconv.Itoa(id)
	newOrm := orm.NewOrm()
	err = newOrm.Raw(sql).QueryRow(&cate)
	if err != nil && "<QuerySeter> no row found" != err.Error() {
		return nil, err
	}
	return
}

// 单个查询
func (m *Category) GetBySql(sql string) (cate *Category, err error) {
	newOrm := orm.NewOrm()
	err = newOrm.Raw(sql).QueryRow(&cate)
	if err != nil && "<QuerySeter> no row found" != err.Error() {
		return nil, err
	}
	return
}


// 新增分类
func (m *Category) InsertMulti(pid int, cates string) (err error) {
	slice := strings.Split(cates, "\n")
	if len(slice) == 0 {
		return
	}

	o := orm.NewOrm()
	//开启事务
	err = o.Begin()
	var sqlError error
	for _,item := range slice {
		if item = strings.TrimSpace(item); item != "" {
			var cate = Category{
				Pid : pid,
				Title : item,
				Status: true,
			}
			if o.Read(&cate, "title"); cate.Id == 0 {
				_, sqlError = o.Insert(&cate)
			}
			if sqlError != nil {
				break;
			}
		}
	}
	if sqlError != nil {
		logs.Error("execute transaction's sql fail, rollback.", err)
		err = o.Rollback()
		if err != nil {
			logs.Error("roll back transaction failed", err)
		}
		return sqlError
	} else {
		err = o.Commit()
		if err != nil {
			logs.Error("commit transaction failed.", err)
		}
		return
	}
}


// 删除分类
func (m *Category) Delete (id int) (err error) {
	var cate = Category{Id : id}
	o := orm.NewOrm()
	if err = o.Read(&cate); cate.Cnt > 0 { //当前分类下文档图书数量不为0，不允许删除
		return errors.New("删除失败，当前分类下的问下图书不为0，不允许删除")
	}
	if _, err = o.Delete(&cate, "id"); err != nil {
		return
	}
	_, err = o.QueryTable(TNCategory()).Filter("pid", id).Delete()
	if err != nil { //删除分类图标
		return err
	}
	return
}

//更新分类字段  todo 这个很重要
func (m *Category) UpdateField (id int, field, val string) (err error) {
	_, err = orm.NewOrm().QueryTable(TNCategory()).Filter("id", id).Update(orm.Params{field: val})
	return
}

// 统计分类书籍
var counting = false

type Count struct {
	Cnt        int
	CategoryId int
}

func CountCategory() {
	if counting {
		return
	}
	counting = true

	defer func() {
		counting = false
	}()

	var count []Count
	o := orm.NewOrm()
	sql := "select count(bc.id) cnt, bc.category_id from " + TNBookCategory() + " bc left join " + TNBook() +
		" b on b.book_id=bc.book_id where b.privately_owned=0 group by bc.category_id"
	o.Raw(sql).QueryRows(&count)
	if len(count) == 0 {
		return
	}
	// cnt 1   category_id 1
	// cnt 1   category_id 4
	var cates []Category

	// 查询所有的 md_category 取出的字段是 id pid cnt
	o.QueryTable(TNCategory()).All(&cates, "id", "pid", "cnt")
	if len(cates) == 0 {
		return
	}
	// 初始化所有的 md_category 的 cnt 是0
	o.QueryTable(TNCategory()).Update(orm.Params{"cnt": 0})
	// todo 这个很重要
	cateChild := make(map[int]int)   // 这个有啥用
	for _, item := range count {
		if item.Cnt > 0 {
			cateChild[item.CategoryId] = item.Cnt
			_, err := o.QueryTable(TNCategory()).Filter("id", item.CategoryId).Update(orm.Params{"cnt" : item.Cnt})
			if err != nil {
				return
			}
		}
	}
}













