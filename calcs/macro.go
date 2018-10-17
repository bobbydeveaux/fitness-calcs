package calc

import (
	"math"
)

func CalcMacros(calorieGoal float64, mass float64, leanmass float64, lifestyle string) (float64, float64, float64) {
	proteinGoal := mass * 2.2
	proteinCals := proteinGoal * 4
	carbGoal := 0.0
	carbCals := 0.0
	fatGoal := 0.0
	fatCals := 0.0

	switch lifestyle {
	case "lchf":
		carbCals = calorieGoal * 0.10
		carbGoal = carbCals / 4
		fatCals = calorieGoal - carbCals - proteinCals
		fatGoal = fatCals / 9
	case "hclf":
		fatCals = calorieGoal * 0.15
		fatGoal = fatCals / 9
		carbCals = calorieGoal - fatCals - proteinCals
		carbGoal = carbCals / 4
	case "psmf":
		fatGoal = 22.5
		carbGoal = 22.5
		proteinGoal = leanmass * 2.2 * 1.75
	case "balanced":
		remainingCals := calorieGoal - proteinCals
		fatCals = remainingCals * 0.5
		fatGoal = fatCals / 9
		carbCals = remainingCals * 0.5
		carbGoal = carbCals / 4
	}

	return math.Round(proteinGoal), math.Round(carbGoal), math.Round(fatGoal)
}
