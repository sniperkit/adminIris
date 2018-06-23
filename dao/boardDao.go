package dao

import (
	"github.com/senseoki/adminIris/models"
	"github.com/senseoki/adminIris/vo"
)

// NewBoardDao ...
func NewBoardDao() BoardDao {
	return &boardDao{}
}

// BoardDao ...
type BoardDao interface {
	SelectList(vo *vo.BoardVO) []*models.Board
}

type boardDao struct{}

func (d *boardDao) SelectList(vo *vo.BoardVO) []*models.Board {
	boards := []*models.Board{}
	vo.RDB.Find(&boards)
	return boards
}
