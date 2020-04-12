package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/internal/render"
	"github.com/gorilla/mux"
)

type Router interface {
	SubRouter(path string) Router
	HandleJSON(path string, method string, f func(ctx context.Context, r *http.Request) (interface{}, error))
	HandleView(path string, method string, f func(ctx context.Context, r *http.Request) (RenderRequest, error))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type RenderRequest struct {
	Template string
	Data     interface{}
}

type routerImpl struct {
	engine *mux.Router
	render render.Client
}

func New(render render.Client) *routerImpl {
	router := mux.NewRouter()
	return &routerImpl{
		engine: router,
		render: render,
	}
}

func (r *routerImpl) SubRouter(path string) Router {
	return &routerImpl{
		engine: r.engine.PathPrefix(path).Subrouter(),
		render: r.render,
	}
}

func (r *routerImpl) HandleView(path string, method string, f func(ctx context.Context, r *http.Request) (RenderRequest, error)) {
	ctx := context.Background()
	handler := func(w http.ResponseWriter, req *http.Request) {
		request, _ := f(ctx, req)
		r.render.Render(w, request.Template, request.Data)
	}
	r.engine.HandleFunc(path, handler).Methods(method)
}

func (r *routerImpl) HandleJSON(path string, method string, f func(ctx context.Context, r *http.Request) (interface{}, error)) {
	ctx := context.Background()
	handler := func(w http.ResponseWriter, req *http.Request) {
		result, err := f(ctx, req)
		var response *Response
		if err != nil {
			response = r.buildResponse(http.StatusInternalServerError, "internal server", nil)
		} else {
			response = r.buildResponse(http.StatusOK, "success", result)
		}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.ErrorDetail(log.TagHandler, "error encode response", err)
		}
	}
	r.engine.HandleFunc(path, handler).Methods(method)
}

func (r *routerImpl) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.engine.ServeHTTP(w, req)
}

func (r *routerImpl) buildResponse(status int, message string, payload interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Payload: payload,
	}
}
