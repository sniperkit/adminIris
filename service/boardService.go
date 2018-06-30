package service

import (
	"github.com/senseoki/adminIris/dao"
	"github.com/senseoki/adminIris/models"
	"github.com/senseoki/adminIris/vo"
)

// NewBoardService ...
func NewBoardService(dao dao.BoardDao) BoardService {
	return &boardService{
		dao: dao,
	}
}

// BoardService ...
type BoardService interface {
	SelectList(vo *vo.Board) []*models.Board
	Regist(vo *vo.Board)
}

type boardService struct {
	dao dao.BoardDao
}

func (s *boardService) SelectList(vo *vo.Board) []*models.Board {
	return s.dao.SelectList(vo)
}

func (s *boardService) Regist(vo *vo.Board) {
	s.dao.Regist(vo)
}
