package testsetup

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ory/dockertest/v3"
	"github.com/psbernardo/dockertest/config"
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

var (
	TestConfig *config.Config
)

func NewTestService(options ...containerOptions) (*TestService, error) {
	var s TestService
	TestConfig = config.Read()
	pool, err := NewDockerServer()
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

func NewDockerServer() (*dockertest.Pool, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("could not construct pool: %v", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not connect to Docker: %v", err)
	}

	return pool, nil
}
