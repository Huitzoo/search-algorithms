package algorithms

import (
	"search/desing"
	"search/structs"
	"search/utils"
)

type BreadthFirstSearch struct {
	Size          []int   `yaml:"size"`
	BlockedPoints [][]int `yaml:"blocked_points"`
	InitialState  []int   `yaml:"initial_state"`
	Goal          []int   `yaml:"goal"`
	BlockedIds    []int
	frontier      *structs.Queue
	steps         [][]int
	exploredIDs   map[int]struct{}
}

func NewBreadthFirstSearch(
	size []int,
	initialState []int,
	goalState []int,
	blockedNodes [][]int,
) StageInterface {
	return &BreadthFirstSearch{
		Size:          size,
		InitialState:  initialState,
		Goal:          goalState,
		BlockedPoints: blockedNodes,
	}
}

func (board *BreadthFirstSearch) BuildStage() {
	columns := board.Size[0]
	blockedPosition := map[int]struct{}{}

	for _, blockedPoint := range board.BlockedPoints {
		blockedPosition[utils.CalculateIDStateByCoords(blockedPoint, columns)] = struct{}{}
	}

	board.frontier = structs.NewQueue()
	board.steps = make([][]int, 0)
	board.exploredIDs = make(map[int]struct{})

	goalID := utils.CalculateIDStateByCoords(board.Goal, columns)

	board.breadthFirstSearchAlgoritm(
		board.InitialState,
		blockedPosition,
		board.Size,
		goalID,
	)

	startID := utils.CalculateIDStateByCoords(board.InitialState, columns)

	desing.PaintBoard(
		board.Size,
		goalID,
		startID,
		board.exploredIDs,
		blockedPosition,
	)

}

func (board *BreadthFirstSearch) breadthFirstSearchAlgoritm(
	coords []int,
	blockedPosition map[int]struct{},
	maxMin []int,
	goalID int,
) bool {
	actualID := utils.CalculateIDStateByCoords(
		coords, maxMin[0],
	)

	for _, aroundBreadthFirstSearchCell := range utils.AroundBoardCells {
		newX, newY := coords[0]+aroundBreadthFirstSearchCell[0], coords[1]+aroundBreadthFirstSearchCell[1]

		nextID := utils.CalculateIDStateByCoords(
			[]int{newX, newY}, maxMin[0],
		)

		if _, exist := blockedPosition[nextID]; exist || newX == -1 || newX == maxMin[0] || newY == -1 || newY == maxMin[1] {
			continue
		}

		if nextID == goalID {
			board.exploredIDs[actualID] = struct{}{}
			return true
		}

		if _, exist := board.exploredIDs[nextID]; !exist {
			board.frontier.Push(
				[]int{newX, newY},
			)
		}
	}

	board.exploredIDs[actualID] = struct{}{}
	if board.frontier.IsEmpty() {
		return false
	} else {
		next := board.frontier.PopFIFO().([]int)
		board.steps = append(board.steps, next)
		return board.breadthFirstSearchAlgoritm(
			next,
			blockedPosition,
			maxMin,
			goalID,
		)
	}
}
