package models

import "time"

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

func (m *Score) TableName() string {
	return TNScore()
}



