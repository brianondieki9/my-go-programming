package main

//define type to hold binary tree node
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

//define type to hold all the nodes in the binary tree
type BinaryTree struct {
	root *BinaryNode
}

//function to insert data into the binary tree.
func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}

	return t
}

func (n *BinaryNode) insert(data int64) {
}

func main() {

}
