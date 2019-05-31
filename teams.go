package main

import (
	"github.com/anaskhan96/soup"
)

// Get All Teams (Baseball)
func bsTeams(bs Sport) (teams []string, err error) {
	res, err := soup.Get("https://" + bs.Host + "/" + bs.Teams)
	if err != nil {
		return
	}

	doc := soup.HTMLParse(res)
	links := doc.FindAll("td", "data-stat", "franchise_name")
	for _, link := range links {
		teams = append(teams, link.FullText())
	}
	return
}
