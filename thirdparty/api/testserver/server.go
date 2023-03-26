package testserver

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

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

func SetupThirdPartyAPI(pool *dockertest.Pool, contextDir string) (*dockertest.Resource, string, error) {
	exposePort := "8000"

	if len(strings.TrimSpace(contextDir)) == 0 {
		contextDir = "../"
	}

	bOpts := &dockertest.BuildOptions{
		ContextDir: contextDir,
		Dockerfile: "./thirdparty/api/test/Dockerfile",
	}

	rOpts := &dockertest.RunOptions{
		Name:         "test-third-party-api",
		ExposedPorts: []string{exposePort},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"8000/tcp": {
				{HostIP: "127.0.0.1", HostPort: "8000/tcp"},
			},
		},
	}

	hcOptions := func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	}

	resource, err := pool.BuildAndRunWithBuildOptions(bOpts, rOpts, hcOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	err = pool.Retry(func() error {
		HTTP_PORT := resource.GetPort("8000/tcp")
		_, err := net.Dial("tcp", net.JoinHostPort("localhost", HTTP_PORT))
		return err
	})

	if err != nil {
		return nil, exposePort, fmt.Errorf("could not connect to container: %s", err)
	}
	return resource, exposePort, nil

}
