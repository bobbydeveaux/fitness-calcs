package main

import (
	"github.com/bobbydeveaux/fitness-calcs/calcs"
)

func main() {

	bobby := calc.Person{
		Height:    71,
		Waist:     31.00,
		Neck:      15.00,
		Mass:      75.00,
		Bia:       14.00,
		Hip:       0,
		Activity:  "low",
		Deficit:   "medium",
		Lifestyle: "lchf",
	}

	steph := calc.Person{
		Height:    65,
		Waist:     29.00,
		Neck:      12.50,
		Mass:      53.00,
		Bia:       23.00,
		Hip:       35.00,
		Activity:  "low",
		Deficit:   "easy",
		Lifestyle: "lchf",
	}

	lee := calc.Person{
		Height:    71.6,
		Waist:     34.00,
		Neck:      15.50,
		Mass:      78.50,
		Bia:       00.00,
		Hip:       00.00,
		Activity:  "low",
		Deficit:   "maintain",
		Lifestyle: "lchf",
	}

	lee.Calc()
	bobby.Calc()
	steph.Calc()
}
