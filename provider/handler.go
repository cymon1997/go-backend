package provider

import (
	"sync"

	"github.com/cymon1997/go-backend/handler"
	"github.com/cymon1997/go-backend/internal/router"
)

var (
	articleHandler     handler.BaseHandler
	syncArticleHandler sync.Once
)

func GetHandlers() router.Router {
	GetArticleHandler().Register(GetRouter())
	return GetRouter()
}

func GetArticleHandler() handler.BaseHandler {
	syncArticleHandler.Do(func() {
		articleHandler = handler.NewArticleHandler(GetArticleFactory())
	})
	return articleHandler
}
