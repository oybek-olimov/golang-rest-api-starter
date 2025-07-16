package handler

import (
	"encoding/json"
	"net/http"

	"go-crud-api/internal/service"
	"go-crud-api/pkg/jwt"
)

type AuthHandler struct {
	UserService *service.UserService
	JWT         *jwt.Manager
}

func NewAuthHandler(userService *service.UserService, jwtManager *jwt.Manager) *AuthHandler {
	return &AuthHandler{
		UserService: userService,
		JWT:         jwtManager,
	}
}


type signupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signinRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register godoc
// @Summary Ro'yxatdan o'tish
// @Description Foydalanuvchini ro'yxatdan o'tkazadi
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   request  body  model.SignUpRequest  true  "Ro'yxatdan o'tish ma'lumotlari"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Router /auth/signup [post]
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req signupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.UserService.Register(r.Context(), req.Username, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req signinRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.UserService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := h.JWT.Generate(user.ID)
	if err != nil {
		http.Error(w, "could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"user":  user,
	})
}
