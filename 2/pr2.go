package main

import (
    "fmt"
    "math"
)

/*
task #2
Написать программу, которая конкурентно рассчитает значение
квадратов чисел взятых из массива (2,4,6,8,10) и выведет
их квадраты в stdout.
*/

func main() {
    ch := make(chan float64)
	array := []float64{2, 4, 6, 8, 10}
	for _, a := range array {
		go pow(a, ch)
	    fmt.Println(<-ch)
	}
}

func pow(a float64, ch chan float64) {
    ch <- math.Pow(a, 2)
}
