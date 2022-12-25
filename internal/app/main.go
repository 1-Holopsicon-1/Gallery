package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/"))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})

	fmt.Println(http.ListenAndServe(":4000", r))
}
