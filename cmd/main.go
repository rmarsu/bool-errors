package main

import (
	boolerrors "boolerrors/source"
	"fmt"
)


func main() {
	error1 := boolerrors.NewError()
	error2 := boolerrors.NoError()
	megaError := boolerrors.CombineErrors(error1, error2)
	fmt.Println(megaError.Error())
	fmt.Println(megaError.BoolButCool())
	fmt.Println(megaError.Negative().Error())
	//megaError.Secret()
}