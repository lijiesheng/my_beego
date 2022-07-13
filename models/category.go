package models


type Category struct {
	Id     int
	Pid    int    //分类id
	Title  string `orm:"size(30);unique"` // 唯一
	Intro  string //介绍
	Icon   string // 图标路径
	Cnt    int  //统计分类下图书
	Sort   int  //排序
	Status bool //状态，true 显示
}

/**
	表名
 */
func TNCategory() string{
	return "md_category"
}

// orm 回调
func (m *Category) TableName() string {
	return TNCategory()
}


// 获取所有分类
// -1 获取全部分类
func (m *Category) GetCates(pid , status int) (caets []Category, err error) {

	return nil, nil
}
