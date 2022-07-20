package utils

import "sync"

// 书籍发布锁 todo 这个很重要
type BookLock struct {
	Books map[int]bool
	Lock sync.RWMutex
}

//
var BookRelease = BookLock{
	Books: make(map[int]bool),  // 初始化一个 map
}

// 查询发布任务
// bool 的默认值是 false
func (this BookLock) Exist(bookId int) (exist bool) {
	this.Lock.RLock()
	defer this.Lock.RUnlock()
	_, exist = this.Books[bookId]
	return
}

// 设置
func (this BookLock) Set(bookId int) {
	this.Lock.RLock()
	defer this.Lock.RUnlock()
	this.Books[bookId] = true
}

// 删除
func (this BookLock) Delete(bookId int) {
	this.Lock.RLock()
	defer this.Lock.RUnlock()
	delete(this.Books, bookId)
}


