package router

import (
	"net/http"

	"github.com/VladVes/SES-Journal/internal/handler"
)

type Router struct {
	mux *http.ServeMux
}

func New() *Router {
	return &Router{mux: http.NewServeMux()}
}

func (r *Router) Register() http.Handler {
	h := handler.NewHealthHandler()

	r.mux.HandleFunc("GET /", h.Index)

	return r.mux
}
