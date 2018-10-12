package calc

import ()

func CalcTdee(bmr float64, activityLevel string) float64 {

	factor := 1.0

	switch activityLevel {
	case "low":
		factor = 1.2
	case "light":
		factor = 1.375
	case "moderate":
		factor = 1.55
	case "very":
		factor = 1.725
	case "extreme":
		factor = 1.9
	}

	return bmr * factor
}
