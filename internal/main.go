package main

import (
	"fmt"

	"github.com/psbernardo/dockertest/internal/model"
)

func main() {
	fmt.Println("hi patrick")
}

type ThirdPartyAPI interface {
	Get() (int, error)
	GetPerson(personId int) (*model.Person, error)
}

type Repository interface {
	CreatePerson(person *model.Person) (*model.Person, error)
}

type TestRest struct {
	TestAPI          ThirdPartyAPI
	PersonRepository Repository
}

func NewTestRest(testAPI ThirdPartyAPI, repository Repository) *TestRest {
	return &TestRest{
		TestAPI:          testAPI,
		PersonRepository: repository,
	}
}

func (r *TestRest) Get() (int, error) {
	return r.TestAPI.Get()
}

func (r *TestRest) FetchAndCreate(personId int) (*model.Person, error) {
	// fetch the data to the third party api
	Person, err := r.TestAPI.GetPerson(personId)
	if err != nil {
		return nil, err
	}
	// save the data to database
	return r.PersonRepository.CreatePerson(Person)
}
