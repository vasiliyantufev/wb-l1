package main

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// Target предоставляет интерфейс, с которым должна работать система.
type Target interface {
	Request() string
}

// Adaptee реализует систему, которую нужно адаптировать.
type Adaptee struct {
}

// NewAdapter — конструктор адаптера.
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

// SpecificRequest - реализация конкретного запроса
func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

// Adapter реализует Target интерфейс и является адаптером.
type Adapter struct {
	*Adaptee
}

// Request - адаптивный метод.
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}
