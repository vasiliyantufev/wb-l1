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

//функция для проверки уникальности символов в строке
func unionString(str string) (string, bool) {

	//каждый символ в строке проверяется на уникальность
	//если находятся одинаковые, выбрасываем false
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return str, false
			}
		}
	}
	return str, true
}
