//	Разработать программу, которая перемножает, делит, складывает, вычитает две
//	числовых переменных a,b, значение которых > 2^20.

package tasks

import (
	"fmt"
	"math/big"
)

func Task22_1() {
	x := &big.Float{}
	y := &big.Float{}
	x.SetString("200000000000000000000")
	y.SetString("100000000000000000001")

	answer := new(big.Float)
	fmt.Println("x+y=", answer.Add(x, y))
	fmt.Println("x-y=", answer.Sub(x, y))
	fmt.Println("x*y=", answer.Mul(x, y))
	fmt.Println("x/y=", answer.Quo(x, y))
}
