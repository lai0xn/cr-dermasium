package handlers

import (
	"crypto/hmac"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lai0xn/cr-dermasuim/api/http/middlewares"
	"github.com/lai0xn/cr-dermasuim/app/orders"
	"github.com/lai0xn/cr-dermasuim/app/utils"
	uuid "github.com/satori/go.uuid"
)

type OrdersController struct {
	service *orders.Service
}

func NewOrder() *OrdersController {
	return &OrdersController{}
}

func (c *OrdersController) RegistrRoutes(r chi.Router) {
	or := chi.NewRouter()
	r.Post("/webhook/", c.Webhook)

	or.Use(middlewares.IsAuthenticated)
	or.Post("/", c.CheckOut)
	r.Mount("/checkout/", or)
}

func (c *OrdersController) CheckOut(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	id, err := uuid.FromString(userID)

	order, err := c.service.CheckOut(id)

	response := utils.CreateCheckout(
		order.TotalPrice,
		"dzd",
		"https://"+r.Host+"/checkout/webhook/"+order.ID.String(),
	)
	print(string(response))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var checkout orders.CheckoutResponse
	json.Unmarshal(response, &checkout)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(checkout)
}

func (c *OrdersController) Webhook(w http.ResponseWriter, r *http.Request) {
	var checkout orders.WebhookPayload
	id := chi.URLParam(r, "orderID")
	signature := r.Header.Get("signature")

	// Getting the raw payload from the request body
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// If there is no signature, ignore the request
	if signature == "" {
		http.Error(w, "No signature provided", http.StatusBadRequest)
		return
	}

	// Calculate the signature
	computedSignature := utils.ComputeHMAC(payload)

	// If the calculated signature doesn't match the received signature, ignore the request
	if !hmac.Equal([]byte(computedSignature), []byte(signature)) {
		http.Error(w, "Invalid signature", http.StatusForbidden)
		return
	}

	orderID, _ := uuid.FromString(id)
	json.NewDecoder(r.Body).Decode(&checkout)
	c.service.HandleWebhook(checkout, orderID)
}
