package controller

import (
	"log"

	"github.com/senseoki/adminIris/vo"

	"github.com/jinzhu/gorm"
	"github.com/senseoki/adminIris/middleware"
	"github.com/senseoki/adminIris/service"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// BoardController ...
type BoardController struct {
	Ctx iris.Context

	BoardService service.BoardService
}

// BeforeActivation ...
func (c *BoardController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/board/list", "BoardList", middleware.RDBTransaction)
}

// BoardList ...
func (c *BoardController) BoardList() {
	boardVO := &vo.BoardVO{
		RDB: c.Ctx.Values().Get("rdb").(*gorm.DB),
	}
	boards := c.BoardService.SelectList(boardVO)
	log.Printf("BoardList : %+v\n", boards)
}
