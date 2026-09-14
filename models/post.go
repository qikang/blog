package models

import (
	"html/template"
	"time"
)

// Post represents a blog post
type Post struct {
	Slug      string
	Category  string
	Title     string
	Date      time.Time
	Author    string
	Tags      []string
	Summary   string
	Content   string
	HTML      template.HTML
}

// PostList represents a list of posts with pagination info
type PostList struct {
	Posts      []Post
	Page       int
	TotalPages int
	TotalCount int
}
