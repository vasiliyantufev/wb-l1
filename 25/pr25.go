package main

/*
Реализовать собственную функцию sleep.
*/

import (
    "fmt"
    "time"
)

// time.After() возвращает канал. И значение будет отправлено на канал после указанной продолжительности.
// получение будет заблокировано до тех пор, пока значение не будет отправлено:
func Sleep(x int) {
      <-time.After(time.Second * time.Duration(x))
}

func main() {
    fmt.Println("Start")
    Sleep(2)
    fmt.Println("Finish")
}
