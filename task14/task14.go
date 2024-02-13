package main

import "fmt"

/*
	=== Задача №14 ===

	Разработать программу, которая в рантайме способна определить
	тип переменной: int, string, bool, channel из переменной типа interface{}.

*/

func getType(in interface{}) {
	switch v := in.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	case chan int:
		fmt.Println(v, "is chan int")
	case chan string:
		fmt.Println(v, "is chan string")
	case chan bool:
		fmt.Println(v, "is chan bool")
	default:
		fmt.Println("I don't know what is it")
	}
}

func main() {
	var aint int
	var bstring = "sss"
	var cbool bool
	chanInt := make(chan int)
	chanString := make(chan string)
	chanBool := make(chan bool)

	getType(aint)
	getType(bstring)
	getType(cbool)
	getType(chanInt)
	getType(chanString)
	getType(chanBool)

}
