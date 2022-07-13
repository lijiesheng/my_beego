package models

import "github.com/astaxie/beego/orm"

func init()  {
	// 自动生成 model
	orm.RegisterModel(
		new (Category),   // 如果没有Member 自动注册【】
		new (Book),
		new (Document),
		new (Attachment),
		new (DocumentStore),
		new (BookCategory),
		new (Member),
		new (Collection),
		new (Relationship),
		new (Fans),
		new (Comments),
		new (Score),
	)
}
