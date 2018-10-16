package calc

import (
	"log"
)

type Person struct {
	Height    float64
	Waist     float64
	Neck      float64
	Mass      float64
	Bia       float64
	Hip       float64
	Activity  string
	Deficit   string
	Lifestyle string
}

func CalcAll(p Person) {

	navyfat := CalcNavyFat(p.Height, p.Waist, p.Neck, p.Mass, p.Hip)
	fatPercent := navyfat
	if p.Bia > 0 {
		fatPercent = (navyfat + p.Bia) / 2
	}
	leanmass, fatmass := CalcLeanMass(p.Mass, fatPercent)
	bmr := CalcBmr(leanmass)
	tdee := CalcTdee(bmr, p.Activity)
	calorieGoal := CalcCalorieGoal(tdee, p.Deficit)
	proteinGoal, carbGoal, fatGoal := CalcMacros(calorieGoal, p.Mass, p.Lifestyle)
	log.Print("navyfat: ")
	log.Println(navyfat)
	log.Print("averagefat: ")
	log.Println(fatPercent)
	log.Print("leanmass: ")
	log.Println(leanmass)
	log.Print("fatmass: ")
	log.Println(fatmass)
	log.Print("bmr: ")
	log.Println(bmr)
	log.Print("tdee: ")
	log.Println(tdee)
	log.Print("calorie goal: ")
	log.Println(calorieGoal)
	log.Print("protein goal: ")
	log.Println(proteinGoal)
	log.Print("carb goal: ")
	log.Println(carbGoal)
	log.Print("fat goal: ")
	log.Println(fatGoal)
}
