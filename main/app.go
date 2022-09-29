package main

import (
	"flag"
	"gosimpletest/helpers"
	"gosimpletest/services"
)

func main() {
	num1 := flag.Int("num1", 0, "first number")
	num2 := flag.Int("num2", 0, "second number")
	opr := flag.String("opr", "add", "Calculator Operator")
	flag.Parse()
	showResult := helpers.Output{}
	var result = 0
	switch *opr {
	case "add":
		myCalc := services.Calculator{
			*num1, *num2,
		}
		res, err := myCalc.Add()
		if err != nil {
			err := showResult.Error(err)
			if err != nil {
				return
			}
		} else {
			result = res
		}
	case "sub":
		myCalc := services.Calculator{
			*num1, *num2,
		}
		res, err := myCalc.Sub()
		if err != nil {
			err := showResult.Error(err)
			if err != nil {
				return
			}
		} else {
			result = res
		}
	}
	err := showResult.Console(result)
	if err != nil {
		return
	}
}
