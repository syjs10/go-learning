package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"github.com/syjs10/day11/github"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	// fmt.Println(t)
	return int(time.Since(t).Hours() / 24)
}

func main() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
