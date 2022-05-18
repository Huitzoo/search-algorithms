package main

import (
	"fmt"
	"search/structs"
	"search/structs/algorithms"
)

func main() {

	agent1 := structs.NewAgent("x", 1)
	agent2 := structs.NewAgent("y", 2)
	board := algorithms.NewTicTacToe()

	agent1.Move(board.GetBoard())

	agent2.Move(board.GetBoard())
	fmt.Println(board.GetBoard())
}
