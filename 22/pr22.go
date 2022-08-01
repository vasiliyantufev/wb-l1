package main

/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.
*/

import (
	"fmt"
	"math/big"
)

func main() {

	x := big.NewInt(4000000000000000)
	y := big.NewInt(2000000000000000)

	fmt.Printf("Первая переменная: %v\nВторая пременная: %v\n", x, y)

	//производим операции с большими числами с помощью библиотеки math/big
	addition(x, y)
	multiplication(x, y)
	subtraction(x, y)
	division(x, y)

}

func addition(x, y *big.Int) {
	fmt.Printf("Сложения: %v\n", big.NewInt(0).Add(x, y))
}

func multiplication(x, y *big.Int) {
	fmt.Printf("Умножение: %v\n", big.NewInt(0).Mul(x, y))
}

func subtraction(x, y *big.Int) {
	fmt.Printf("Вычитание: %v\n", big.NewInt(0).Sub(x, y))
}

func division(x, y *big.Int) {
	fmt.Printf("Деление: %v\n", big.NewInt(0).Quo(x, y))
}
