package services

import (
	"blog/models"
	"fmt"
)

// GetArchive returns posts grouped by year
func (s *MarkdownService) GetArchive() (map[string][]models.Post, error) {
	posts, err := s.GetAllPosts()
	if err != nil {
		return nil, err
	}

	archive := make(map[string][]models.Post)
	for _, post := range posts {
		key := fmt.Sprintf("%d年", post.Date.Year())
		archive[key] = append(archive[key], post)
	}

	return archive, nil
}
