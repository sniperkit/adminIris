package service

import (
	"github.com/senseoki/adminIris/dao"
	"github.com/senseoki/adminIris/vo/thread"
)

// NewThreadService ...
func NewThreadService(threadDao dao.ThreadDao) ThreadService {
	return &threadService{
		dao: threadDao,
	}
}

// ThreadService ...
type ThreadService interface {
	SelectList(thread *thread.ThreadVO)
}

type threadService struct {
	dao dao.ThreadDao
}

func (s *threadService) SelectList(thread *thread.ThreadVO) {
	s.dao.SelectList(thread)
}
