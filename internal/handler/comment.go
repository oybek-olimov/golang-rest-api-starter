package handler

import (
	"encoding/json"
	"go-crud-api/internal/model"
	"go-crud-api/pkg/jwt"
	"go-crud-api/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CommentHandler struct {
	Service    *service.CommentService
	JWTManager *jwt.Manager
}

func NewCommentHandler(s *service.CommentService, jwtManager *jwt.Manager) *CommentHandler {
	return &CommentHandler{Service: s, JWTManager: jwtManager}
}

func (h *CommentHandler) RegisterRoutes(r chi.Router) {
	r.Route("/comments", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(h.JWTManager.Middleware)
			r.Post("/", h.Create)
			r.Delete("/{id}", h.Delete)
		})
		r.Get("/post/{postID}", h.GetByPostID)
	})
}

func (h *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, _ := jwt.GetUserIDFromContext(r.Context())
	var c model.Comment
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	c.UserID = userID
	if err := h.Service.Create(r.Context(), &c); err != nil {
		http.Error(w, "failed to create", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(c)
}

func (h *CommentHandler) GetByPostID(w http.ResponseWriter, r *http.Request) {
	postID, _ := strconv.Atoi(chi.URLParam(r, "postID"))
	comments, err := h.Service.GetByPostID(r.Context(), postID)
	if err != nil {
		http.Error(w, "failed to get comments", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comments)
}

func (h *CommentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, _ := jwt.GetUserIDFromContext(r.Context())
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Service.Delete(r.Context(), id, userID); err != nil {
		http.Error(w, "failed to delete", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
