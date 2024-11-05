package main

type Template interface {
	IsGameOver() bool
	PrintState()
	MovePlayer1()
	MovePlayer2()
	Winner() int
}
