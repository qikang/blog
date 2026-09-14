package handlers

import (
	"blog/models"
	"net/http"
	"strconv"
)

// IndexHandler handles the home page
func (h *BlogHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if pageNum, err := strconv.Atoi(p); err == nil && pageNum > 1 {
			page = pageNum
		}
	}

	postList, err := h.mdService.GetPagedPosts(page, h.pageSize)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Build page numbers
	var pageNumbers []struct {
		Number   int
		IsActive bool
	}
	for i := 1; i <= postList.TotalPages; i++ {
		pageNumbers = append(pageNumbers, struct {
			Number   int
			IsActive bool
		}{
			Number:   i,
			IsActive: i == page,
		})
	}

	data := struct {
		SiteName      string
		Posts         []models.Post
		Page          int
		TotalPages    int
		TotalCount    int
		LatestUpdate  string
		HasPagination bool
		HasPrev       bool
		HasNext       bool
		PrevPage      int
		NextPage      int
		PageNumbers   []struct {
			Number   int
			IsActive bool
		}
	}{
		SiteName:      h.siteName,
		Posts:         postList.Posts,
		Page:          postList.Page,
		TotalPages:    postList.TotalPages,
		TotalCount:    postList.TotalCount,
		LatestUpdate:  "",
		HasPagination: false,
		HasPrev:       page > 1,
		HasNext:       page < postList.TotalPages,
		PrevPage:      page - 1,
		NextPage:      page + 1,
		PageNumbers:   pageNumbers,
	}

	// Get latest post update time
	if len(postList.Posts) > 0 {
		data.LatestUpdate = postList.Posts[0].Date.Format("2006-01-02")
	}

	h.renderTemplate(w, r, "index.html", data)
}
