package common

import "github.com/astaxie/beego"

// session
const SessionName = "__mbook_session__"

// 用户权限
const (
	// 超级管理员.
	MemberSuperRole = 0
	//普通管理员.
	MemberAdminRole = 1
	//普通用户.
	MemberGeneralRole = 2
)

// app_key
func AppKey() string {
	return beego.AppConfig.DefaultString("app_key", "godoc")
}

func Role(role int) string {
	if role == MemberSuperRole {
		return "超级管理员"
	} else if role == MemberAdminRole {
		return "管理员"
	} else if role == MemberGeneralRole {
		return "普通用户"
	} else {
		return ""
	}
}