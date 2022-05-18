package utils

var AroundBoardCells = [][]int{
	//{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1},
	{1, 0}, {0, 1}, {0, -1}, {-1, 0},
}

const (
	RuneBoardBlocked  rune = '#'
	LinesBoardWide    int  = 1
	SizeSquareForCell int  = 10 + LinesBoardWide
)
