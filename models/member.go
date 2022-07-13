package models


type Member struct {
	MemberId int `orm:"pk;auto" json:"member_id"`
	Account string `orm:"size(30);unique" json:"account"`
}
