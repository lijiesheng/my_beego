package models

type Relationship struct {
	RelationshipId int `orm:"pk;auto;" json:"relationship_id"`
	MemberId       int `json:"member_id"`
	BookId         int ` json:"book_id"`
	RoleId         int `json:"role_id"` // common.BookRole
}


func (m *Relationship) TableName() string {
	return TNRelationship()
}

func TNRelationship() string {
	return "md_relationship"
}