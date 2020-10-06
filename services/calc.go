package services

type Calculator struct {
	Num1 int
	Num2 int
}

func (c *Calculator) Add() int {
	return c.Num1 + c.Num2
}

func (c *Calculator) Sub() int {
	return c.Num1 - c.Num2
}
