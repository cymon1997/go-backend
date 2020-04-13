package database

import (
	"fmt"

	"github.com/cymon1997/go-backend/internal/config"
)

const (
	mysqlDriver    = "mysql"
	postgresDriver = "postgres"
)

//DSN FORMAT : mysql://root:root_pass@127.0.0.1:3306/test_db
//LIB FORMAT : root:root_pass@127.0.0.1:3306/test_db
func parseDSN(cfg config.DBConfig) string {
	switch cfg.Driver {
	case mysqlDriver:
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	case postgresDriver:
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	}
	return fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.SSLMode)
}
