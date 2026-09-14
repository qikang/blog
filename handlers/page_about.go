package handlers

import (
	"blog/models"
	"net/http"
	"path/filepath"
)

// AboutHandler handles the about page
func (h *BlogHandler) AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Read about_me.md file
	filename := filepath.Join(h.mdService.GetPostsDir(), "about_me.md")
	post, err := h.mdService.ParsePost(filename)
	if err != nil {
		// If file doesn't exist, use default content
		post = &models.Post{
			Title:   "关于我",
			Content: "欢迎访问关于我页面",
		}
	}

	data := struct {
		SiteName string
		Post     *models.Post
	}{
		SiteName: h.siteName,
		Post:     post,
	}

	h.renderTemplate(w, r, "about.html", data)
}
