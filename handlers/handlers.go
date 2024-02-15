// handlers/handlers.go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/seanyudhistira91/technical-test-YC-W22/entity"
	"github.com/seanyudhistira91/technical-test-YC-W22/service"
	"github.com/seanyudhistira91/technical-test-YC-W22/utils"
)

// Handler struct handles HTTP requests.
type Handler struct {
	service service.Service
}

// NewHandler creates a new instance of the Handler.
func NewHandler(svc service.Service) *Handler {
	return &Handler{service: svc}
}

type reqRegistration struct {
	Name        string
	PhoneNumber string
	Password    string
	Email       string
}

// GetResource handles GET requests for resource data.
func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var entityUser entity.Users
	var req reqRegistration

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// map from request to entities
	hashPassword, err := utils.HashAndSalt(req.Password)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	entityProfile := entity.Profiles{
		Name: req.Name,
	}
	entityUser = entity.Users{
		PhoneNumber:  req.PhoneNumber,
		HashPassword: hashPassword,
		Email:        req.Email,
		Profile:      &entityProfile,
	}

	err = h.service.CreateUser(ctx, entityUser)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the resource data as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "success created data"})
}

type reqLogin struct {
	PhoneNumber string
	Password    string
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var entityUser entity.Users
	var req reqLogin

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	entityUser = entity.Users{
		PhoneNumber:  req.PhoneNumber,
		HashPassword: req.Password,
	}

	err := h.service.Login(ctx, entityUser)
	if err != nil {
		http.Error(w, "Unauthorize", http.StatusUnauthorized)
		return
	}

	// Return the resource data as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "success login"})
}
