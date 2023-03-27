package testsetup

type SuiteTest struct {
	TestService *TestService
}

func (s *SuiteTest) SetupTestServices() error {
	testservice, err := NewTestService(
		WithThirdPartyAPITest(),
		WithMariaDBTest(),
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
