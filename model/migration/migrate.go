package migration

import (
	"github.com/ygutara/xsis-test/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&model.Movie{})

	return nil
}
