package model

import (
	"context"

	"github.com/cymon1997/go-backend/entity"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/internal/validator"
	"github.com/cymon1997/go-backend/module/article/repo"
)

const insertArticleTag = "Article|Insert"

type InsertArticleModel struct {
	dbRepo repo.ArticleDBRepo
	req    entity.Article
}

func (m *InsertArticleModel) Do(ctx context.Context) (entity.InsertArticleResponse, error) {
	var response entity.InsertArticleResponse
	err := m.Validate(ctx)
	if err != nil {
		log.ErrorDetail(insertArticleTag, "error validation: %v", err)
		return response, err
	}
	response.ID, err = m.dbRepo.Insert(ctx, m.req)
	if err != nil {
		log.ErrorDetail(insertArticleTag, "error insert to db: %v", err)
		return response, err
	}
	return response, nil
}

func (m *InsertArticleModel) Validate(ctx context.Context) error {
	v := validator.New()
	if m.req.Title == "" {
		v.Missing("title")
	}
	if m.req.Description == "" {
		v.Missing("description")
	}
	if m.req.Content == "" {
		v.Missing("content")
	}
	return v.Error()
}
