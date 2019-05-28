package main

import (
	"github.com/anaskhan96/soup"
)

func bsYears(bs Sport) ([]string, error) {
	years := []string{}

	res, err := soup.Get("https://" + bs.Host + "/" + bs.Years)
	if err != nil {
		return years, err
	}
	doc := soup.HTMLParse(res)
	links := doc.FindAll("th", "data-stat", "year_ID")
	for _, link := range links {
		if link.Text() != "" {
			years = append(years, link.Text())
		}
	}

	return years, nil
}
