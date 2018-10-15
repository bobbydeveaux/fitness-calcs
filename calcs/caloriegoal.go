package calc

import (
	"math"
)

func CalcCalorieGoal(tdee float64, goal string) float64 {

	factor := 1.0

	switch goal {
	case "leangains":
		factor = 1.2
	case "maintenance":
		factor = 1.0
	case "slow":
		factor = 0.95
	case "easy":
		factor = 0.90
	case "medium":
		factor = 0.85
	case "hard":
		factor = 0.80
	}

	return math.Round(tdee * factor)
}
