package services

import "errors"

type Calculator struct {
	Num1 int
	Num2 int
}

func (c *Calculator) Add() (int, error) {
	if c.Num1 < 0 || c.Num2 < 0 {
		return -1, errors.New("Tidak boleh negatif")
	}
	return c.Num1 + c.Num2, nil
}

func (c *Calculator) Sub() (int, error) {
	return c.Num1 - c.Num2, nil
}
