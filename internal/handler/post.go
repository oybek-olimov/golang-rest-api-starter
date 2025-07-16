package handler

import (
	"encoding/json"
	"go-crud-api/internal/model"
	"go-crud-api/internal/service"
	"go-crud-api/pkg/jwt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	Service    *service.PostService
	JWTManager *jwt.Manager
}

func NewPostHandler(service *service.PostService, jwtManager *jwt.Manager) *PostHandler {
	return &PostHandler{
		Service:    service,
		JWTManager: jwtManager,
	}
}

func (h *PostHandler) RegisterRoutes(r chi.Router) {
	r.Route("/posts", func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Get("/{id}", h.GetByID)

		r.Group(func(r chi.Router) {
			r.Use(h.JWTManager.Middleware)
			r.Post("/", h.Create)
			r.Put("/{id}", h.Update)
			r.Delete("/{id}", h.Delete)
		})
	})
}

func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Service.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	post, err := h.Service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, err := jwt.GetUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var post model.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	post.UserID = userID
	if err := h.Service.Create(r.Context(), &post); err != nil {
		http.Error(w, "Failed to create", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID, _ := jwt.GetUserIDFromContext(r.Context())
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var post model.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	post.ID = id
	post.UserID = userID
	if err := h.Service.Update(r.Context(), &post); err != nil {
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, _ := jwt.GetUserIDFromContext(r.Context())
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Service.Delete(r.Context(), id, userID); err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
