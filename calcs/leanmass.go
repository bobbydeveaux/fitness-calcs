package calc

import ()

func CalcLeanMass(mass float64, fatPercent float64) (float64, float64) {
	fatmass := mass * (fatPercent / 100)
	leanmass := mass - fatmass

	return leanmass, fatmass
}
