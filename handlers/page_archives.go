package handlers

import (
	"blog/models"
	"net/http"
	"strconv"
)

// ArchivesHandler handles the archives page
func (h *BlogHandler) ArchivesHandler(w http.ResponseWriter, r *http.Request) {
	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if pageNum, err := strconv.Atoi(p); err == nil && pageNum > 1 {
			page = pageNum
		}
	}

	pageSize := 10

	allPosts, _ := h.mdService.GetAllPosts()
	totalCount := len(allPosts)
	totalPages := (totalCount + pageSize - 1) / pageSize

	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= totalCount {
		start = 0
		page = 1
	}
	if end > totalCount {
		end = totalCount
	}

	pagedPosts := allPosts[start:end]

	// Rebuild archive for current page
	archive := make(map[string][]models.Post)
	for _, post := range pagedPosts {
		year := post.Date.Format("2006")
		archive[year] = append(archive[year], post)
	}

	// Build page numbers
	var pageNumbers []struct {
		Number   int
		IsActive bool
	}
	for i := 1; i <= totalPages; i++ {
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
		Archive       map[string][]models.Post
		TotalCount    int
		Page          int
		TotalPages    int
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
		Archive:       archive,
		TotalCount:    totalCount,
		Page:          page,
		TotalPages:    totalPages,
		HasPagination: true,
		HasPrev:       page > 1,
		HasNext:       page < totalPages,
		PrevPage:      page - 1,
		NextPage:      page + 1,
		PageNumbers:   pageNumbers,
	}

	h.renderTemplate(w, r, "archives.html", data)
}
