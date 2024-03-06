package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lai0xn/cr-dermasuim/app/auth"
	"github.com/lai0xn/cr-dermasuim/app/utils"
)

type authController struct {
	service *auth.Service
}

func NewAuth() *authController {
	return &authController{}
}

func (c *authController) RegisterRoutes(r chi.Router) {
	r.Post("/signin", c.Signin)
	r.Post("/token/refresh", c.RefreshJwt)
	r.Get("/verification/verify", c.Verify)
	r.Post("/verification/resend", c.SendOTP)
}

func (c *authController) Signin(w http.ResponseWriter, r *http.Request) {
	var payload auth.LoginPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})

		return
	}
	user, err := c.service.Signin(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})

		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(H{
		"message": "code sent successfully",
		"userID":  base64.RawStdEncoding.EncodeToString([]byte(user.ID.String())),
	})
}

func (c *authController) Signup(w http.ResponseWriter, r *http.Request) {
	var payload auth.SignUpPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	err = c.service.Signup(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(H{
		"message": "user created",
	})
}

func (c *authController) Verify(w http.ResponseWriter, r *http.Request) {
	id_base64 := r.URL.Query().Get("token")
	code := r.URL.Query().Get("code")
	id, err := base64.RawStdEncoding.DecodeString(id_base64)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})

		return
	}
	user, err := c.service.VerifyOTP(string(id), code)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(H{
			"error": err.Error(),
		})
		return
	}
	new_user := false
	if user.FirstName == "" {
		new_user = true
	}
	token, err := utils.GenerateToken(*user)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(H{
		"verified": true,
		"new_user": new_user,
		"token":    token,
	})
}

func (c *authController) RefreshJwt(w http.ResponseWriter, r *http.Request) {
	type tokenPayload struct {
		Token string `json:"token"`
	}
	var token tokenPayload
	json.NewDecoder(r.Body).Decode(&token)
	t, err := utils.RefreshToken(token.Token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(H{
		"token": t,
	})
}

func (c *authController) SendOTP(w http.ResponseWriter, r *http.Request) {
	type phone struct {
		PhoneNumber string `json:"phoneNumber"`
	}
	var number phone
	json.NewDecoder(r.Body).Decode(&number)
	err := c.service.SendOTP(number.PhoneNumber)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(H{
		"message": "verification code sent",
	})
}
