package github

import "time"

const SearchIssuesURL = "https://api.github.com/search/issues"

// Maybe Pointers in structs can be compared to foreign keys
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
	Url       string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
