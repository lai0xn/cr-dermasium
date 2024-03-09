package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lai0xn/cr-dermasuim/app/products"
	uuid "github.com/satori/go.uuid"
)

type productsController struct {
	service *products.Service
}

func (c *productsController) RegisterRoutes(r chi.Router) {
	r.Get("/products/all", c.GetProduct)
	r.Get("/product/:id", c.GetProduct)
}

func (c *productsController) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.service.GetAllProducts()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})
		return

	}
	json.NewEncoder(w).Encode(H{
		"products": products,
	})
}

func (c *productsController) GetProduct(w http.ResponseWriter, r *http.Request) {
	idURL := chi.URLParam(r, "id")
	id, err := uuid.FromString(idURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err,
		})
		return
	}
	product, err := c.service.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err,
		})
		return

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(H{
		"product": product,
	})
}
