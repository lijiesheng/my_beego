package models

import "github.com/astaxie/beego/orm"

func init()  {
	// 自动生成 model
	orm.RegisterModel(
		new (Member),   // 如果没有Member 自动注册
	)
}
