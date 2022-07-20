package models

import "github.com/astaxie/beego/orm"

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





// 关注或者取消关注 【通过 member_id 和 fan_id 找到一条记录就是关注】

// bool 的默认值是 false
// error 的默认值是
func (m *Fans) FollowOrCancel(mid, fansId int) (cancel bool, err error) {
	var fans Fans
	o := orm.NewOrm()
	qs := o.QueryTable(TNFans()).Filter("member_id", mid).Filter("fans_id", fansId)
	qs.One(&fans)
	if fans.Id > 0 {  // 【以前关注过】取消关注
		_, err = qs.Delete()  // 删除一条记录
		cancel = true
	} else {  // 关注
		fans.MemberId = mid
		fans.FansId = fansId
		_, err = o.Insert(&fans)
	}
	return
}