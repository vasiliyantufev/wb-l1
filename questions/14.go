package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) { // новая локальная переменная
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
