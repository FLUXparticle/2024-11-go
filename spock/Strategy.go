package spock

type Strategy interface {
	NextMove(lastMove Move) Move
}
