package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CalculatorTestSuite struct {
	suite.Suite
	calcService Calculator
}

//Secara otomatis SetupTest dijalankan
func (suite *CalculatorTestSuite) SetupTest() {
	suite.calcService = Calculator{}
}

func (suite *CalculatorTestSuite) TestCalculator_Add() {
	testTable := []struct {
		num1 int
		num2 int
		res  int
	}{
		{1, 1, 2},
		{-1, 2, -1},
	}

	for _, table := range testTable {
		suite.calcService.Num1 = table.num1
		suite.calcService.Num2 = table.num2
		result, err := suite.calcService.Add()
		if err != nil {
			assert.NotNil(suite.T(), err)
			assert.Equal(suite.T(), table.res, result)
		} else {
			assert.Nil(suite.T(), err)
			assert.Equal(suite.T(), table.res, result)
		}

	}
}

func (suite *CalculatorTestSuite) TestCalculator_Sub() {
	suite.calcService.Num1 = 10
	suite.calcService.Num2 = 1

	result := suite.calcService.Sub()
	assert.Equal(suite.T(), 9, result)
}

func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
