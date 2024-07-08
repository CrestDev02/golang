package server

import (
	"database/sql"
	"my-app/config"
	"my-app/internal/auth"
	"my-app/internal/handlers"
	"my-app/internal/repository"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Config *config.Config
	DB     *sql.DB
}

func NewServer(config *config.Config, db *sql.DB) *Server {
	return &Server{Config: config, DB: db}
}

func (s *Server) Start() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(auth.Middleware)

	userRepo := &repository.UserRepositoryPostgres{DB: s.DB}
	userHandler := &handlers.UserHandler{Repo: userRepo}

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Post("/register", userHandler.Register)
	r.Post("/login", userHandler.Login)

	return http.ListenAndServe(":"+s.Config.Server.Port, r)
}
