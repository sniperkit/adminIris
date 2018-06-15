package controller

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/senseoki/adminIris/service"
	"github.com/senseoki/adminIris/vo/home"
)

// HomeController ...
type HomeController struct {
	Ctx iris.Context

	HomeService service.HomeService
}

// GetGreeting ...
// GET : /v1/greeting
func (c *HomeController) GetGreeting() mvc.Result {
	home := &home.Home{}
	fmt.Println(home)
	greeting := c.HomeService.Greeting()
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>" + greeting + "</h1>",
	}
}
