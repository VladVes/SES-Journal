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

func (r *Router) Register(uh *handler.UserHandler) http.Handler {
	h := handler.NewHealthHandler()

	r.mux.HandleFunc("GET /", h.Index)
	r.mux.HandleFunc("GET /users", uh.HandleUsersList)

	return r.mux
}
