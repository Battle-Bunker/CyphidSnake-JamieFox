package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	_ "log"
)

func HeuristicEat(snapshot agent.GameSnapshot) float64 {
	if snapshot.You().Health() == 100 {
		return 5.0
	}
	return 0
}
