package main

import (
	"github.com/anaskhan96/soup"
)

func bsTeams(bs Sport) ([]string, error) {
	teams := []string{}

	res, err := soup.Get("https://" + bs.Host + "/" + bs.Teams)
	if err != nil {
		return teams, err
	}
	doc := soup.HTMLParse(res)
	links := doc.FindAll("td", "data-stat", "franchise_name")
	for _, link := range links {
		teams = append(teams, link.FullText())
	}

	return teams, nil
}
