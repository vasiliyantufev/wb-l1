package main

import (
	"fmt"
	"sort"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	x := []int{33, 303, 303, 40, 50}

	// Element to search
	v := 303

	ind := sort.Search(len(x), func(ind int) bool { return x[ind] >= v })
	if ind < len(x) && x[ind] == v {
		fmt.Printf("found %d at ind %d in %v\n", v, ind, x)
	} else {
		fmt.Printf("%d not found in %v\n", v, x)
	}
}
