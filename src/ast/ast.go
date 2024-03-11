package ast

//"fmt"

// ast (abstract syntax tree)
type Node struct {
	operation_type string
	content        string
}

type Tree struct {
	root  *Node
	left  *Node
	right *Node
}

func (t *Tree) Walk_left() {
	// traverse the left side of the tree
}

func (t *Tree) Walk_right() {
	// traverse the right side of the tree
}

func (t *Tree) Check_root() {
	// check if the node is the root
}

func (t *Tree) Append_node(n *Node) {
	// append node to the tree
}

func (t *Tree) Check_node_type(n *Node) {
	// check the type of the node with the tree node
}
