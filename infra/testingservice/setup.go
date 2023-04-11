package testingservice

import (
	"github.com/psbernardo/dockertest/infra/api/thirdpartyapi"
	"github.com/psbernardo/dockertest/infra/database/maria"
	"github.com/psbernardo/dockertest/infra/testingservice/testsetup"
	"gorm.io/gorm"
)

type DatabaseLoader func(tx *gorm.DB) error

type SuiteTest struct {
	TestService *testsetup.TestService
}

func (s *SuiteTest) SetupTestServices() error {
	testservice, err := testsetup.NewTestService(
		testsetup.WithThirdPartyAPITest(),
		testsetup.WithMariaDBTest(),
		testsetup.WithThirdPartyAPITest(),
	)
	if err != nil {
		return err
	}

	s.TestService = testservice
	return nil
}

func (s *SuiteTest) TearDownTestServices() error {
	return s.TestService.TearDownServices()
}

func (s *SuiteTest) NewThirdPartyAPITestClient() *thirdpartyapi.Client {
	return thirdpartyapi.NewClient(&s.TestService.Config.TestAPIConfig)
}

//	use variadic(function signature may have a type prefixed with ...)
//
// function which accept with zero or more of that  parameter
func (s *SuiteTest) NewMariaDBTestClient(dbLoader ...DatabaseLoader) (*maria.PersonRepository, error) {
	tx := s.TestService.Config.MariaDB.ConnectDB()
	if err := tx.Transaction(func(tx *gorm.DB) error {
		for _, loadData := range dbLoader {
			if err := loadData(tx); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return maria.NewPersonRepository(tx), nil
}
