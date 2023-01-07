package main

import (
	"fmt"
	"skyshi/config"
	"skyshi/factory"
	"skyshi/migration"
	"skyshi/utils/database/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitMysqlDB(cfg)
	e := echo.New()
	migration.InitMigrate(db)
	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 3030)))

}
