package thirdpartyapi

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func SetupThirdPartyAPI(pool *dockertest.Pool, contextDir string) (*dockertest.Resource, error) {
	exposePort := "8000"

	if len(strings.TrimSpace(contextDir)) == 0 {
		contextDir = "../"
	}

	bOpts := &dockertest.BuildOptions{
		ContextDir: contextDir,
		Dockerfile: "./infra/testingservice/api/thirdpartyapi/implementation/Dockerfile",
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
		config.RestartPolicy = docker.NeverRestart()
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
		return nil, fmt.Errorf("could not connect to container: %s", err)
	}
	return resource, nil

}
