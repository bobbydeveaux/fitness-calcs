package main

import (
	"github.com/bobbydeveaux/fitness-calcs/calcs"
	"log"
)

func main() {
	log.Println("Test")
	bobby := calc.Person{
		Height:    71,
		Waist:     31.00,
		Neck:      15.00,
		Mass:      75.00,
		Hip:       0,
		Activity:  "low",
		Deficit:   "medium",
		Lifestyle: "lchf",
	}

	stephanie := calc.Person{
		Height:    65,
		Waist:     29.00,
		Neck:      12.50,
		Mass:      53.00,
		Hip:       35.00,
		Activity:  "low",
		Deficit:   "easy",
		Lifestyle: "lchf",
	}

	calc.CalcAll(bobby)
	calc.CalcAll(stephanie)

}
