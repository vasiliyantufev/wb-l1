package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func main() {

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(deleteElement(s, 2))
}

//делаем добавление с первого элемента до эл,который надо удалить и до последнего элемента
func deleteElement(slice []int, i int) []int {
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}
