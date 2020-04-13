package model

import (
	"github.com/cymon1997/go-backend/entity"
	"github.com/cymon1997/go-backend/internal/mq"
	"github.com/cymon1997/go-backend/internal/redis"
	"github.com/cymon1997/go-backend/module/article/repo"
)

type Factory interface {
	NewInsertModel(req entity.Article) *InsertArticleModel
	NewGetByIDModel(req entity.GetArticleRequest) *GetArticleModel
	NewHealthModel() *HealthModel
}

type articleFactory struct {
	dbRepo    repo.ArticleDBRepo
	redis     redis.Client
	publisher mq.Publisher
}

func NewArticleFactory(dbRepo repo.ArticleDBRepo, redis redis.Client, publisher mq.Publisher) Factory {
	return &articleFactory{
		dbRepo:    dbRepo,
		redis:     redis,
		publisher: publisher,
	}
}

func (f *articleFactory) NewInsertModel(req entity.Article) *InsertArticleModel {
	return &InsertArticleModel{
		dbRepo: f.dbRepo,
		req:    req,
	}
}

func (f *articleFactory) NewGetByIDModel(req entity.GetArticleRequest) *GetArticleModel {
	return &GetArticleModel{
		dbRepo: f.dbRepo,
		req:    req,
	}
}

func (f *articleFactory) NewHealthModel() *HealthModel {
	return &HealthModel{
		dbRepo:      f.dbRepo,
		redisClient: f.redis,
		publisher:   f.publisher,
	}
}
