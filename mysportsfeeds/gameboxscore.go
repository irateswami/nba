package mysportsfeeds

import (
	"encoding/json"
	"log"
	"net/http"
)

type Fielder struct {
	Runs         int
	HomeRuns     int
	RunsBattedIn int
	StolenBases  int
	Hits         int
	AtBats       int
}

type Pitcher struct {
	Wins           int
	Saves          int
	StrikeOuts     int
	EarnedRuns     int
	InningsPitched float32
	Walks          int
	Hits           int
}

func GetBoxScore(password string) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://api.mysportsfeeds.com/v2.1/pull/mlb/2019-regular/games/20190330-DET-TOR/boxscore.json", nil)
	if err != nil {
		log.Printf("request failed: %v\n", err)
	}

	request.SetBasicAuth(password, "MYSPORTSFEEDS")

	response, err := client.Do(request)
	if err != nil {
		log.Printf("response failed: %v\n", err)
	}

	boxScore := new(BoxScore)
	json.NewDecoder(response.Body).Decode(boxScore)

	var fielders []Fielder
	var pitchers []Pitcher

	for i := range boxScore.Stats.Home.Players {
	}

}
