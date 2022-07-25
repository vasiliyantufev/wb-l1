package main

import (
	"fmt"
	"sort"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {

	array := []int{-12, 4, 9, -9, 11, -8}
	sort.Ints(array)
	fmt.Println(array)
}
