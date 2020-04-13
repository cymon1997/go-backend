package repo

import (
	"context"
	"database/sql"
	"github.com/cymon1997/go-backend/entity"
	"github.com/cymon1997/go-backend/internal/base/repo"
	"github.com/cymon1997/go-backend/internal/errors"
)

type ArticleDBRepo interface {
	Insert(ctx context.Context, data entity.Article) (string, error)
	Get(ctx context.Context, id string) (entity.Article, error)
	Ping(ctx context.Context) error
}

type articleDBRepoImpl struct {
	db repo.DBRepo
}

func NewArticleDBRepo(db repo.DBRepo) ArticleDBRepo {
	return &articleDBRepoImpl{
		db: db,
	}
}

func (r *articleDBRepoImpl) Insert(ctx context.Context, data entity.Article) (id string, err error) {
	rows, err := r.db.QueryNamed(ctx, insertArticleQuery, data)
	if err != nil {
		return "", errors.New(errors.InternalServer).WithMessage(err.Error())
	}
	if rows.Next() {
		_ = rows.Scan(&id)
	}
	return
}

func (r *articleDBRepoImpl) Get(ctx context.Context, id string) (entity.Article, error) {
	var result entity.Article
	err := r.db.Get(ctx, &result, getArticleQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, errors.New(errors.NoDataFound)
		}
		return result, err
	}
	return result, nil
}

func (r *articleDBRepoImpl) Ping(ctx context.Context) error {
	return r.db.GetDB().Ping()
}
