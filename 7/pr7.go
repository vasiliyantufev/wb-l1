package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

type DataMap struct {
	mx   sync.RWMutex
	Data map[string]string
}

//
//Создаём map
func NewDM() *DataMap {
	return &DataMap{
		Data: make(map[string]string),
	}
}

//пишем в map
func (data *DataMap) Put(id string, o string) {
	data.mx.Lock()
	defer data.mx.Unlock()
	data.Data[id] = o
}

//читаем из map
func (data *DataMap) Get(id string) (o string, b bool) {
	data.mx.RLock()
	defer data.mx.RLock()
	o, b = data.Data[id]
	return
}

func main() {

	data := NewDM()
	data.Put("test", "test")
	fmt.Println(data.Get("test"))
}
