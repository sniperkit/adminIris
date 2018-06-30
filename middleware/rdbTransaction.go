package middleware

import (
	"github.com/kataras/iris"
	"github.com/senseoki/adminIris/config"
	"github.com/senseoki/adminIris/util/mylog"
)

// RDBTransaction ...
func RDBTransaction(ctx iris.Context) {
	tx := config.CF.Storage.RDB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			CommonRecover(ctx, err)
			mylog.MyLogger.Info("[DBTransaction] Rollback() !!!")
		} else {
			tx.Commit()
			mylog.MyLogger.Info("[DBTransaction] Commit() !!!")
		}
	}()

	mylog.MyLogger.Info("[DBTransaction] Begin() !!!")

	ctx.Values().Set("rdb", tx)

	ctx.Next()
}
