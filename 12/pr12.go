package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {

	str := []string{"cat", "cat", "dog", "cat", "tree"}

	mapSet := make(map[string]struct{})

	for _, key := range str {
		// создает собственное множество для str
		mapSet[key] = struct{}{}
	}
	fmt.Println(mapSet)
}
