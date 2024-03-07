package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lai0xn/cr-dermasuim/api/http/middlewares"
	"github.com/lai0xn/cr-dermasuim/app/users"
	uuid "github.com/satori/go.uuid"
)

type profileController struct {
	service *users.Service
}

func NewProfile() *profileController {
	return &profileController{}
}

func (c *profileController) RegisterRoutes(r chi.Router) {
	subRouter := chi.NewRouter()
	subRouter.Use(middlewares.IsAuthenticated)
	subRouter.Post("/me/set", c.SetUserInfo)
	r.Mount("/users/profile", subRouter)
}

func (c *profileController) SetUserInfo(w http.ResponseWriter, r *http.Request) {
	var payload users.ProfilePayload
	userID := r.Context().Value("userID").(string)
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, err := uuid.FromString(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.service.SetUserProfile(id, payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(H{
		"message": "user profile set successfully",
	})
}
