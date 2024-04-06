package user

import (
	"encoding/json"
	"net/http"

	"github.com/devjunaeid/ecom-userSide/types"
	"github.com/devjunaeid/ecom-userSide/utils"
	"github.com/rs/zerolog/log"
)

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
	// Receieve and Process JSON Payload
	var payload types.RegisterPayload
	err := utils.ParseJSON(r, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// Check If user already exsist.
	// Vaidate the register actions. (Handle error)
}
