package repo

import (
	"context"

	"github.com/cymon1997/go-backend/internal/database"
	"github.com/jmoiron/sqlx"
)

type DBRepo interface {
	GetDB() *sqlx.DB
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecNamed(ctx context.Context, query string, data interface{}) error
	QueryNamed(ctx context.Context, query string, data interface{}) (*sqlx.Rows, error)
}

type dbRepoImpl struct {
	db database.Client
}

func NewBaseDBRepo(db database.Client) DBRepo {
	return &dbRepoImpl{
		db: db,
	}
}

func (d *dbRepoImpl) GetDB() *sqlx.DB {
	return d.db.GetInstance()
}

//Get
//Retrieve single row
func (d *dbRepoImpl) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.db.GetInstance().GetContext(ctx, dest, query, args...)
}

//Select
//Retrieve multiple rows
func (d *dbRepoImpl) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.db.GetInstance().SelectContext(ctx, dest, query, args...)
}

//ExecNamed
//Query that not need to retrieve data
//Operation: INSERT
func (d *dbRepoImpl) ExecNamed(ctx context.Context, query string, data interface{}) error {
	_, err := d.db.GetInstance().NamedExecContext(ctx, query, data)
	return err
}

//QueryNamed
//Query that returning data
//Operation: INSERT returning data
func (d *dbRepoImpl) QueryNamed(ctx context.Context, query string, data interface{}) (*sqlx.Rows, error) {
	return d.db.GetInstance().NamedQueryContext(ctx, query, data)
}
