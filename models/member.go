package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Member struct {
	MemberId int `orm:"pk;auto" json:"member_id"`
	// member_id   int auto_increment   primary key,

	Account string `orm:"size(30);unique" json:"account"`
	// account   varchar(30)  default ''     not null,

	Nickname string `orm:"size(30);unique" json:"nickname"`
	// nickname        varchar(30)  default ''     not null,

	Password string ` json:"-"`
	// password        varchar(255) default ''     not null,

	Description string `orm:"size(640)" json:"description"`
	//  description     varchar(640) default ''     not null,

	Email         string    `orm:"size(100);unique" json:"email"`
	// email           varchar(100) default ''     not null,

	Phone         string    `orm:"size(20);null;default(null)" json:"phone"`
	//phone           varchar(20)  default 'null' null,

	Avatar        string    `json:"avatar"`
	//  avatar          varchar(255) default ''     not null,

	Role          int       `orm:"default(1)" json:"role"`
	// role            int          default 1      not null,

	RoleName      string    `orm:"-" json:"role_name"`
	// 数据库中不存在

	Status        int       `orm:"default(0)" json:"status"`
	// status          int          default 0      not null,

	CreateTime    time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`
	// create_time     datetime                    not null,

	CreateAt      int       `json:"create_at"`
	// create_at       int          default 0      not null,

	LastLoginTime time.Time `orm:"type(datetime);null" json:"last_login_time"`
	// last_login_time datetime                    null,

	// constraint account
	//        unique (account),
	//    constraint email
	//        unique (email),
	//    constraint nickname
	//        unique (nickname)
}


func (m *Member) TableName() string {
	return TNMembers()
}

func TNMembers() string {
	return "md_members"
}

func (m *Member) Find(id int) (member *Member, err error) {
	sql := "select * from md_members where member_id = " + strconv.Itoa(id)
	newOrm := orm.NewOrm()
	err = newOrm.Raw(sql).QueryRow(&member)
	if err != nil && "<QuerySeter> no row found" != err.Error() {
		return nil, err
	}
	return
}
