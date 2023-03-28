package testingservice

import (
	"github.com/psbernardo/dockertest/infra/api/thirdpartyapi"
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

func (s *SuiteTest) NewThirdPartyAPIClient() *thirdpartyapi.Client {
	return thirdpartyapi.NewClient(&s.TestService.Config.TestAPIConfig)
}
