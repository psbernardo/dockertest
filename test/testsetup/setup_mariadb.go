package testsetup

import "github.com/psbernardo/dockertest/test/database/mariadb"

func WithMariaDBTest() containerOptions {
	return func(s *TestService) error {
		resource, port, err := mariadb.SetupMariaDb(s.pool, *TestConfig.MariaDB)
		if err != nil {
			return err
		}

		host := buildURL(port)
		return s.AddContainer(host, resource)

	}
}
