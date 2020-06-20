package mysportsfeeds

import (
	"time"
)

type Games struct {
	LastUpdatedOn time.Time `json:"lastUpdatedOn"`
	Games         []struct {
		Schedule struct {
			ID        int         `json:"id"`
			StartTime time.Time   `json:"startTime"`
			EndedTime interface{} `json:"endedTime"`
			AwayTeam  struct {
				ID           int    `json:"id"`
				Abbreviation string `json:"abbreviation"`
			} `json:"awayTeam"`
			HomeTeam struct {
				ID           int    `json:"id"`
				Abbreviation string `json:"abbreviation"`
			} `json:"homeTeam"`
			Venue struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"venue"`
			VenueAllegiance          string        `json:"venueAllegiance"`
			ScheduleStatus           string        `json:"scheduleStatus"`
			OriginalStartTime        interface{}   `json:"originalStartTime"`
			DelayedOrPostponedReason interface{}   `json:"delayedOrPostponedReason"`
			PlayedStatus             string        `json:"playedStatus"`
			Attendance               interface{}   `json:"attendance"`
			Officials                []interface{} `json:"officials"`
			Broadcasters             []interface{} `json:"broadcasters"`
			Weather                  interface{}   `json:"weather"`
		} `json:"schedule"`
		Score struct {
			CurrentInning       interface{} `json:"currentInning"`
			CurrentInningHalf   interface{} `json:"currentInningHalf"`
			CurrentIntermission interface{} `json:"currentIntermission"`
			PlayStatus          interface{} `json:"playStatus"`
			AwayScoreTotal      int         `json:"awayScoreTotal"`
			AwayHitsTotal       int         `json:"awayHitsTotal"`
			AwayErrorsTotal     int         `json:"awayErrorsTotal"`
			HomeScoreTotal      int         `json:"homeScoreTotal"`
			HomeHitsTotal       int         `json:"homeHitsTotal"`
			HomeErrorsTotal     int         `json:"homeErrorsTotal"`
			Innings             []struct {
				InningNumber int `json:"inningNumber"`
				AwayScore    int `json:"awayScore"`
				HomeScore    int `json:"homeScore"`
			} `json:"innings"`
		} `json:"score"`
	} `json:"games"`
	References struct {
		TeamReferences []struct {
			ID           int    `json:"id"`
			City         string `json:"city"`
			Name         string `json:"name"`
			Abbreviation string `json:"abbreviation"`
			HomeVenue    struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"homeVenue"`
			TeamColoursHex      []string `json:"teamColoursHex"`
			SocialMediaAccounts []struct {
				MediaType string `json:"mediaType"`
				Value     string `json:"value"`
			} `json:"socialMediaAccounts"`
			OfficialLogoImageSrc string `json:"officialLogoImageSrc"`
		} `json:"teamReferences"`
		VenueReferences []struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			City           string `json:"city"`
			Country        string `json:"country"`
			GeoCoordinates struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"geoCoordinates"`
			CapacitiesByEventType []struct {
				EventType string `json:"eventType"`
				Capacity  int    `json:"capacity"`
			} `json:"capacitiesByEventType"`
			PlayingSurface     string `json:"playingSurface"`
			BaseballDimensions []struct {
				DimensionType string `json:"dimensionType"`
				DistanceFeet  int    `json:"distanceFeet"`
			} `json:"baseballDimensions"`
			HasRoof            bool `json:"hasRoof"`
			HasRetractableRoof bool `json:"hasRetractableRoof"`
		} `json:"venueReferences"`
	} `json:"references"`
}
