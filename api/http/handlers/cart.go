package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lai0xn/cr-dermasuim/api/http/middlewares"
	"github.com/lai0xn/cr-dermasuim/app/cart"
	uuid "github.com/satori/go.uuid"
)

type CartController struct {
	service *cart.Service
}

func NewCart() *CartController {
	return &CartController{}
}

func (c *CartController) RegisterRoutes(r chi.Router) {
	r2 := chi.NewRouter()
	r2.Use(middlewares.IsAuthenticated)
	r2.Post("/add/{productID}", c.AddItemToCart)
	r2.Delete("/delete/{productID}", c.RemoveFromCart)
	r2.Patch("/decrease/{productID}", c.DecreaseQuantity)
	r2.Get("/me/", c.ViewCart)

	r.Mount("/cart", r2)
}

func (c CartController) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	productID := chi.URLParam(r, "productID")
	uID, err := uuid.FromString(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})

	}
	pID, err := uuid.FromString(productID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})

	}
	err = c.service.AddToCart(uID, pID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(H{
		"message": "product added successfully",
	})
}

func (c CartController) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	productID := chi.URLParam(r, "productID")
	uID, err := uuid.FromString(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})

	}
	pID, err := uuid.FromString(productID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})

	}
	err = c.service.DeleteCartItem(uID, pID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(H{
		"message": "product removed successfully",
	})
}

func (c CartController) DecreaseQuantity(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	productID := chi.URLParam(r, "productID")
	uID, err := uuid.FromString(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})

	}
	pID, err := uuid.FromString(productID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})

	}
	err = c.service.DecreaseItemQuantity(uID, pID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(H{
		"message": "quantity decreased successfully",
	})
}

func (c CartController) ViewCart(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	uID, err := uuid.FromString(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})

	}
	items, err := c.service.ViewCart(uID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(H{
		"items": items,
	})
}
