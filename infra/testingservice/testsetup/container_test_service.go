package testsetup

import (
	"errors"
	"fmt"

	"github.com/ory/dockertest/v3"
	"github.com/psbernardo/dockertest/config"
)

var (
	contextDIR = "../"

	resouse_not_set_error = errors.New("resource not set")
	duplicate_container   = errors.New("duplicate container initialize")
)

type containerOptions func(s *TestService) error

type TestService struct {
	containers map[string]*dockertest.Resource
	pool       *dockertest.Pool
	Config     *config.Config
}

func NewTestService(options ...containerOptions) (*TestService, error) {
	s := TestService{
		Config:     config.Read(),
		containers: make(map[string]*dockertest.Resource),
	}

	pool, err := newDockerServer()
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

func (s *TestService) addContainer(resource *dockertest.Resource) error {
	if resource == nil {
		return resouse_not_set_error
	}

	// handle duplicate container
	if _, ok := s.containers[resource.Container.Name]; ok {
		return nil
	}

	s.containers[resource.Container.Name] = resource
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

func newDockerServer() (*dockertest.Pool, error) {
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
