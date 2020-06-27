package mysportsfeeds

func ParseBoxScore(boxScore BoxScore) {

	allPlayers := []co
	awayPlayers := boxScore.Stats.Away.Players
	homePlayers := boxScore.Stats.Away.Players

	for i := range awayPlayers {

		if awayPlayers[i].Player.Position == "P" {
			pitcher := new(Pitcher)
			pitcher.Wins = awayPlayers[i].PlayerStats[0].Pitching.Wins
			pitcher.Saves = awayPlayers[i].PlayerStats[0].Pitching.Saves
			pitcher.StrikeOuts = awayPlayers[i].PlayerStats[0].Pitching.PitcherStrikeouts
			pitcher.EarnedRuns = awayPlayers[i].PlayerStats[0].Pitching.EarnedRunsAllowed
			pitcher.InningsPitched = awayPlayers[i].PlayerStats[0].Pitching.InningsPitched

		}

	}

	for i := range homePlayers {
		if homePlayers[i].Player.Position == "P" {

		}

	}

}
