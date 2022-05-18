package tree

type Tree struct {
	FatherNode *Node
	depth      int
}

func NewTree(depth int, board *[][]string) *Tree {

	node := NewNode(0, board)
	return &Tree{depth: depth, FatherNode: node}
}

func (t *Tree) InsertNode(board *[][]string) *Tree {

}
