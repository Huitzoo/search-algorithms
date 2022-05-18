package algorithms

/*
import (
	"search/desing"
	"search/structs"
	"search/utils"
)

type StackFrontier struct {
	Size          []int   `yaml:"size"`
	BlockedPoints [][]int `yaml:"blocked_points"`
	InitialState  []int   `yaml:"initial_state"`
	Goal          []int   `yaml:"goal"`
	Kind          string  `yaml:"kind"`
	Nodes         []structs.Node
	BlockedIds    []int
	frontier      *structs.Queue
	steps         [][]int
	exploredIDs   map[int]struct{}
}

var QueueStates *structs.Queue
var steps [][]int
var exploredIDs map[int]struct{}

func init() {
	QueueStates = structs.NewQueue()
	steps = make([][]int, 0)
	exploredIDs = make(map[int]struct{})
}

func (board *StackFrontier) BuildStage() {
	columns := board.Size[0]
	blockedPosition := map[int]struct{}{}

	for _, blockedPoint := range board.BlockedPoints {
		blockedPosition[utils.CalculateIDStateByCoords(blockedPoint, columns)] = struct{}{}
	}

	goalID := utils.CalculateIDStateByCoords(board.Goal, columns)

	calculateNodesArround(
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
		exploredIDs,
		blockedPosition,
	)

}

func (board *StackFrontier) calculateNodesArround(
	coords []int,
	blockedPosition map[int]struct{},
	maxMin []int,
	goalID int,
) bool {

	actualID := utils.CalculateIDStateByCoords(
		coords, maxMin[0],
	)

	for _, aroundBoardCell := range utils.AroundBoardCells {
		newX, newY := coords[0]+aroundBoardCell[0], coords[1]+aroundBoardCell[1]

		nextID := utils.CalculateIDStateByCoords(
			[]int{newX, newY}, maxMin[0],
		)

		if _, exist := blockedPosition[nextID]; exist || newX == -1 || newX == maxMin[0] || newY == -1 || newY == maxMin[1] {
			continue
		}

		if nextID == goalID {
			exploredIDs[actualID] = struct{}{}
			return true
		}

		if _, exist := exploredIDs[nextID]; !exist {
			QueueStates.Push(
				[]int{newX, newY},
			)
		}
	}

	exploredIDs[actualID] = struct{}{}
	if QueueStates.IsEmpty() {
		return false
	} else {
		next := QueueStates.PopLIFO().([]int)
		steps = append(steps, next)
		return calculateNodesArround(
			next,
			blockedPosition,
			maxMin,
			goalID,
		)
	}
}
*/
