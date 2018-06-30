package controller

import (
	"github.com/senseoki/adminIris/util/mylog"
	"github.com/senseoki/adminIris/vo"

	"github.com/jinzhu/gorm"
	"github.com/senseoki/adminIris/middleware"
	"github.com/senseoki/adminIris/service"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// BoardController ...
type BoardController struct {
	Ctx          iris.Context
	BoardService service.BoardService
}

// BeforeActivation ...
func (c *BoardController) BeforeActivation(b mvc.BeforeActivation) {
	// GET: /adm/board/list
	b.Handle("GET", "/board/list", "List", middleware.RDBTransaction)
	// GET: /adm/board/register
	b.Handle("GET", "/board/register", "RegisterPage")
	// POST: /adm/board/register
	b.Handle("POST", "/board/register", "Register", middleware.RDBTransaction)

}

// List ...
func (c *BoardController) List() {
	board := &vo.Board{
		RDB: c.Ctx.Values().Get("rdb").(*gorm.DB),
	}
	boards := c.BoardService.SelectList(board)
	mylog.MyLogger.Infof("BoardList : %+v\n", boards)
}

// RegisterPage ...
func (c *BoardController) RegisterPage() mvc.Result {
	return mvc.View{
		Layout: "layout/default.layout.html",
		Name:   "layout/default.yield.html",
		Data: iris.Map{
			"header":  "header/default.header.html",
			"sidebar": "menu/default.sidebar.html",
			"content": "board/register.html",
			"footer":  "footer/default.footer.html",
			"js": []string{
				"pages/board/register.js",
			},
		},
	}
}

// Register ...
func (c *BoardController) Register() {
	board := new(vo.Board)
	board.RDB = c.Ctx.Values().Get("rdb").(*gorm.DB)

	err := c.Ctx.ReadJSON(board)
	if err != nil {
		panic(err)
	}
	//mylog.MyLogger.Infof("Register boardVO : %+v\n", boardVO)

	c.BoardService.Regist(board)

	c.Ctx.StatusCode(iris.StatusOK)
}
