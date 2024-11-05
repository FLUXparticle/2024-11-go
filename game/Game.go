package main

import "fmt"

type Game struct {
	Template
	curPlayer int
}

func (g *Game) Play() {
	g.curPlayer = 1
	for !g.IsGameOver() {
		g.PrintState()
		switch g.curPlayer {
		case 1:
			g.MovePlayer1()
		case 2:
			g.MovePlayer2()
		}
		g.curPlayer = 3 - g.curPlayer
	}
	win := g.Winner()
	fmt.Printf("Player %d wins.\n", win)
}
