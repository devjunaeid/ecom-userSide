package user

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /login", h.loginUser)
	router.HandleFunc("POST /register", h.registerUser)
}

func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {

}
