package tree

type Node struct {
	Nodes []*Node
	Score int
	Board *[][]string
}

func NewNode(score int, board *[][]string) *Node {
	return &Node{Score: score, Board: board}
}
