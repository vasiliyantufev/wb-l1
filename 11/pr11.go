package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {

	//два множества
	firstSet := []string{"one", "two", "three", "four", "fiw"}
	secondSet := []string{"three", "four", "four", "one"}

	fmt.Println("", intersection(firstSet, secondSet))
}

func intersection(first, secod []string) []string {

	out := []string{}

	for _, f := range first {
		for _, s := range secod {

			// если значение первого и второго множества одинаковы и в результирующем множестве текущего значения нет,
			// то записываем
			if f == s && !stringInSlice(f, out) {
				out = append(out, f)
			}
		}
	}
	return out
}

//поиск элемента в результирующем множестве
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
