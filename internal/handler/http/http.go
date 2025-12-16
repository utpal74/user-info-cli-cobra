package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/user-info-cli-tool/internal/controller/user"
)

// Handler defines a user http handler.
type Handler struct {
	ctrl *user.Controller
}

// New creates a user http handler.
func New(ctrl *user.Controller) *Handler{
	return &Handler{ctrl}
}

// GetUser handles GET /user requests.
func (h *Handler) GetUser(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	name = strings.TrimSpace(name) 
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := req.Context()
	u, err := h.ctrl.GetUser(ctx, name)
	if err != nil && errors.Is(err, user.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository got error: %v\n",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(u); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}

// CreateUser handles - POST /create-user requests.
func (h *Handler) CreateUser(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var input struct{
		Name string `json:"name"`
		MobileNo string `json:"mobile_no"`
	}
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil{
		log.Printf("rquest json decode error: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := req.Context()
	if err := h.ctrl.CreateUser(ctx, input.Name, input.MobileNo); err != nil{
		log.Printf("create user error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
