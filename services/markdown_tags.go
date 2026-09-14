package services

import (
	"blog/models"
	"sort"
	"strings"
)

// GetAllTags returns all unique tags from all posts
func (s *MarkdownService) GetAllTags() ([]string, error) {
	posts, err := s.GetAllPosts()
	if err != nil {
		return nil, err
	}

	tagMap := make(map[string]bool)
	for _, post := range posts {
		for _, tag := range post.Tags {
			tagMap[strings.TrimSpace(tag)] = true
		}
	}

	var tags []string
	for tag := range tagMap {
		tags = append(tags, tag)
	}
	sort.Strings(tags)

	return tags, nil
}

// GetPostsByTag returns all posts with a specific tag
func (s *MarkdownService) GetPostsByTag(tag string) ([]models.Post, error) {
	posts, err := s.GetAllPosts()
	if err != nil {
		return nil, err
	}

	var filtered []models.Post
	tag = strings.TrimSpace(tag)
	for _, post := range posts {
		for _, t := range post.Tags {
			if strings.TrimSpace(t) == tag {
				filtered = append(filtered, post)
				break
			}
		}
	}

	return filtered, nil
}
