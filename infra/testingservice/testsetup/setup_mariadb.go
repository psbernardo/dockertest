package testsetup

import (
	"github.com/psbernardo/dockertest/infra/database/maria"
	"github.com/psbernardo/dockertest/infra/testingservice/database/mariadb"
)

func WithMariaDBTest() containerOptions {
	return func(s *TestService) error {
		resource, err := mariadb.SetupMariaDb(s.pool, *s.Config.MariaDB)
		if err != nil {
			return err
		}

		// Migration
		DB := s.Config.MariaDB.ConnectDB()
		if err := maria.Migrate(DB); err != nil {
			return err
		}
		db, err := DB.DB()
		if err != nil {
			return err
		}
		db.Close()

		return s.addContainer(resource)
	}
}
