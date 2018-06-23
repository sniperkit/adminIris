package main

import (
	stdCtx "context"

	"html/template"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/pprof"
	"github.com/kataras/iris/mvc"
	"github.com/senseoki/adminIris/config"
	"github.com/senseoki/adminIris/controller"
	"github.com/senseoki/adminIris/middleware"
	"github.com/senseoki/adminIris/service"
	"github.com/senseoki/adminIris/util"
)

func main() {
	app := iris.New()

	tmpl := iris.HTML("./view", ".html").Delims("[[", "]]")
	tmpl.AddFunc("CSS", func(tags []string) template.HTML {
		return util.TemplateFuncCSSAndJavascript(tags, "css")
	})
	tmpl.AddFunc("SCRIPT", func(tags []string) template.HTML {
		return util.TemplateFuncCSSAndJavascript(tags, "js")
	})

	app.RegisterView(tmpl)
	app.StaticWeb("/resources", "resources")

	appSpecialHandler(app)

	mvc.Configure(
		app.Party("/adm",
			middleware.NewRecover,
			middleware.RDBTransaction,
		), adm)

	mvc.Configure(
		app.Party("/api/v1",
			middleware.NewRecover,
		), api)

	ShutDown(app)
	app.Run(
		iris.Addr(":9000"),
		iris.WithoutVersionChecker,                    // disables updates
		iris.WithoutServerError(iris.ErrServerClosed), // skip err server closed when CTRL/CMD+C pressed:
		iris.WithOptimizations,                        // enables faster json serialization and more:
	)
}

// appSpecialHandler 함수는 mvc controllers 에서 정의되지 않는
// 경로에 대한 핸들러를 등록한다.
func appSpecialHandler(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/adm/dashboard")
	})

	// Any
	app.Any("/debug/pprof/{action:path}", pprof.New())
}

func adm(app *mvc.Application) {
	app.Register(
		service.NewDashBoardService(),
	)
	app.Handle(new(controller.DashBoardController))
}

func api(app *mvc.Application) {
	app.Register(
		service.NewHomeService(),
	)
	app.Handle(new(controller.HomeController))
}

// ShutDown 함수는 서버 강제 종료시 유입된 request를 처리하고
// 서버를 종료한다.
func ShutDown(app *iris.Application) {
	iris.RegisterOnInterrupt(func() {

		timeout := 5 * time.Second
		ctx, cancel := stdCtx.WithTimeout(stdCtx.Background(), timeout)

		defer func() {
			cancel()
			config.CF.Storage.RDB.Close()
		}()

		app.Shutdown(ctx) // close all hosts
	})
}
