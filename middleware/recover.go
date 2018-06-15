package middleware

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/kataras/iris/context"
	"github.com/senseoki/adminIris/util/log"
)

// NewRecover ..
func NewRecover(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ctx.IsStopped() {
				return
			}

			var stacktrace string
			for i := 1; ; i++ {
				_, f, l, got := runtime.Caller(i)
				if !got {
					break

				}

				stacktrace += fmt.Sprintf("%s:%d\n", f, l)
			}

			// when stack finishes
			logMessage := fmt.Sprintf("\n\n============================== ERROR LOG START ================================")
			logMessage += fmt.Sprintf("\n\nRecovered from a route's Handler('%s')\n", ctx.HandlerName())
			logMessage += fmt.Sprintf("At Request: %s\n", getRequestLogs(ctx))
			logMessage += fmt.Sprintf("Trace: %s\n", err)
			logMessage += fmt.Sprintf("\n%s\n", stacktrace)
			logMessage += fmt.Sprintf("\n================================ ERROR LOG END ================================\n\n")
			//ctx.Application().Logger().Warn(logMessage)
			//ctx.Application().Logger().Error(logMessage)
			log.MyLogger.Error(logMessage)

			ctx.StatusCode(500)
			ctx.StopExecution()
		}
	}()
	ctx.Next()
}

func getRequestLogs(ctx context.Context) string {
	var status, ip, method, path string
	status = strconv.Itoa(ctx.GetStatusCode())
	path = ctx.Path()
	method = ctx.Method()
	ip = ctx.RemoteAddr()
	// the date should be logged by iris' Logger, so we skip them
	return fmt.Sprintf("%v %s %s %s", status, path, method, ip)
}
