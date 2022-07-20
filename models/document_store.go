package models


import (
	"github.com/astaxie/beego/orm"
)

//文档编辑
type DocumentStore struct {
	DocumentId int    `orm:"pk;auto;column(document_id)"`
	Markdown   string `orm:"type(text);"` //markdown内容
	Content    string `orm:"type(text);"` //html内容
}


func (m *DocumentStore) TableName() string {
	return TNDocumentStore()
}

func TNDocumentStore() string {
	return "md_document_store"
}

// 编辑文档内容
// todo 这一块没有明白
func (m *DocumentStore) SelectField(docId interface{}, filed string) string {
	var ds = DocumentStore{}
	if "markdown" != filed {
		filed = "content"
	}
	orm.NewOrm().QueryTable(TNDocumentStore()).Filter("document_id", docId).One(&ds, filed)
	if "content" == filed {
		return ds.Content
	}
	return ds.Markdown
}


// 删除记录
func (m *DocumentStore) Delete (docId ... interface{}) {
	if len(docId) > 0 {
		orm.NewOrm().QueryTable(TNDocumentStore()).Filter("document_id__in", docId...).Delete()
	}
}

// 插入或者更新
func (m *DocumentStore) InsertOrUpdate(fileds ... string) (err error) {
	o := orm.NewOrm()
	var one DocumentStore
	o.QueryTable(TNDocumentStore()).Filter("document_id", m.DocumentId).One(&one, "document_id")

	if one.DocumentId > 0 {
		_, err = o.Update(m)
	} else {
		_, err = o.Insert(m)
	}
	if err != nil {
		return err
	}
	return
}