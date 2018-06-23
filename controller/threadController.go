package controller

import (
	"github.com/senseoki/adminIris/service"
	"github.com/senseoki/adminIris/vo/thread"

	"github.com/kataras/iris"
)

// ThreadController ...
type ThreadController struct {
	Ctx iris.Context

	ThreadService service.ThreadService
}

// GetThreads ...
func (c *ThreadController) GetThreads() {
	thread := &thread.ThreadVO{
	//DB: c.Ctx.Values().Get("DB").(*gorm.DB),
	}

	c.ThreadService.SelectList(thread)

}

// PostCreate ...
func (c *ThreadController) PostCreate() {

}
