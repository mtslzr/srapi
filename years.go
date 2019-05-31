package main

import (
	"github.com/anaskhan96/soup"
)

// Get All Years (Baseball)
func bsYears(bs Sport) (years []string) {
	res, err := soup.Get("https://" + bs.Host + "/" + bs.Years)
	checkError(err)

	doc := soup.HTMLParse(res)
	links := doc.FindAll("th", "data-stat", "year_ID")
	for _, link := range links {
		if link.Text() != "" {
			years = append(years, link.Text())
		}
	}
	return
}
