package solutions

import (
	"fmt"
	"math/big"
)

func Example22() {
	// math/big предоставляет инструмент для работы с большими числами
	a, _ := big.NewFloat(0).SetString("11144444444444444441111111111111114444444444451515155")
	b, _ := big.NewFloat(0).SetString("11144444444444444441111111111111114444444444451515155")
	fmt.Println(sum(a, b).String())
	fmt.Println(subtraction(a, b).String())
	fmt.Println(division(a, b).String())
	fmt.Println(multiplication(a, b).String())
}

func sum(a, b *big.Float) *big.Float {
	return big.NewFloat(0).Add(a, b)
}

func subtraction(a, b *big.Float) *big.Float {
	return big.NewFloat(0).Sub(a, b)
}

func division(a, b *big.Float) *big.Float {
	if f, _ := b.Float64(); f == 0 {
		panic("error")
	}
	return big.NewFloat(0).Quo(a, b)
}

func multiplication(a, b *big.Float) *big.Float {
	return big.NewFloat(0).Mul(a, b)
}
