package calc

import (
	"math"
)

func CalcBmr(leanmass float64) float64 {
	bmr := 370 + (21.6 * leanmass)
	return math.Round(bmr)
}
