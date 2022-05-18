package algorithms

import (
	"fmt"
	"math"
	"search/desing"
	"search/structs"
	"search/utils"
)

type scoreAndStep struct {
	Score float64
	Step  float64
}

type AStarSearch struct {
	Size          []int   `yaml:"size"`
	BlockedPoints [][]int `yaml:"blocked_points"`
	InitialState  []int   `yaml:"initial_state"`
	Goal          []int   `yaml:"goal"`
	Kind          string  `yaml:"kind"`
	Nodes         map[int]*structs.Node
	BlockedIds    []int
	stackNodes    map[int]scoreAndStep
	steps         [][]int
	exploredIDs   map[int]struct{}
}

func NewAStarSearch(
	size []int,
	initialState []int,
	goalState []int,
	blockedNodes [][]int,
) StageInterface {
	return &AStarSearch{
		Size:          size,
		InitialState:  initialState,
		Goal:          goalState,
		BlockedPoints: blockedNodes,
	}
}

func (board *AStarSearch) BuildStage() {
	columns := board.Size[0]
	blockedNodes := map[int]struct{}{}
	board.Nodes = make(map[int]*structs.Node)
	board.stackNodes = map[int]scoreAndStep{}

	for _, blockedPoint := range board.BlockedPoints {
		blockedNodes[utils.CalculateIDStateByCoords(blockedPoint, columns)] = struct{}{}
	}

	goalID := utils.CalculateIDStateByCoords(board.Goal, columns)

	board.buildNodes(blockedNodes, goalID, board.Goal)

	board.steps = make([][]int, 0)
	board.exploredIDs = make(map[int]struct{})
	startID := utils.CalculateIDStateByCoords(board.InitialState, columns)
	board.aStarSearchAlgorithm(
		startID, goalID,
	)

	desing.PaintBoard(
		board.Size,
		goalID,
		startID,
		board.exploredIDs,
		blockedNodes,
	)

}

func (board *AStarSearch) buildNodes(
	blockedNodes map[int]struct{},
	goalID int,
	goalCoords []int,
) {
	columns := board.Size[0]
	rows := board.Size[1]

	for x := 0; x < board.Size[0]; x++ {
		for y := 0; y < board.Size[0]; y++ {
			currentID := utils.CalculateIDStateByCoords([]int{x, y}, columns)
			edges := make([]int, 0)
			isGoalID := false
			if _, exist := blockedNodes[currentID]; !exist {
				for _, aroundCells := range utils.AroundBoardCells {
					newX, newY := x+aroundCells[0], y+aroundCells[1]

					arroundID := utils.CalculateIDStateByCoords(
						[]int{newX, newY}, columns,
					)

					if _, exist := blockedNodes[arroundID]; exist ||
						newX == -1 ||
						newX == columns ||
						newY == -1 ||
						newY == rows {

						continue
					}
					if currentID == goalID {
						isGoalID = true
					}
					edges = append(edges, arroundID)
				}
				score := utils.CalculateDistance(goalCoords, []int{x, y})

				board.Nodes[currentID] = structs.NewNode(
					currentID, edges, x, y, isGoalID, score,
				)
			}
		}
	}
}

func (board *AStarSearch) aStarSearchAlgorithm(
	currentID int, goalID int,
) bool {
	getGoal := false
	steps := 0.0

	for !getGoal {

		steps++
		node := board.Nodes[currentID]
		board.steps = append(board.steps, []int{node.CoordX, node.CoordY})

		if node.Goal {
			getGoal = true
			break
		}

		max := math.MaxFloat64
		nextNodeID := 0
		_ = nextNodeID
		wasStakedNode := true
		auxSteps := steps

		for idNode, node := range board.stackNodes {
			if node.Score < max {
				nextNodeID = idNode
				max = node.Score + node.Step
				auxSteps = node.Step
			}
		}

		for _, nextNode := range node.Edges {
			if _, exist := board.exploredIDs[nextNode]; !exist && board.Nodes[nextNode].Score+steps < max {
				nextNodeID = nextNode
				max = board.Nodes[nextNode].Score + steps
				wasStakedNode = false
			} else if _, exist := board.exploredIDs[nextNode]; !exist {
				board.stackNodes[nextNode] = scoreAndStep{
					Score: board.Nodes[nextNode].Score,
					Step:  steps,
				}
			}
		}

		fmt.Println(board.Nodes[nextNodeID].CoordX, board.Nodes[nextNodeID].CoordY, steps)

		if wasStakedNode {
			delete(board.stackNodes, nextNodeID)
			steps = auxSteps
		}

		board.exploredIDs[currentID] = struct{}{}
		currentID = nextNodeID

		fmt.Println(board.stackNodes)
	}

	return false
}
