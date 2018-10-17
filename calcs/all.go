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
	Calcs     Calcs
}

type Calcs struct {
	NavyFat     float64
	FatPercent  float64
	LeanMass    float64
	FatMass     float64
	Bmr         float64
	Tdee        float64
	CalorieGoal float64
	ProteinGoal float64
	CarbGoal    float64
	FatGoal     float64
}

func (c *Calcs) Check(p Person) bool {
	totCals := (c.ProteinGoal+c.CarbGoal)*4 + c.FatGoal*9
	if p.Lifestyle == "psmf" {
		c.CalorieGoal = totCals
		return true
	}
	if c.CalorieGoal > totCals+5 || c.CalorieGoal < totCals-5 {
		log.Println("Calorie calc mishap")
		return false
	}

	return true
}

func CalcAll(p Person) Person {

	navyfat := CalcNavyFat(p.Height, p.Waist, p.Neck, p.Mass, p.Hip)
	fatPercent := navyfat
	if p.Bia > 0 {
		fatPercent = (navyfat + p.Bia) / 2
	}
	leanmass, fatmass := CalcLeanMass(p.Mass, fatPercent)
	bmr := CalcBmr(leanmass)
	tdee := CalcTdee(bmr, p.Activity)
	calorieGoal := CalcCalorieGoal(tdee, p.Deficit)
	proteinGoal, carbGoal, fatGoal := CalcMacros(calorieGoal, p.Mass, leanmass, p.Lifestyle)
	calcs := Calcs{
		NavyFat:     navyfat,
		FatPercent:  fatPercent,
		LeanMass:    leanmass,
		FatMass:     fatmass,
		Bmr:         bmr,
		Tdee:        tdee,
		CalorieGoal: calorieGoal,
		ProteinGoal: proteinGoal,
		FatGoal:     fatGoal,
		CarbGoal:    carbGoal,
	}

	if !calcs.Check(p) {
		return p
	}

	p.Calcs = calcs

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

	return p
}
