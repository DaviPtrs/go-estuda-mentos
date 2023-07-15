package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name     string
	m_played int
	m_won    int
	m_tied   int
	m_lost   int
	points   int
}

func sortTeamSlice(teamSlice []Team) {
	sort.Slice(teamSlice, func(i, j int) bool {
		if teamSlice[i].points == teamSlice[j].points {
			return teamSlice[i].name < teamSlice[j].name
		}
		return teamSlice[i].points > teamSlice[j].points
	})
}

func cleanStringSlice(slice []string) []string {
	clear := []string{}
	for _, elem := range slice {
		if len(elem) != 0 && !strings.HasPrefix(elem, "#") {
			clear = append(clear, elem)
		}
	}
	return clear
}

func parseGames(gameString []string) (map[string]*Team, error) {
	results := map[string]*Team{}

	for _, elem := range gameString {
		gameInfo := strings.Split(elem, ";")
		if len(gameInfo) != 3 {
			return results, errors.New("Invalid")
		} else {
			team1, check := results[gameInfo[0]]
			if !check {
				team1 = new(Team)
				team1.name = gameInfo[0]
				results[gameInfo[0]] = team1
			}

			team2, check := results[gameInfo[1]]
			if !check {
				team2 = new(Team)
				team2.name = gameInfo[1]
				results[gameInfo[1]] = team2
			}

			team1.m_played++
			team2.m_played++

			matchResult := gameInfo[2]
			if matchResult == "win" {
				team1.m_won++
				team1.points += 3
				team2.m_lost++
			} else if matchResult == "loss" {
				team2.m_won++
				team2.points += 3
				team1.m_lost++
			} else if matchResult == "draw" {
				team1.m_tied++
				team2.m_tied++
				team1.points++
				team2.points++
			} else {
				return results, errors.New("Invalid")
			}
		}
	}
	return results, nil
}

func Tally(reader io.Reader, writer io.Writer) error {
	buff, _ := io.ReadAll(reader)
	games := strings.Split(string(buff), "\n")
	games = cleanStringSlice(games)
	results, err := parseGames(games)
	if err != nil {
		return err
	}

	resultsSlice := []Team{}
	for _, team := range results {
		resultsSlice = append(resultsSlice, *team)
	}
	sortTeamSlice(resultsSlice)

	output := fmt.Sprintf("%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	for _, team := range resultsSlice {
		output += fmt.Sprintf("%-30s | %2v | %2v | %2v | %2v | %2v\n",
			team.name, team.m_played, team.m_won,
			team.m_tied, team.m_lost, team.points)
	}
	_, err = fmt.Fprint(writer, output)

	if err != nil {
		return err
	}

	return nil

}
