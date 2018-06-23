package dao

import "github.com/senseoki/adminIris/vo/thread"

// NewThreadDao ...
func NewThreadDao() ThreadDao {
	return &threadDao{}
}

// ThreadDao ...
type ThreadDao interface {
	SelectList(thread *thread.ThreadVO)
}

type threadDao struct{}

func (d *threadDao) SelectList(vo *thread.ThreadVO) {
	vo.DB.Find(&vo.Thread.Threads)
}
