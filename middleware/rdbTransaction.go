package middleware

import (
	"fmt"

	"github.com/senseoki/adminIris/models"

	"github.com/kataras/iris"
	"github.com/senseoki/adminIris/config"
)

// RDBTransaction ...
func RDBTransaction(ctx iris.Context) {
	logger := ctx.Application().Logger()
	tx := config.CF.Storage.RDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			logger.Info("[DBTransaction] Rollback() !!!")
		}
		tx.Commit()
		logger.Info("[DBTransaction] Commit() !!!")
	}()

	logger.Info("[DBTransaction] Begin() !!!")

	ctx.Values().Set("DB", tx)

	var product models.Product

	tx.First(&product, 1)

	fmt.Printf("%+v\n", product)

	ctx.Next()
}
