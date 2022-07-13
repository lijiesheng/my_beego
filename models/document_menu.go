package models

//图书目录
type DocumentMenu struct {
	DocumentId   int             `json:"id"`
	DocumentName string          `json:"text"`
	BookIdentify string          `json:"-"`
	Identify     string          `json:"identify"`
	ParentId     interface{}     `json:"parent"`
	Version      int64           `json:"version"`
	State        *highlightState `json:"state,omitempty"` //如果字段为空，则json中不会有该字段
}

type highlightState struct {
	Selected bool `json:"selected"`
	Opened   bool `json:"opened"`
}