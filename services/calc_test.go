package services

import "testing"

func TestCalculator_Add(t *testing.T) {
	t.Run("Calculator Add operator testing", func(t *testing.T) {
		numSample1 := 10
		numSample2 := 20

		calc := Calculator{
			Num1: numSample1,
			Num2: numSample2,
		}
		if calc.Add() != 30 {
			t.Error("It should return 30")
		}
	})
}
func TestCalculator_Sub(t *testing.T) {
	t.Run("Calculator Subtraction operator testing", func(t *testing.T) {
		numSample1 := 3
		numSample2 := 2

		calc := Calculator{
			Num1: numSample1,
			Num2: numSample2,
		}
		if calc.Sub() != 1 {
			t.Error("It should return 1")
		}
	})
}
