package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lai0xn/cr-dermasuim/api/http/routes"
	"github.com/lai0xn/cr-dermasuim/migrations"
	"github.com/lai0xn/cr-dermasuim/storage"
)

func main() {
	r := chi.NewRouter()
	storage.Conenct()

	migrations.Migrate()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	// media files fodler
	r.Handle(
		"/uploads/*",
		http.StripPrefix("/uploads/", http.FileServer(http.Dir("../../uploads"))),
	)
	routes.RegisterAllRoutes(r)
	log.Println("Server Started listening .....")
	http.ListenAndServe(":8080", r)
}
