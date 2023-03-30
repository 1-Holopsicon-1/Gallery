package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/ping"))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	router.Route("/auth", func(auth chi.Router) {
		auth.Post("/register", h.register)
		auth.Post("/login", h.login)
	})
	router.Route("/api", func(api chi.Router) {
		api.Route("/post", func(p chi.Router) {
			p.Get("/{id}", h.findPostById)
			p.Get("/all", h.getAll)
			p.With().Post("/publish", h.createPost)
			p.With().Delete("/{id}", h.deletePost)
			p.With().Patch("/edit/{id}", h.editTags)
		})
		api.Route("/tag", func(tag chi.Router) {
			tag.Get("/{tagNames}", h.findByTag)
			tag.Post("/name", h.TagsByLetters)
		})
	})
	return router
}
