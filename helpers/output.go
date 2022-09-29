package helpers

import (
	"fmt"
)

type Output struct {
}

func (c Output) Console(res int) error {
	_, err := fmt.Println(res)
	return err
}
func (c Output) Error(err error) error {
	_, err = fmt.Println(err)
	return err
}
