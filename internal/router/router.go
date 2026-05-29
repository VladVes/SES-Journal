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
	u := handler.NewUserHandler()

	r.mux.HandleFunc("GET /", h.Index)
	r.mux.HandleFunc("GET /users", u.Index)

	return r.mux
}
