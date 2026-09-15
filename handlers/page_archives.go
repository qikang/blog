package handlers

import (
	"blog/models"
	"net/http"
	"sort"
	"strconv"
)

// ArchiveYear 表示某一年的文章列表(有序结构,避免 map 遍历乱序)
type ArchiveYear struct {
	Year  string
	Posts []models.Post
}

// ArchivesHandler handles the archives page
func (h *BlogHandler) ArchivesHandler(w http.ResponseWriter, r *http.Request) {
	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if pageNum, err := strconv.Atoi(p); err == nil && pageNum > 1 {
			page = pageNum
		}
	}

	// 每页大小:支持 10 / 20 / 50,默认 10
	pageSize := 10
	if s := r.URL.Query().Get("size"); s != "" {
		if n, err := strconv.Atoi(s); err == nil {
			switch n {
			case 10, 20, 50:
				pageSize = n
			}
		}
	}

	allPosts, _ := h.mdService.GetAllPosts()
	totalCount := len(allPosts)
	totalPages := (totalCount + pageSize - 1) / pageSize

	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= totalCount {
		start = 0
		page = 1
		end = pageSize
	}
	if end > totalCount {
		end = totalCount
	}

	pagedPosts := allPosts[start:end]

	// 按年分组到 map
	archiveMap := make(map[string][]models.Post)
	for _, post := range pagedPosts {
		year := post.Date.Format("2006")
		archiveMap[year] = append(archiveMap[year], post)
	}

	// 提取年份并倒序排序(最新年份在前)
	years := make([]string, 0, len(archiveMap))
	for y := range archiveMap {
		years = append(years, y)
	}
	sort.Strings(years) // 升序

	// 转成 ArchiveYear 切片,年份倒序;每年内文章保持传入顺序(整体已按 date 倒序)
	archive := make([]ArchiveYear, 0, len(years))
	for i := len(years) - 1; i >= 0; i-- {
		archive = append(archive, ArchiveYear{
			Year:  years[i],
			Posts: archiveMap[years[i]],
		})
	}

	// 构建分页页码(带上当前 size,保证切换页码后每页大小不丢)
	pageNumbers := make([]PageLink, 0, totalPages)
	for i := 1; i <= totalPages; i++ {
		pageNumbers = append(pageNumbers, PageLink{
			Number:   i,
			IsActive: i == page,
			Size:     pageSize,
		})
	}

	data := ArchivesPageData{
		SiteName:      h.siteName,
		Archive:       archive,
		TotalCount:    totalCount,
		Page:          page,
		TotalPages:    totalPages,
		PageSize:      pageSize,
		SizeOptions:   []int{10, 20, 50},
		HasPagination: true,
		HasPrev:       page > 1,
		HasNext:       page < totalPages,
		PrevPage:      page - 1,
		NextPage:      page + 1,
		PageNumbers:   pageNumbers,
	}

	h.renderTemplate(w, r, "archives.html", data)
}

// PageLink 描述分页器里的页码链接
type PageLink struct {
	Number   int
	IsActive bool
	Size     int
}

// ArchivesPageData 聚合 archives 页所需的模板数据
type ArchivesPageData struct {
	SiteName      string
	Archive       []ArchiveYear
	TotalCount    int
	Page          int
	TotalPages    int
	PageSize      int
	SizeOptions   []int
	HasPagination bool
	HasPrev       bool
	HasNext       bool
	PrevPage      int
	NextPage      int
	PageNumbers   []PageLink
}