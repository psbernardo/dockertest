package loadtestdata

import (
	"github.com/psbernardo/dockertest/infra/database/maria/entities"
	"gorm.io/gorm"
)

func WithNewPerson() func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		entitiesPerson := entities.Person{
			Name:     "test ingest name",
			LastName: "test ingest lastname",
			Age:      25,
		}

		if err := tx.Model(entities.Person{}).Create(&entitiesPerson).Error; err != nil {
			return err
		}
		return nil
	}

}

func WithNewPerson2() func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		entitiesPerson := entities.Person{
			Name:     "test ingest name2",
			LastName: "test ingest lastname2",
			Age:      25,
		}

		if err := tx.Model(entities.Person{}).Create(&entitiesPerson).Error; err != nil {
			return err
		}
		return nil
	}

}
