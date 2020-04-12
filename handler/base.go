package handler

import "github.com/cymon1997/go-backend/internal/router"

type BaseHandler interface {
	Register(router.Router)
}
