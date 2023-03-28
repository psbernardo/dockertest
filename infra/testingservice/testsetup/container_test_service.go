package testsetup

import (
	"errors"
	"fmt"

	"github.com/ory/dockertest/v3"
	"github.com/psbernardo/dockertest/config"
)

var (
	contextDIR = "../"
)

type containerOptions func(s *TestService) error

type TestService struct {
	containers []*dockertest.Resource
	pool       *dockertest.Pool
	Config     *config.Config
}

func NewTestService(options ...containerOptions) (*TestService, error) {
	s := TestService{
		Config: config.Read(),
	}

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

func (s *TestService) AddContainer(resource *dockertest.Resource) error {
	if resource == nil {
		return errors.New("resource not set")
	}

	s.containers = append(s.containers, resource)
	return nil
}

func (s *TestService) TearDownServices() error {
	for _, r := range s.containers {
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
