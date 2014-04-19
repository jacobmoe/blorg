package model

import "strings"
import "regexp"

type Post struct {
	Title   string   `json:"title"`
	Section []string `json:"sections"`
	Date    string   `json:"date"`
	page    *Page
}

// move to gorg
func AsciiTableToHtml(ascii string) string {
	r := regexp.MustCompile(`\A\s*\|\-+\++.*\|$`)
	html := "<table class=\"table table-hover\">"

	for rowIndex, row := range strings.Split(ascii, "\n") {
		if r.MatchString(row) {
			continue
		}

		html = html + "<tr>"
		for _, col := range strings.Split(row, "|") {
			if rowIndex == 0 {
				html = html + "<th>" + col + "</th>"
			} else {
				html = html + "<td>" + col + "</td>"
			}
		}
	}

	return html + "</table>"
}
