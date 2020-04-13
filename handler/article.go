package handler

import (
	"context"
	"github.com/cymon1997/go-backend/entity"
	"net/http"

	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/internal/router"
	"github.com/cymon1997/go-backend/module/article/model"
)

type articleHandlerImpl struct {
	factory model.Factory
}

func NewArticleHandler(factory model.Factory) *articleHandlerImpl {
	return &articleHandlerImpl{
		factory: factory,
	}
}

func (h *articleHandlerImpl) Register(r router.Router) {
	r = r.SubRouter("/article")
	r.HandleJSON("", http.MethodGet, h.index)
	r.HandleJSON("/", http.MethodPost, h.insert)
	r.HandleJSON("/{id}", http.MethodGet, h.get)
	//r.HandleView("/view", http.MethodGet, h.view)
}

func (h *articleHandlerImpl) get(ctx context.Context, r *http.Request) (interface{}, error) {
	return h.factory.NewGetByIDModel(entity.GetArticleRequest{
		ID: GetURLParam(r, "id"),
	}).Do(ctx)
}

func (h *articleHandlerImpl) insert(ctx context.Context, r *http.Request) (interface{}, error) {
	//auth := r.Header.Get("Authorization")
	var data entity.Article
	err := ParseBody(r.Body, &data)
	if err != nil {
		log.ErrorDetail("Article", "error parse request body", err)
		return nil, err
	}
	return h.factory.NewInsertModel(data).Do(ctx)
}

func (h *articleHandlerImpl) index(ctx context.Context, r *http.Request) (interface{}, error) {
	return struct {
		Version string `json:"version"`
		Build   string `json:"build_version"`
	}{
		Version: "0.0.1",
		Build:   "alpha",
	}, nil
}

func (h *articleHandlerImpl) view(ctx context.Context, r *http.Request) (router.RenderRequest, error) {
	type invoice struct {
		Invoice string
		OrderID string
	}
	return router.RenderRequest{
		Template: "invoice.html",
		Data: invoice{
			Invoice: "INV/2018/123",
			OrderID: "123",
		},
	}, nil
}

func (h *articleHandlerImpl) health(ctx context.Context, r *http.Request) (interface{}, error) {
	return h.factory.NewHealthModel().Do(ctx)
}
