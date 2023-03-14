package main

import (
	"fmt"
	"productapi/config"
	"productapi/infrastruktur/database/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	mysql.SeedCategoryData(db)
	mysql.SeedProductData(db)
	e := echo.New()

	fmt.Println("Running program ....")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
