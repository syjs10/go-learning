package github

import (
	"net/http"
	"net/url"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number  int
	HTMLURL string
	Title   string
	State   string
	User    *User
	Create  time.Time `json:"create_at"`
	Body    string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(string.Join(terms, " "))
	resq, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
}
