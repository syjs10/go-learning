package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number   int
	HTMLURL  string
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"created_at"`
	Body     string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resq, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resq.StatusCode != http.StatusOK {
		resq.Body.Close()
		return nil, fmt.Errorf("Search query faild: %s", resq.Status)
	}
	var result IssueSearchResult
	if err := json.NewDecoder(resq.Body).Decode(&result); err != nil {
		resq.Body.Close()
		return nil, err
	}
	resq.Body.Close()
	return &result, nil
}
