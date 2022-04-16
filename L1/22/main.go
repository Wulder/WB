package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	result := new(big.Int)

	a.SetString("1000000000000000000000000000000000000", 10)
	b.SetString("2", 10)

	result.Add(a, b)
	fmt.Println("Add: ", result)

	result.Div(a, b)
	fmt.Println("Div: ", result)

	result.Mul(a, b)
	fmt.Println("Mul: ", result)

	result.Sub(a, b)
	fmt.Println("Sub: ", result)
}
