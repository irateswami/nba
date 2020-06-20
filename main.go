package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	msf "github.com/irateswami/nba/mysportsfeeds"
)

type secrets struct {
	Key string `json:"key"`
}

// Will never be modified, therefore, is global to package
var secretsStruct secrets

func init() {

	// Open the secrets
	secrets, err := os.Open("./secrets.json")
	if err != nil {
		log.Fatalf("failed to open secrets: %+s\n", err)
	}

	// Read in the byte string
	secretsBytes, err := ioutil.ReadAll(secrets)
	if err != nil {
		log.Fatalf("failed to read secrets: %+s\n", err)
	}

	json.Unmarshal(secretsBytes, &secretsStruct)

	err = secrets.Close()
	if err != nil {
		log.Fatalf("failed to close secrets: %+s\n", err)
	}

	fmt.Printf("%+v\n", secretsStruct)
}

func main() {

	msf.GetDailyGames(secretsStruct.Key)
}
