package main

type Cocktail struct {
	Name         string
	Instructions []Instruction
}

type Instruction struct {
	Cl   int
	Text string
}
