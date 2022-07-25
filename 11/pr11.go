package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	firstArr := []string{"one", "two", "three", "four", "fiw"}
	secondArr := []string{"three", "four", "four", "one"}

	fmt.Println("", intersection(firstArr, secondArr))
}

func intersection(first, secod []string) []string {

	out := []string{}

	for _, f := range first {
		for _, s := range secod {
			if f == s {

				if !stringInSlice(f, out) {
					out = append(out, f)
				}
			}
		}
	}
	return out
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
