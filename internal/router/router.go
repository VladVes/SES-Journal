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

func (r *Router) Register(uh *handler.UserHandler, rh *handler.RecordHandler) http.Handler {
	h := handler.NewHealthHandler()

	r.mux.HandleFunc("GET /", h.Index)
	r.mux.HandleFunc("GET /users", uh.HandleUsersList)
	r.mux.HandleFunc("POST /users", uh.HandleCreateUser)

	r.mux.HandleFunc("GET /records", rh.HandleRecordsList)
	r.mux.HandleFunc("GET /records/{id}", rh.HandleGetRecordByID)

	return r.mux
}
