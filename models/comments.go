package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

//评论表
type Score struct {
	Id         int
	BookId     int
	Uid        int
	Score      int //评分
	TimeCreate time.Time
}

//评论表
type Comments struct {
	Id         int
	Uid        int       `orm:"index"` //用户id
	BookId     int       `orm:"index"` //文档项目id
	Content    string    //评论内容
	TimeCreate time.Time //评论时间
}


func TNComments() string {
	return "md_comments"
}

func (m *Comments) TableName() string {
	return TNComments()
}

func TNScore() string {
	return "md_score"
}

// 这个表示唯一
func (m *Score) TableName() string {
	return TNScore()
}



// 多字段唯一键
// todo 不知道啥意思
func (m *Score) TableUnique() [][]string {
	return [][]string{
		[]string{"Uid", "BookId"},
	}
}

// 评分内容
type BookScoresResult struct {
	Avatar   string  	`json:"avatar"`
	Nickname  string     `json:"nickname"`
	Score     string     `json:"score"`
	TimeCreate time.Time  `json:"time_create"`
}


// 获取评分内容
func (m *Score) BookScores(p, listRows, bookId int) (scores []BookScoresResult, err error) {
	sql := `select s.score,s.time_create,m.avatar,m.nickname from ` + TNScore() + ` s left join ` + TNMembers() + ` m on m.member_id=s.uid where s.book_id=? order by s.id desc limit %v offset %v`
	sql = fmt.Sprintf(sql, listRows, (p-1)*listRows)
	_, err = orm.NewOrm().Raw(sql, bookId).QueryRows(&scores)
	return
}


//查询用户对文档的评分
func (m *Score) BookScoreByUid(uid, bookId interface{}) int {
	var score Score
	orm.NewOrm().QueryTable(TNScore()).Filter("uid", uid).Filter("book_id", bookId).One(&score, "score")
	return score.Score
}

func (m *Score) AddScore(uid, bookId, score int) (err error) {
	//查询评分是否已存在
	o := orm.NewOrm()
	var scoreObj = Score{Uid: uid, BookId: bookId}
	o.Read(&scoreObj, "uid", "book_id")
	if scoreObj.Id > 0 { //评分已存在
		err = errors.New("您已给当前文档打过分了")
		return
	}

	//评分不存在，添加评分记录
	score = score * 10
	scoreObj.Score = score
	scoreObj.TimeCreate = time.Now()
	o.Insert(&scoreObj)
	if scoreObj.Id > 0 { //评分添加成功，更行当前书籍项目的评分
		//评分人数+1
		var book = Book{BookId: bookId}
		o.Read(&book, "book_id")
		if book.CntScore == 0 {
			book.CntScore = 1
			book.Score = 0
		} else {
			book.CntScore = book.CntScore + 1
		}
		book.Score = (book.Score*(book.CntScore-1) + score) / book.CntScore
		_, err = o.Update(&book, "cnt_score", "score")
		if err != nil {
			beego.Error(err.Error())
			err = errors.New("评分失败，内部错误")
		}
	}
	return
}
