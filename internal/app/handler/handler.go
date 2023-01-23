package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/"))
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
			tag.Get("/{tagName}", h.findByTag)
		})
	})
	return router
}
