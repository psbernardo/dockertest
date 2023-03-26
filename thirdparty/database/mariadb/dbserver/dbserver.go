package dbserver

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SetupMariaDb(pool *dockertest.Pool) (*dockertest.Resource, string, error) {
	exposePort := "3306"
	// pulls an image, creates a container based on it and runs it
	// resource, err := pool.Run("mariadb", "latest", []string{"MARIADB_ROOT_PASSWORD=secret", "MARIADB_USER=psbernardo", "MARIADB_PASSWORD=trustno1"})

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mariadb",
		Tag:        "latest",
		Env: []string{
			"MARIADB_ROOT_PASSWORD=secret",
			"MARIADB_USER=psbernardo",
			"MARIADB_PASSWORD=trustno1",
			"MARIADB_DATABASE=mydatabase",
		},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3306/tcp": {
				{HostIP: "127.0.0.1", HostPort: "3306/tcp"},
			},
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	if err != nil {
		return nil, "", err
	}

	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mydatabase?parseTime=true", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to mariadb database: %s", err)
		return nil, "", err
	}

	return resource, exposePort, nil

}
