package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

// Get Current Standings (Baseball)
func bsStandings(bs Sport, year string) (map[string][]string, error) {
	stands := make(map[string][]string)
	leagues := []string{"AL", "NL"}
	divisions := []string{"E", "C", "W"}
	if year == "" {
		year = strconv.Itoa(time.Now().Year())
	}

	res, err := soup.Get("https://" + bs.Host + "/" +
		strings.Replace(bs.Standings, "{year}", year, 1))
	if err != nil {
		return stands, err
	}

	doc := soup.HTMLParse(res)
	for _, div := range divisions {
		tables := doc.FindAll("div", "id", "all_standings_"+div)
		for x, table := range tables {
			name := table.Find("h2").Text()
			divName := leagues[x] + " " + strings.Replace(name, " Division", "", 1)
			teams := table.FindAll("a")
			for _, team := range teams {
				stands[divName] = append(stands[divName], team.Text())
			}
		}
	}
	return stands, nil
}
