package mariadb

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/psbernardo/dockertest/config/database_maria"

	_ "github.com/go-sql-driver/mysql"
)

func SetupMariaDb(pool *dockertest.Pool, config database_maria.Config) (*dockertest.Resource, error) {
	tcpPort := fmt.Sprintf("%d/tcp", config.Port)

	constainerName := "MariaTestDB"

	// finds a container with the given name and returns it if present
	if r, ok := pool.ContainerByName(constainerName); ok {
		return r, nil
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mariadb",
		Tag:        "latest",
		Name:       constainerName,
		Env: []string{
			"MARIADB_ROOT_PASSWORD=secret",
			fmt.Sprintf("MARIADB_USER=%s", config.Username),
			fmt.Sprintf("MARIADB_PASSWORD=%s", config.Password),
			fmt.Sprintf("MARIADB_DATABASE=%s", config.Database),
		},
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port(tcpPort): {
				{HostIP: config.Host, HostPort: tcpPort},
			},
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.NeverRestart()
	})

	if err != nil {
		return nil, err
	}
	// add six seconds waiting time to up the DB
	time.Sleep(time.Second * 6)
	if err := pool.Retry(func() error {
		var err error
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", config.Username, config.Password, config.Host, resource.GetPort(tcpPort), config.Database))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to mariadb database: %s", err)
		return nil, err
	}

	return resource, nil

}
