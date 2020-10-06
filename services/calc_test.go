package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	t.Run("Calculator Add operator testing, it should return int", func(t *testing.T) {
		numSample1 := 10
		numSample2 := 20

		calc := Calculator{
			Num1: numSample1,
			Num2: numSample2,
		}
		result, err := calc.Add()
		assert.Nil(t, err)
		assert.Equal(t, 30, result)
	})

	t.Run("It should return error when receive negative value", func(t *testing.T) {
		numSample1 := -1
		numSample2 := 10
		calc := Calculator{
			Num1: numSample1,
			Num2: numSample2,
		}
		result, err := calc.Add()

		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
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
		result := calc.Sub()
		assert.Equal(t, 1, result)
	})
}
