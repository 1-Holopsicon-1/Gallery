package handler

import (
	"Gallery/internal/app/services"
	"encoding/json"
	"log"
	"net/http"
)

var tagService services.TagService

func (h *Handler) findByTag(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(tagService.FindPostByNameTag(h.DB, r))
	if err != nil {
		log.Fatalln(err)
	}
}
