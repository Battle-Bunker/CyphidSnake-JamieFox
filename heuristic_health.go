package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
)

// heuristicHealth calculates the sum of health for all snakes in your team,
// including the player's snake.
// Calculates all of the health of all the agents in your team and returns it as an integer. (written by jacob)
func HeuristicHealth(snapshot agent.GameSnapshot) float64 {
	var totalHealth float64 = 0
	for _, snake := range snapshot.YourTeam() {
		totalHealth += float64(snake.Health())
	}
	return totalHealth
}
