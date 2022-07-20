package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

//拼接返回到接口的图书信息  todo 没有对应数据库的表
type BookData struct {
	BookId         int       `json:"book_id"`
	BookName       string    `json:"book_name"`
	Identify       string    `json:"identify"`
	OrderIndex     int       `json:"order_index"`
	Description    string    `json:"description"`
	PrivatelyOwned int       `json:"privately_owned"`
	PrivateToken   string    `json:"private_token"`
	DocCount       int       `json:"doc_count"`
	CommentCount   int       `json:"comment_count"`
	CreateTime     time.Time `json:"create_time"`
	CreateName     string    `json:"create_name"`
	ModifyTime     time.Time `json:"modify_time"`
	Cover          string    `json:"cover"`
	MemberId       int       `json:"member_id"`
	Username       int       `json:"user_name"`
	Editor         string    `json:"editor"`
	RelationshipId int       `json:"relationship_id"`
	RoleId         int       `json:"role_id"`
	RoleName       string    `json:"role_name"`
	Status         int
	Vcnt           int    `json:"vcnt"`
	Collection     int    `json:"star"`
	Score          int    `json:"score"`
	CntComment     int    `json:"cnt_comment"`
	CntScore       int    `json:"cnt_score"`
	ScoreFloat     string `json:"score_float"`
	LastModifyText string `json:"last_modify_text"`
	Author         string `json:"author"`
	AuthorURL      string `json:"author_url"`
}

// 通过 memberId 和 identify [书的唯一标识]
func (m *BookData) SelectByIdentify(identify string, memberId int) (result *BookData, err error) {
	if identify == "" || memberId <= 0 {
		return result, errors.New("Invalid parameter")
	}
	book := (&Book{})
	o := orm.NewOrm()
	err = o.QueryTable(TNBook()).Filter("identify", identify).One(book)
	if err != nil {
		return
	}
	// 查看权限
	relationship := (&Relationship{})
	err = o.QueryTable(TNRelationship()).Filter("book_id", book.BookId).Filter("role_id", 0).One(relationship)
	if err != nil {
		return result, errors.New("Permission denied")
	}
	member, err := (&Member{}).Find(relationship.MemberId)
	if err != nil {
		return result, err
	}
	err = o.QueryTable(TNRelationship()).Filter("book_id", book.BookId).Filter("member_id", memberId).One(relationship)
	if err != nil {
		return
	}


}
