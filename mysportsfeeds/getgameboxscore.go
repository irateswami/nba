package mysportsfeeds

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func GetBoxScore(password, year, date, awayTeam, homeTeam string) BoxScore {
	client := &http.Client{}

	//"https://api.mysportsfeeds.com/v2.1/pull/mlb/2019-regular/games/20190330-DET-TOR/boxscore.json"
	url := "https://api.mysportsfeeds.com/v2.1/pull/mlb/<YEAR>-regular/games/<DATE>-<AWAYTEAM>-<HOMETEAM>/boxscore.json"
	strings.Replace(url, "<YEAR>", year, -1)
	strings.Replace(url, "<DATE>", year, -1)
	strings.Replace(url, "<AWAYTEAM>", year, -1)
	strings.Replace(url, "<HOMETEAM>", year, -1)

	request, err := http.NewRequest("GET", url, nil)
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

	return *boxScore
}
