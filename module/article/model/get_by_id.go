package model

import (
	"context"

	"github.com/cymon1997/go-backend/entity"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/internal/uuid"
	"github.com/cymon1997/go-backend/internal/validator"
	"github.com/cymon1997/go-backend/module/article/repo"
)

const getArticleTag = "Article|Get"

type GetArticleModel struct {
	dbRepo repo.ArticleDBRepo
	req    entity.GetArticleRequest
}

func (m *GetArticleModel) Do(ctx context.Context) (entity.GetArticleResponse, error) {
	var response entity.GetArticleResponse
	err := m.Validate(ctx)
	if err != nil {
		log.ErrorDetail(getArticleTag, "error validation: %v", err)
		return response, err
	}
	response.Data, err = m.dbRepo.Get(ctx, m.req.ID)
	if err != nil {
		log.ErrorDetail(getArticleTag, "error get from db: %v", err)
		return response, err
	}
	return response, nil
}

func (m *GetArticleModel) Validate(ctx context.Context) error {
	v := validator.New()
	if !uuid.IsValid(m.req.ID) {
		v.Message("invalid uuid")
	}
	return v.Error()
}
