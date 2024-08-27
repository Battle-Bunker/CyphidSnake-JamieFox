package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/BattlesnakeOfficial/rules"
	"math"
	"slices"
)

func ManhattenDistance(head rules.Point, food rules.Point) float64 {
	return math.Abs(float64(head.X-food.X)) + math.Abs(float64(head.Y-food.Y))
}

func HeuristicFood(snapshot agent.GameSnapshot) float64 {
	var foodManhattenDistances = make([]float64, len(snapshot.Food()))
	for i, food := range snapshot.Food() {
		foodManhattenDistances[i] = ManhattenDistance(snapshot.You().Head(), food)
	}
	return float64(-slices.Min(foodManhattenDistances))
}
