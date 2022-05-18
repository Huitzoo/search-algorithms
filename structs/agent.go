package structs

import (
	"math/rand"
	"search/structs/tree"
	"time"
)

type Agent struct {
	ID        int
	signal    string
	movements int
}

type BestMovementAndScore struct {
}

func NewAgent(signal string, id int) *Agent {
	return &Agent{signal: signal, ID: id}
}

func (ag *Agent) firstMovement(board [][]string) {
	rand.Seed(time.Now().UnixNano())
	yrand := rand.Intn(len(board))
	xrand := rand.Intn(len(board[0]))
	board[yrand][xrand] = ag.signal
}

func (ag *Agent) Move(
	board [][]string,
) {
	if ag.movements == 0 && ag.ID == 1 {
		ag.firstMovement(board)
		return
	}

	ptree := tree.NewTree(1)

	ag.createTreeMovements(board, ptree)
}

func (ag *Agent) createTreeMovements(
	board SimpleBoard,
	ptree *tree.Tree,
) {

	for y, rows := range board {
		for x, column := range rows {
			if column == "" {
				ptree.
			}
		}
	}

	createTreeMovements()
}
