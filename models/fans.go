package models

type Fans struct {
	Id       int //PK
	MemberId int
	FansId   int `orm:"index"` //粉丝id
}

type FansData struct {
	MemberId int
	Nickname string
	Avatar   string
	Account  string
}

func (m *Fans) TableName() string {
	return TNFans()
}

func TNFans() string {
	return "md_fans"
}