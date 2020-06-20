package mysportsfeeds

import "encoding/json"
import "fmt"
import "log"
import "net/http"

func GetDailyGames(password string) {
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

	for i := range games.Games {
		fmt.Printf("%d = %+v %+v\n", i, games.Games[i].Schedule.AwayTeam.Abbreviation, games.Games[i].Schedule.HomeTeam.Abbreviation)
	}

}
