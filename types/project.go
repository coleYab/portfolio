package types

import "time"

// this will contain the featured projects
type Project struct {
	ImageURL    string
	GithubURL   string
	Title       string
	Description string
	TechStacks  []string
	Stars       int
	Likes       int
}

// Repository represents a GitHub repository
type Repository struct {
	Name        string    `json:"name"`
	HTMLURL     string    `json:"html_url"`
	Description string    `json:"description"`
	PushedAt    time.Time `json:"pushed_at"`
	Stargazers  int       `json:"stargazers_count"`
	Language    string    `json:"language"`
}
