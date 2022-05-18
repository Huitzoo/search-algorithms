package algorithms

type SimpleBoard struct {
	board [][]string
}

func NewSimpleBoard(board [][]string) *SimpleBoard {
	return &SimpleBoard{board: board}
}

func (b *SimpleBoard) GetScoreBoard(
	wonSignal string,
) {
	for _, rows := range b.board {
		for _, columns := range rows {
			if columns == "" {

			}
		}
	}
}

func (t *SimpleBoard) PaintMovement(
	x, y int,
	signal string,
) {
	t.board[x][y] = signal
}

type TicTacToe struct {
	board *SimpleBoard
}

func NewTicTacToe() *TicTacToe {
	board := [][]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}

	return &TicTacToe{board: NewSimpleBoard(board)}
}

func (t *TicTacToe) GetBoard() *SimpleBoard {
	return t.board
}
