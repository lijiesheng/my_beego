package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

type Book struct {
	BookId         int       `orm:"pk;auto" json:"book_id"`  // 自增主键 如果是id 自动是自增主键
	BookName       string    `orm:"size(500)" json:"book_name"`
	Identify       string    `orm:"size(100);unique" json:"identify"` //唯一标识
	OrderIndex     int       `orm:"default(0)" json:"order_index"`
	Description    string    `orm:"size(1000)" json:"description"`       //图书描述
	Cover          string    `orm:"size(1000)" json:"cover"`                        //封面地址
	Editor         string    `orm:"size(50)" json:"editor"`              //编辑器类型: "markdown"
	Status         int       `orm:"default(0)" json:"status"`            //状态:0 正常 ; 1 已删除
	PrivatelyOwned int       `orm:"default(0)" json:"privately_owned"`   // 是否私有: 0 公开 ; 1 私有
	PrivateToken   string    `orm:"size(500);null" json:"private_token"` // 私有图书访问Token
	MemberId       int       `orm:"size(100)" json:"member_id"`
	CreateTime     time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`
	// create_time     datetime                 not null,

	//auto_now 每次 model 保存时都会对时间自动更新;
	//auto_now_add 第一次保存时才设置时间   对于批量的 update 此设置是不生效的
	ModifyTime     time.Time `orm:"type(datetime);auto_now_add" json:"modify_time"`  //修改时间
	// modify_time     datetime                 not null,

	ReleaseTime    time.Time `orm:"type(datetime);" json:"release_time"`  //发布时间
	//  release_time    datetime                 not null,

	DocCount       int       `json:"doc_count"`                      //文档数量
	CommentCount   int       `orm:"type(int)" json:"comment_count"`
	Vcnt           int       `orm:"default(0)" json:"vcnt"`              //阅读次数
	Score          int       `orm:"default(40)" json:"score"`            //评分
	CntScore       int       //评分人数
	CntComment     int       //评论人数
	Author         string    `orm:"size(50)"`                      //来源
	AuthorURL      string    `orm:"column(author_url);size(1000)"` //来源链接
}

func TNBook() string {
	return "md_books"
}

func (m *Book) TableName() string {
	return TNBook()
}

// 获取分类下面的所有书籍
func (m *Book) HomeData(pageIndex, pageSize int, cid int, fields ...string)(books []Book, totalCount int, err error) {
	if len(fields) == 0 {
		fields = append(fields, "book_id", "book_name", "identify", "cover", "order_index")  // 数组
	}
	fieldStr := "b." + strings.Join(fields, " b, ") // 字符串
	sqlFmt := "select %v from " + TNBook() + " b left join " + TNBookCategory() + " c on b.book_id=c.book_id where c.category_id=" +
		strconv.Itoa(cid)
	sql := fmt.Sprintf(sqlFmt, fieldStr)
	sqlCount := fmt.Sprintf(sqlFmt, "count(*) cnt")
	o := orm.NewOrm()

	// todo 这里是这样处理的
	var params []orm.Params
	if _, err := o.Raw(sqlCount).Values(&params); err == nil {
		if len(params) > 0 {
			totalCount, _ = strconv.Atoi(params[0]["cnt"].(string))
		}
	}
	_, err = o.Raw(sql).QueryRows(&books)
	return
}


// 搜索书 通过 book 的 bookName 或 Descrition
func (m *Book) SearchBook(wd string, page, size int) (books []Book, cnt int, err error) {
	sqlFmt := "select %v from md_books where book_name like ? or description like ? order by star desc"
	sql := fmt.Sprintf(sqlFmt, "book_id")
	sqlCount := fmt.Sprintf(sqlFmt, "count(book_id) cnt")

	wd = "%" + wd + "%"
	o := orm.NewOrm()
	var count struct{ Cnt int }
	err = o.Raw(sqlCount, wd, wd).QueryRow(&count)  // 获取 count
	if count.Cnt > 0 {
		cnt = count.Cnt
		_, err = o.Raw(sql+" limit ? offset ?", wd, wd, size, (page-1)*size).QueryRows(&books)  // 获取分页的内容
	}
	return
}


// field 返回的字段
// 返回的 books 序列和传入的 ids 序列是需要一致 todo 这个好
func (m *Book) GetBooksByIds(ids []int, fields ...string) (books []Book, err error) {
	if len(ids) == 0 {
		return
	}
	var bs []Book
	var idArr []interface{}
	for _, i := range ids {
		idArr = append(idArr, i)  //  2, 1, 3,4
	}
	// in 查询
	rows , err := orm.NewOrm().QueryTable(TNBook()).Filter("book_id__in", idArr).All(&bs, fields...)
	if rows > 0 {
		//  返回的 books 序列和传入的 ids 序列是需要一致
		bookMap := make(map[interface{}]Book)
		for _, book := range bs {
			bookMap[book.BookId] = book
		}
		for _, i := range ids {  // 按照 ids 的顺序取出  用了 map 作为一个过渡
			if book, ok := bookMap[i]; ok {
				books = append(books, book)
			}
		}
	}
	return
}

// cols 是返回的字段
// cols 传入的是表中的字段
func (m *Book) Select(filed string, value interface{}, cols ...string) (book *Book, err error) {
	if len(cols) == 0 {
		err = orm.NewOrm().QueryTable(m.TableName()).Filter(filed, value).One(m)
	} else {
		err = orm.NewOrm().QueryTable(m.TableName()).Filter(filed, value).One(m, cols...)
	}
	return m, err
}
