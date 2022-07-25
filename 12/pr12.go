package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {

	str := []string{"cat", "cat", "dog", "cat", "tree"}

	mapArray := make(map[string]struct{})

	for _, key := range str {
		mapArray[key] = struct{}{}
	}
	fmt.Println(mapArray)
}
