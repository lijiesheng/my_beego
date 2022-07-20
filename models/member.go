package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"regexp"
	"strings"
	"time"
	"zs403_mbook_copy/common"
	"zs403_mbook_copy/utils"
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
	m.MemberId = id
	if err := orm.NewOrm().Read(m);err != nil {
		return m, err
	}
	m.RoleName = common.Role(m.Role)
	return m, err
}

func (m *Member) IsAdministrator () bool{
	if m == nil || m.MemberId <= 0 {
		return false
	}
	return m.Role == 0 || m.Role == 1
}


// cols 要更新的列 默认是全部
func (m *Member) Update(cols ...string) error {
	if m.Email == "" {
		return errors.New("邮箱不能为空")
	}
	if _,err := orm.NewOrm().Update(m, cols...);err != nil {
		return err
	}
	return nil
}


// 用户密码登录
func (m *Member) Login(account, password string) (member *Member, err error) {
	if err = orm.NewOrm().QueryTable(m.TableName()).Filter("account", account).
		Filter("status" , 0).One(member); err != nil {
			return member, errors.New("用户不存在")
	}
	ok, err := utils.PasswordVerify(member.Password, password)
	if ok && err == nil {
		m.RoleName = common.Role(m.Role)
		return member, nil
	}
	return member, errors.New("密码错误")
}


// 添加用户
func (m *Member) Add() error {
	if m.Email == "" {
		return errors.New("请填写邮箱")
	}
	if ok, err := regexp.MatchString(common.RegexpEmail, m.Email);
		!ok || err != nil || m.Email == "" {
		return errors.New("邮箱格式错误")
	}
	if l := strings.Count(m.Password, ""); l < 6 || l >= 20 {
		return errors.New("密码请输入6-20个字符")
	}

	cond := orm.NewCondition().Or("email", m.Email).Or("nickname", m.Nickname).Or("account", m.Account)
	var one Member
	o := orm.NewOrm()
	if o.QueryTable(m.TableName()).SetCond(cond).One(&one, "member_id", "nickname", "account", "email");one.MemberId > 0 {
		if one.Nickname == m.Nickname {
			return errors.New("昵称已存在")
		}
		if one.Email == m.Email {
			return errors.New("邮箱已存在")
		}
		if one.Account == m.Account {
			return errors.New("用户已存在")
		}
	}
	hash, err := utils.PasswordHash(m.Password)  // 密码用 hash 存起来

	if err != nil {
		return err
	}
	m.Password = hash
	_, err = o.Insert(m)

	if err != nil {
		return err
	}
	m.RoleName = common.Role(m.Role)
	return nil
}
