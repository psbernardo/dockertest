package testingservice

import (
	"github.com/psbernardo/dockertest/infra/api/thirdpartyapi"
	"github.com/psbernardo/dockertest/infra/database/maria"
	"github.com/psbernardo/dockertest/infra/testingservice/testsetup"
)

type SuiteTest struct {
	TestService *testsetup.TestService
}

func (s *SuiteTest) SetupTestServices() error {
	testservice, err := testsetup.NewTestService(
		testsetup.WithThirdPartyAPITest(),
		testsetup.WithMariaDBTest(),
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

func (s *SuiteTest) NewMariaDBTestClient() *maria.PersonRepository {
	tx := s.TestService.Config.MariaDB.ConnectDB()
	return maria.NewPersonRepository(tx)
}
