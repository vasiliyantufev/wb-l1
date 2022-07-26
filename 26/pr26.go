package main

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

import "fmt"

func main() {
    fmt.Println(unionString("abcd"))
    fmt.Println(unionString("abCdefAaf"))
    fmt.Println(unionString("aabcd"))
}

func unionString(str string) (string, bool) {

    for i := 0; i < len(str); i++ {
        for j := i + 1; j < len(str); j++ {
            if(str[i] == str[j]) {
                return str, false
            }
        }
    }
    return str, true
}