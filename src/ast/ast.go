package ast

import(
	"fmt"
)

// ast (abstract syntax tree)
type node struct{
	operation_type string
	content string
}

type tree struct{
	root *node 
	left *node  
	right *node 
}

func (t *tree) Walk_left(){
	// traverse the left side of the tree
}

func (t *tree) Walk_right(){
	// traverse the right side of the tree
}

func (t *tree) Check_root(){
	// check if the node is the root
}

func (t *tree) Append_node(n *node){
	// append node to the tree
}

func (t *tree) Check_node_type(n *node){
	// check the type of the node with the tree node
}
