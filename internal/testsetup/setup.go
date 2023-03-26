package testsetup

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ory/dockertest/v3"
	"github.com/psbernardo/dockertest/thirdparty/api/testserver"
	"github.com/psbernardo/dockertest/thirdparty/database/mariadb/dbserver"
)

var (
	contextDIR = "../"
	localHost  = "http://localhost:"
)

type containerOptions func(s *TestService) error

type TestService struct {
	containersMap     map[string]*dockertest.Resource
	pool              *dockertest.Pool
	ThirdPartyAPIHost string
}

func NewTestService(options ...containerOptions) (*TestService, error) {
	var s TestService

	pool, err := testserver.NewDockerServer()
	if err != nil {
		return nil, err
	}
	s.pool = pool

	for _, opts := range options {
		if err := opts(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil

}

func (s *TestService) AddContainer(host string, resource *dockertest.Resource) error {
	if s.containersMap == nil {
		s.containersMap = make(map[string]*dockertest.Resource)
	}

	if len(strings.TrimSpace(host)) == 0 {
		return errors.New("host not set")
	}

	if resource == nil {
		return errors.New("resource not set")
	}

	s.containersMap[host] = resource
	return nil
}

func (s *TestService) TearDownServices() error {
	for _, r := range s.containersMap {
		if err := s.pool.Purge(r); err != nil {
			return err
		}
	}
	return nil
}

func buildURL(port string) string {
	return fmt.Sprintf("%s%s", localHost, port)
}

func WithThirdPartyAPITest() containerOptions {
	return func(s *TestService) error {
		resource, port, err := testserver.SetupThirdPartyAPI(s.pool, contextDIR)
		if err != nil {
			return err
		}

		host := buildURL(port)
		s.ThirdPartyAPIHost = host
		s.AddContainer(host, resource)
		return nil
	}
}

func WithMariaDBTest() containerOptions {
	return func(s *TestService) error {
		resource, port, err := dbserver.SetupMariaDb(s.pool)
		if err != nil {
			return err
		}

		host := buildURL(port)
		return s.AddContainer(host, resource)

	}
}
