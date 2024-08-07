package mocking

import (
	"context"
	"errors"

	"github.com/psbernardo/dockertest/internal/model"
)

type MockThirdPartyAPI struct {
	GetFn       func(ctx context.Context) (int, error)
	GetPersonFn func(ctx context.Context, personId int) (*model.Person, error)
}

func (m *MockThirdPartyAPI) Get(ctx context.Context) (int, error) {
	if m.GetFn == nil {
		return 0, errors.New("GetFn is not set")
	}

	return m.GetFn(ctx)
}

func (m *MockThirdPartyAPI) GetPerson(ctx context.Context, personId int) (*model.Person, error) {
	if m.GetPersonFn == nil {
		return nil, errors.New("GetPersonFn is not set")
	}

	return m.GetPersonFn(ctx, personId)
}

type MockRepository struct {
	CreatePersonFn func(ctx context.Context, person *model.Person) (*model.Person, error)
}

func (m *MockRepository) CreatePerson(ctx context.Context, person *model.Person) (*model.Person, error) {
	if m.CreatePersonFn == nil {
		return nil, errors.New("GetPersonFn is not set")
	}

	return m.CreatePersonFn(ctx, person)
}
