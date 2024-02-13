package main

import "fmt"

/*
	=== Задача №1 ===

	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской
	структуры Human (аналог наследования).

*/

// Структура Human со свойствами name (тип - строковый) и age (тип - целочисленный)
type Human struct {
	Name string
	Age  uint
}

// Метод структуры Human. Функционал - вывод свойств структуры в форме, описанной в методе.
func (h *Human) Print() {
	fmt.Printf("Hi! My name is %s. I'm %d years old.", h.Name, h.Age)
}

// Структура Acnion - структура, в которую встраиваться структура Human
type Action struct {
	Human
}

func main() {
	// Инициализация переменной action структурой Action
	action := Action{Human: Human{Name: "Ubulllka", Age: 20}}
	// Вызов метода структуры Human от труктуры Action
	action.Print()
}
