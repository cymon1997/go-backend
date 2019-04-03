package database

import (
	"github.com/cymon1997/go-backend/internal/config"
	"github.com/cymon1997/go-backend/internal/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBClient interface {
	Dial() error
	GetInstance() *sqlx.DB
}

type dbClient struct {
	db *sqlx.DB
}

func NewDBClient(cfg *config.DBConfig) DBClient {
	db, err := sqlx.Open(cfg.Driver, parseDSN(cfg))
	if err != nil {
		log.FatalDetail(log.TagDB, "error create db client", err)
	}
	return &dbClient{
		db: db,
	}
}

func (c *dbClient) Dial() error {
	err := c.db.Ping()
	if err != nil {
		log.ErrorDetail(log.TagDB, "error ping db", err)
		return err
	}
	return nil
}

func (c *dbClient) GetInstance() *sqlx.DB {
	return c.db
}
