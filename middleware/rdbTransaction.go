package middleware

import (
	"github.com/kataras/iris"
	"github.com/senseoki/adminIris/config"
	"github.com/senseoki/adminIris/util/log"
)

// RDBTransaction ...
func RDBTransaction(ctx iris.Context) {
	tx := config.CF.Storage.RDB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			CommonRecover(ctx, err)
			log.MyLogger.Info("[DBTransaction] Rollback() !!!")
		} else {
			tx.Commit()
			log.MyLogger.Info("[DBTransaction] Commit() !!!")
		}
	}()

	log.MyLogger.Info("[DBTransaction] Begin() !!!")

	ctx.Values().Set("rdb", tx)

	ctx.Next()
}
