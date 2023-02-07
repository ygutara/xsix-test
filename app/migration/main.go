package main

import (
	"fmt"

	"github.com/ygutara/xsis-test/app/config"
	"github.com/ygutara/xsis-test/model/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("error connect to DB")
	}

	migration.Migrate(db)
}
