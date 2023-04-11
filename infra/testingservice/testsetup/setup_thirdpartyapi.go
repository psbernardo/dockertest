package testsetup

import (
	"github.com/psbernardo/dockertest/infra/testingservice/thirdpartyapi"
)

func WithThirdPartyAPITest() containerOptions {
	return func(s *TestService) error {
		resource, err := thirdpartyapi.SetupThirdPartyAPI(s.pool, contextDIR)
		if err != nil {
			return err
		}
		return s.addContainer(resource)
	}
}
