package calc

import (
	"math"
)

func CalcNavyFat(height float64, naval float64, neck float64, mass float64, hip float64) float64 {

	// Female
	if hip > 0 {
		return math.Log10(naval-hip-neck) - 97.684*math.Log10(height/2.54) - 78.387
	}

	// Male
	return 86.010*math.Log10(naval-neck) - 70.041*math.Log10(height/2.54) + 36.76
}
