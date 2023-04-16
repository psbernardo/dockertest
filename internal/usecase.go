package usecase

import (
	"github.com/psbernardo/dockertest/internal/model"
)

type ThirdPartyAPI interface {
	Get() (int, error)
	GetPerson(personId int) (*model.Person, error)
}

type Repository interface {
	CreatePerson(person *model.Person) (*model.Person, error)
}

type UseCase struct {
	testAPI          ThirdPartyAPI
	personRepository Repository
}

func NewUseCase(testAPI ThirdPartyAPI, repository Repository) *UseCase {
	return &UseCase{
		testAPI:          testAPI,
		personRepository: repository,
	}
}

func (r *UseCase) Get() (int, error) {
	return r.testAPI.Get()
}

func (r *UseCase) FetchAndCreate(personId int) (*model.Person, error) {
	// fetch the data to the third party api
	Person, err := r.testAPI.GetPerson(personId)
	if err != nil {
		return nil, err
	}
	// save the data to database
	return r.personRepository.CreatePerson(Person)
}
