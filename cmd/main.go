package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

	"go-crud-api/config"
	"go-crud-api/internal/handler"
	"go-crud-api/internal/routes"
	"go-crud-api/internal/service"
	"go-crud-api/internal/store/postgres"
	"go-crud-api/pkg/hash"
	"go-crud-api/pkg/jwt"

	_ "go-crud-api/docs" // auto generated

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go CRUD API
// @version 1.0
// @description This is a CRUD API with User, Post, Comment functionality.
// @host localhost:8000
// @BasePath /

func main() {

	// 📦 Konfiguratsiyalarni yuklash
	cfg := config.LoadConfig()

	// 📦 PostgreSQL'ga ulanish
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("DB ulanish xatosi: %v", err)
	}
	defer db.Close()

	// 🔐 JWT Manager
	jwtManager := jwt.NewManager(cfg.JWTSecret, time.Minute*time.Duration(cfg.TokenDurationMinute))

	// 🔧 Store/service/handler init
	userStore := postgres.NewUserStore(db)
	userService := service.NewUserService(userStore, hash.NewHasher())
	authHandler := handler.NewAuthHandler(userService, jwtManager)

	postStore := postgres.NewPostStore(db)
	postService := service.NewPostService(postStore)
	postHandler := handler.NewPostHandler(postService, jwtManager)

	commentStore := postgres.NewCommentStore(db)
	commentService := service.NewCommentService(commentStore)
	commentHandler := handler.NewCommentHandler(commentService, jwtManager)

	// 🛣 Router
	r := chi.NewRouter()
	routes.RegisterRoutes(r, authHandler, postHandler, commentHandler, jwtManager)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// 🚀 Serverni ishga tushirish
	fmt.Printf("🚀 Server running at http://localhost:%s\n", cfg.AppPort)
	if err := http.ListenAndServe(":"+cfg.AppPort, r); err != nil {
		log.Fatalf("Server xatosi: %v", err)
	}
}
