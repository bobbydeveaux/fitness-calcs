package calc

import (
	"log"
)

type Person struct {
	Height    float64 `json:"height"`
	Waist     float64 `json:"waist"`
	Neck      float64 `json:"neck"`
	Mass      float64 `json:"mass"`
	Bia       float64 `json:"bia"`
	Hip       float64 `json:"hip"`
	Activity  string  `json:"activity"`
	Deficit   string  `json:"deficit"`
	Lifestyle string  `json:"lifestyle"`
	Calcs     Calcs   `json:"calcs,omitempty"`
}

func (p *Person) Calc() bool {

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

	if !calcs.Check(*p) {
		return false
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

	return true
}

type Calcs struct {
	NavyFat     float64 `json:"navyfat,omitempty"`
	FatPercent  float64 `json:"fatpercent,omitempty"`
	LeanMass    float64 `json:"leanmass,omitempty"`
	FatMass     float64 `json:"fatmass,omitempty"`
	Bmr         float64 `json:"bmr,omitempty"`
	Tdee        float64 `json:"tdee,omitempty"`
	CalorieGoal float64 `json:"caloriegoal,omitempty"`
	ProteinGoal float64 `json:"proteingoal,omitempty"`
	CarbGoal    float64 `json:"carbgoal,omitempty"`
	FatGoal     float64 `json:"fatgoal,omitempty"`
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
