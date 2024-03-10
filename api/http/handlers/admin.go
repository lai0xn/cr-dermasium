package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/lai0xn/cr-dermasuim/app/products"
	uuid "github.com/satori/go.uuid"
)

var uploadDir = "../../uploads/"

type AdminController struct {
	service products.Service
}

func NewAdmin() *AdminController {
	return &AdminController{}
}

func (c *AdminController) RegisterRoutes(r chi.Router) {
	r.Post("/admin/products/create", c.CreateProduct)
}

func (c AdminController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload products.ProductPayload

	// Parse the uploaded file
	image, fileHeader, err := r.FormFile("image")
	fmt.Println(r.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer image.Close()
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := uuid.NewV4().String()
	// Create a file to store the uploaded image
	f, err := os.Create(filepath.Join(uploadDir, id+fileHeader.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Write the uploaded image to the file
	if _, err := io.Copy(f, image); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Decode the JSON payload
	err = schema.NewDecoder().Decode(&payload, r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the picture field of the payload
	payload.Picture = "uploads/" + id + fileHeader.Filename

	// Create the product
	if err := c.service.CreateProduct(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "product created",
	})
}

func (c AdminController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	pid, err := uuid.FromString(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	err = c.service.DeleteProduct(pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&H{
		"message": "product deleted",
	})
}
