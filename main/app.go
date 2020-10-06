package main

import (
	"flag"
	"fmt"
	"gosimpletest/services"
)

func main() {
	num1 := flag.Int("num1", 0, "first number")
	num2 := flag.Int("num2", 0, "second number")
	opr := flag.String("opr", "add", "Calculator Operator")
	flag.Parse()
	switch *opr {
	case "add":
		myCalc := services.Calculator{
			*num1, *num2,
		}
		fmt.Println(myCalc.Add())
	case "sub":
		myCalc := services.Calculator{
			*num1, *num2,
		}
		fmt.Println(myCalc.Sub())
	}

}
