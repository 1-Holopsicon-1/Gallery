package services

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type TagService struct {
}

func (ts *TagService) FindPostByNameTag(db *gorm.DB, r *http.Request) interface{} {
	var data []map[string]interface{}
	tagStr := chi.URLParam(r, "tagNames")
	tagNames := strings.Split(tagStr, "+")
	db.Raw("select "+
		"posts.id   as id,"+
		"posts.url  as url,"+
		"tags.id    as tag_id,"+
		"tags.name  as tag_name,"+
		"users.id   as user_id,"+
		"users.name as user_name "+
		"from posts "+
		"join users on users.id = posts.user_id "+
		"join posts_tags as pt on posts.id = pt.posts_id "+
		"join tags on tags.id = pt.tags_id "+
		"where pt.posts_id in (select p.id "+
		"from posts p "+
		"join posts_tags pt on p.id = pt.posts_id "+
		"join tags t on t.id = pt.tags_id "+
		"where t.name ~ ? )",
		tagNames).Scan(&data)
	if len(data) == 0 {
		return http.StatusNoContent
	}
	return c.TransformFromSearchByTag(data)
}
