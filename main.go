package main

import (
	"github.com/bobbydeveaux/fitness-calcs/calcs"
	"log"
)

func main() {
	log.Println("Test")
	navyfat := calc.CalcNavyFat(71*2.54, 31, 15, 75, 0)
	fatPercent := navyfat
	leanmass, fatmass := calc.CalcLeanMass(75, fatPercent)
	bmr := calc.CalcBmr(leanmass)
	tdee := calc.CalcTdee(bmr, "low")
	calorieGoal := calc.CalcCalorieGoal(tdee, "medium")
	log.Print("navyfat: ")
	log.Println(navyfat)
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
}
