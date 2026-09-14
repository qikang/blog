package services

import (
	"blog/models"
	"os"
	"path/filepath"
	"sort"
	"time"
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

	// Sort: first by folder (folders first, then files in root), then by date within each group
	sort.Slice(posts, func(i, j int) bool {
		// Get category for each post
		catI := posts[i].Category
		catJ := posts[j].Category

		// Both have categories - compare by category's latest post date
		if catI != "" && catJ != "" {
			if catI != catJ {
				// Compare by folder's latest post time
				latestI := getFolderLatestTime(posts, catI)
				latestJ := getFolderLatestTime(posts, catJ)
				if !latestI.Equal(latestJ) {
					return latestI.After(latestJ)
				}
				return catI < catJ // Alphabetical if same time
			}
			// Same folder - compare by date
			return posts[i].Date.After(posts[j].Date)
		}

		// One has category, one doesn't - folders first
		if catI != "" && catJ == "" {
			latestJ := posts[j].Date // Root file's own date
			latestI := getFolderLatestTime(posts, catI)
			return latestI.After(latestJ)
		}
		if catI == "" && catJ != "" {
			latestI := posts[i].Date // Root file's own date
			latestJ := getFolderLatestTime(posts, catJ)
			return latestI.After(latestJ)
		}

		// Both in root - compare by date
		return posts[i].Date.After(posts[j].Date)
	})

	return posts, nil
}

// getFolderLatestTime returns the latest date among all posts in a folder
func getFolderLatestTime(posts []models.Post, folder string) time.Time {
	var latest time.Time
	for _, p := range posts {
		if p.Category == folder && p.Date.After(latest) {
			latest = p.Date
		}
	}
	return latest
}

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
