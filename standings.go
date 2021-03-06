package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

// Standings is model for full-league standings.
type Standings struct {
	Leagues []League `json:"leagues"`
}

// League is model for leagues/conferences.
type League struct {
	Name      string     `json:"name"`
	Abbr      string     `json:"abbr"`
	Divisions []Division `json:"divisions"`
}

// Division is model for divisions.
type Division struct {
	Name  string `json:"name"`
	Abbr  string `json:"abbr"`
	Teams []Team `json:"teams"`
}

// Team is model for team standings.
type Team struct {
	Pos  int    `json:"pos"`
	Name string `json:"name"`
	Abbr string `json:"abbr"`
	W    int    `json:"w"`
	L    int    `json:"l"`
	D    int    `json:"d"`
}

// Get Current Standings (Baseball)
func bsStandings(bs Sport, year string) (Standings, error) {
	standings := Standings{}
	standings.Leagues = []League{
		{
			Name: "American League",
			Abbr: "AL",
			Divisions: []Division{
				{
					Name:  "East",
					Abbr:  "E",
					Teams: []Team{},
				},
				{
					Name:  "Central",
					Abbr:  "C",
					Teams: []Team{},
				},
				{
					Name:  "West",
					Abbr:  "W",
					Teams: []Team{},
				},
			},
		},
		{
			Name: "National League",
			Abbr: "NL",
			Divisions: []Division{
				{
					Name:  "East",
					Abbr:  "E",
					Teams: []Team{},
				},
				{
					Name:  "Central",
					Abbr:  "C",
					Teams: []Team{},
				},
				{
					Name:  "West",
					Abbr:  "W",
					Teams: []Team{},
				},
			},
		},
	}

	if year == "" {
		year = strconv.Itoa(time.Now().Year())
	}

	res, err := soup.Get("https://" + bs.Host + "/" +
		strings.Replace(bs.Standings, "{year}", year, 1))
	if err != nil {
		return standings, err
	}

	doc := soup.HTMLParse(res)
	for x, league := range standings.Leagues {
		for y, division := range league.Divisions {
			tables := doc.FindAll("div", "id", "all_standings_"+division.Abbr)
			teams := tables[x].FindAll("tr")
			for z, team := range teams {
				if z == 0 {
					continue
				}

				wins, _ := strconv.Atoi(team.Find("td", "data-stat", "W").FullText())
				losses, _ := strconv.Atoi(team.Find("td", "data-stat", "L").FullText())
				division.Teams = append(division.Teams, Team{
					Pos:  z,
					Name: team.Find("a").FullText(),
					Abbr: parseTeamURL(team.Find("a")),
					W:    wins,
					L:    losses,
					D:    0,
				})
			}
			standings.Leagues[x].Divisions[y] = division
		}
		standings.Leagues[x] = league
	}

	return standings, nil
}

func parseTeamURL(link soup.Root) string {
	url := strings.Split(link.Attrs()["href"], "/")
	return url[2]
}
