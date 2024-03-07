package routes

import (
	"github.com/go-chi/chi"
	"github.com/lai0xn/cr-dermasuim/api/http/handlers"
)

func RegisterAllRoutes(r chi.Router) {
	auth := handlers.NewAuth()
	profile := handlers.NewProfile()
	auth.RegisterRoutes(r)
	profile.RegisterRoutes(r)
}
