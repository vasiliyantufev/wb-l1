package main

import (
    "math"
    "fmt"
)

/*
task #3
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func main() {
	array := []float64{2, 4, 6, 8, 10}

	var result float64

	ch := make(chan float64)

	for _, a := range array {

        go pow(a, ch)
		result += <- ch
	}
	fmt.Printf("Сумма квадратов = %v\n", result)
}

func pow(a float64, ch chan float64) {
    ch <- math.Pow(a, 2)
}
