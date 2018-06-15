package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/senseoki/adminIris/service"
)

// DashBoardController ...
type DashBoardController struct {
	Ctx iris.Context

	DashBoardService service.DashBoardService
}

// GetDashboard ...
// GET : /adm/dashboard
func (c *DashBoardController) GetDashboard() mvc.Result {

	return mvc.View{
		Layout: "layout/default.layout.html",
		Name:   "layout/default.yield.html",
		Data: iris.Map{
			"header":  "header/default.header.html",
			"sidebar": "menu/default.sidebar.html",
			"content": "dashboard.html",
			"footer":  "footer/default.footer.html",
		},
	}
}
