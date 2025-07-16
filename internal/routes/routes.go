package routes

import (
	"fmt"
	"net/http"

	"go-crud-api/internal/handler"
	"go-crud-api/pkg/jwt"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, authHandler *handler.AuthHandler, postHandler *handler.PostHandler, commentHandler *handler.CommentHandler, jwtManager *jwt.Manager) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/signup", authHandler.SignUp)
		r.Post("/signin", authHandler.SignIn)

		// post routelar
		postHandler.RegisterRoutes(r)
		// comment routelar
		commentHandler.RegisterRoutes(r)

		// 🔐 Himoyalangan route
		r.Group(func(r chi.Router) {
			r.Use(jwtManager.Middleware)

			r.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
				userID, err := jwt.GetUserIDFromContext(r.Context())
				if err != nil {
					http.Error(w, "unauthorized", http.StatusUnauthorized)
					return
				}

				w.Write([]byte(fmt.Sprintf("User ID from token: %d", userID)))
			})
		})
	})
}
