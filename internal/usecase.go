package usecase

import (
	"context"

	"github.com/psbernardo/dockertest/internal/model"
	"github.com/psbernardo/dockertest/utl/tracing"
)

type ThirdPartyAPI interface {
	Get(ctx context.Context) (int, error)
	GetPerson(ctx context.Context, personId int) (*model.Person, error)
}

type Repository interface {
	CreatePerson(ctx context.Context, person *model.Person) (*model.Person, error)
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

func (r *UseCase) Get(ctx context.Context) (int, error) {
	return r.testAPI.Get(ctx)
}

func (r *UseCase) FetchAndCreate(ctx context.Context, personId int) (*model.Person, error) {
	tracing.NewTraceUseCase().StartSpanWithAttributes(ctx, map[string]interface{}{
		"personId": personId,
	})
	// fetch the data to the third party api
	Person, err := r.testAPI.GetPerson(ctx, personId)
	if err != nil {
		return nil, err
	}
	// save the data to database
	return r.personRepository.CreatePerson(ctx, Person)
}
