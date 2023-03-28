package maria

import (
	"github.com/psbernardo/dockertest/infra/database/maria/entities"
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) error {
	return DB.AutoMigrate(entities.Person{})
}
