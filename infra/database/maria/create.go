package maria

import (
	"github.com/psbernardo/dockertest/infra/database/maria/entities"
	"github.com/psbernardo/dockertest/internal/model"
	"gorm.io/gorm"
)

type PersonRepository struct {
	tx *gorm.DB
}

func NewPersonRepository(tx *gorm.DB) *PersonRepository {
	return &PersonRepository{
		tx: tx,
	}
}

func (p *PersonRepository) CreatePerson(person *model.Person) (*model.Person, error) {
	entitiesPerson := entities.Person{
		ID:       person.ID,
		Name:     person.Name,
		LastName: person.LastName,
		Age:      person.Age,
	}

	if err := p.tx.Model(entities.Person{}).Create(&entitiesPerson).Error; err != nil {
		return nil, err
	}

	if err := p.tx.Model(entities.Person{}).Where("id = ?", entitiesPerson.ID).Find(&entitiesPerson).Error; err != nil {
		return nil, err
	}

	return &model.Person{
		ID:       entitiesPerson.ID,
		Name:     entitiesPerson.Name,
		LastName: entitiesPerson.LastName,
		Age:      entitiesPerson.Age,
	}, nil

}
