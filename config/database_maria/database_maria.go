package database_maria

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config ...
type Config struct {
	Host     string `env:"HOST,default=localhost"`
	Port     int    `env:"PORT,default=3306"`
	Username string `env:"USER,default=psbernardo"`
	Password string `env:"PASSWORD,default=trustno1"`
	Database string `env:"DATABASE,default=mydatabase"`
}

// BuildDSN func
func buildDSN(config *Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}
func ConnectDB(dbConf *Config) *gorm.DB {
	dsn := buildDSN(dbConf)
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		panic(`fatal error: cannot connect to database`)
	}

	return dbConn
}
