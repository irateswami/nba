package mysportsfeeds

import (
	"encoding/json"
	"log"
	"net/http"
)

type Game struct {
	HomeTeam string
	AwayTeam string
}

func getDailyGames(password string) []Game {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.mysportsfeeds.com/v2.1/pull/mlb/2019-regular/date/20190330/games.json", nil)
	if err != nil {
		log.Printf("request failed: %v\n", err)
	}

	req.SetBasicAuth(password, "MYSPORTSFEEDS")

	response, err := client.Do(req)
	if err != nil {
		log.Printf("response failed: %v\n", err)
	}

	games := new(Games)

	json.NewDecoder(response.Body).Decode(games)

	var allDailyGames []Game

	for i := range games.Games {
		game := Game{
			HomeTeam: games.Games[i].Schedule.HomeTeam.Abbreviation,
			AwayTeam: games.Games[i].Schedule.AwayTeam.Abbreviation,
		}
		allDailyGames = append(allDailyGames, game)
	}

	return allDailyGames
}
