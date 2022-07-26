package main

import (
    "fmt"
    "strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {

    str := "one two three"
    array := strings.Split(str, " ")
    reversArray := make([]string, 0)

    for i := len(array) - 1; i>=0; i-- {
        reversArray = append(reversArray, array[i])
    }

    fmt.Println("Строка: ", str)
    fmt.Println("Строка перевернутая: ", strings.Join(reversArray, " "))
}
