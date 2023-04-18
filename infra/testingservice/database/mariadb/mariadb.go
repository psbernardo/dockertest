package mariadb

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/psbernardo/dockertest/config/database_maria"
)

func SetupMariaDb(pool *dockertest.Pool, config *database_maria.Config) (*dockertest.Resource, error) {
	uuidWithHyphen := uuid.New()

	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	constainerName := fmt.Sprintf("MariaTestDB_%s", uuid)

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
		// PortBindings: map[docker.Port][]docker.PortBinding{
		// 	docker.Port(tcpPort): {
		// 		{HostIP: config.Host, HostPort: tcpPort},
		// 	},
		// },
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.NeverRestart()
	})
	tcpPort := "3306/tcp"
	dkPOrt := resource.GetPort(tcpPort)
	Int, _ := strconv.Atoi(dkPOrt)
	config.Port = Int

	if err != nil {
		return nil, err
	}
	// add six seconds waiting time to up the DB
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
