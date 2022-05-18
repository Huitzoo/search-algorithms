package structs

import "fmt"

type Node struct {
	ID     int
	CoordX int
	CoordY int
	Goal   bool
	Score  float64
	Edges  []int
}

type Edge struct {
	NextNodes []Node
}

var QueueStates *Queue
var steps int = 0
var score int = 0

func init() {
	QueueStates = NewQueue()
}

func NewNode(
	id int,
	edges []int,
	CoordX int,
	CoordY int,
	Goal bool,
	Score float64,
) *Node {
	return &Node{
		id, CoordX, CoordY, Goal, Score, edges,
	}
}

func (n *Node) ExploreNode() {
	if n.Goal {
		fmt.Sprintf("You walked with: %d and you get: %d in score", steps, score)
	} else {
		steps++
	}
}

func GetNextNode() *Queue {
	node := QueueStates.PopLIFO()

	return node.(*Queue)
}
