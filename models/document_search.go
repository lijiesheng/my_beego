package models

import "time"

// 文档搜索结果
type DocumentData struct {
	DocumentId   int       `json:"doc_id"`
	DocumentName string    `json:"doc_name"`
	Identify     string    `json:"identify"`
	Release      string    `json:"release"` // Release 发布后的Html格式内容.
	Vcnt         int       `json:"vcnt"`    //文档图书被浏览次数
	CreateTime   time.Time `json:"create_time"`
	BookId       int       `json:"book_id"`
	BookIdentify string    `json:"book_identify"`
	BookName     string    `json:"book_name"`
}
