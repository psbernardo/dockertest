package testsetup

import (
	"fmt"

	"github.com/psbernardo/dockertest/test/api/thirdpartyapi"
)

func buildURL(port string) string {
	return fmt.Sprintf("%s%s", localHost, port)
}

func WithThirdPartyAPITest() containerOptions {
	return func(s *TestService) error {
		resource, port, err := thirdpartyapi.SetupThirdPartyAPI(s.pool, contextDIR)
		if err != nil {
			return err
		}

		host := buildURL(port)
		s.ThirdPartyAPIHost = host
		s.AddContainer(host, resource)
		return nil
	}
}
