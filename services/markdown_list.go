package services

import (
	"blog/models"
	"os"
	"path/filepath"
	"sort"
)

// GetAllPosts returns all posts sorted by folder first, then by date (newest first)
func (s *MarkdownService) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post

	// Use filepath.Walk to recursively find all .md files
	err := filepath.Walk(s.postsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories
		if info.IsDir() {
			return nil
		}
		// Only process .md files
		if filepath.Ext(path) != ".md" {
			return nil
		}
		// Skip about_me.md - it's only for About page
		if filepath.Base(path) == "about_me.md" {
			return nil
		}
		post, err := s.ParsePost(path)
		if err != nil {
			return nil // Skip files that can't be parsed
		}
		posts = append(posts, *post)
		return nil
	})
	if err != nil {
		return nil, err
	}

	// 全部按文章日期倒序排(归档、首页等所有列表都用此排序)
// 注意 Go 1.21+ 推荐 slices.SortStableFunc,这里保留 sort.Slice + 稳定性的兜底(按 Slug 打破平局)
	sort.SliceStable(posts, func(i, j int) bool {
		if posts[i].Date.Equal(posts[j].Date) {
			// 同一日期内,按 Slug 字典序,保证排序稳定
			return posts[i].Slug < posts[j].Slug
		}
		return posts[i].Date.After(posts[j].Date)
	})

	return posts, nil
}

// getFolderLatestTime 已移除:历史实现里按"分类分组"排序,导致跨分类时日期错乱;
// 现统一改为整体按文章 Date 倒序。

// GetPagedPosts returns posts with pagination
func (s *MarkdownService) GetPagedPosts(page, pageSize int) (*models.PostList, error) {
	allPosts, err := s.GetAllPosts()
	if err != nil {
		return nil, err
	}

	totalCount := len(allPosts)
	totalPages := (totalCount + pageSize - 1) / pageSize

	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= totalCount {
		return &models.PostList{
			Posts:      []models.Post{},
			Page:       page,
			TotalPages: totalPages,
			TotalCount: totalCount,
		}, nil
	}
	if end > totalCount {
		end = totalCount
	}

	return &models.PostList{
		Posts:      allPosts[start:end],
		Page:       page,
		TotalPages: totalPages,
		TotalCount: totalCount,
	}, nil
}

// GetRecentPosts returns the N most recent posts
func (s *MarkdownService) GetRecentPosts(n int) ([]models.Post, error) {
	posts, err := s.GetAllPosts()
	if err != nil {
		return nil, err
	}
	if len(posts) > n {
		return posts[:n], nil
	}
	return posts, nil
}
