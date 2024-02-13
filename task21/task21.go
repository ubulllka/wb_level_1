package main

import "fmt"

/*
	=== Задача №21 ===

	Реализовать паттерн «адаптер» на любом примере.

*/


type ForkI interface {
	Stick()
}

type Fork struct {}

func (f *Fork) Stick() {
	fmt.Println("This is fooooork")
}

type SpoonI interface {
	DoSpoon()
}

type Spoon struct {}

func (s *Spoon) DoSpoon(){
	fmt.Println("This is spooooon")
}

type SponAlsoFork struct {
	spoon *Spoon
}

func (s *SponAlsoFork) Stick() {
	s.spoon.DoSpoon()
}

type EatFood struct {}

func (e *EatFood) Eating (forkI ForkI) {
	forkI.Stick()
}


func main() {
	eatFood1 := &EatFood{}
	fork := &Fork{}
	eatFood1.Eating(fork)

	eatFood2 := &EatFood{}
	spoon := &Spoon{}
	sponAlsoFork := &SponAlsoFork{spoon: spoon}
	eatFood2.Eating(sponAlsoFork)
}
