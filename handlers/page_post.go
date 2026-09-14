package handlers

import (
	"blog/models"
	"net/http"

	"github.com/gorilla/mux"
)

// PostHandler handles single post page
func (h *BlogHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]

	if slug == "" {
		http.NotFound(w, r)
		return
	}

	post, err := h.mdService.GetPostBySlug(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Get prev and next posts
	allPosts, _ := h.mdService.GetAllPosts()
	var prevPost, nextPost *models.Post

	for i, p := range allPosts {
		if p.Slug == slug {
			if i > 0 {
				prevPost = &allPosts[i-1]
			}
			if i < len(allPosts)-1 {
				nextPost = &allPosts[i+1]
			}
			break
		}
	}

	data := struct {
		SiteName string
		Post     *models.Post
		PrevPost *models.Post
		NextPost *models.Post
	}{
		SiteName: h.siteName,
		Post:     post,
		PrevPost: prevPost,
		NextPost: nextPost,
	}

	h.renderTemplate(w, r, "post.html", data)
}
