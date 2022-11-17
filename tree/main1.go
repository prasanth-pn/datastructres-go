package main

import "fmt"

type node struct {
	data int

	left   *node
	right  *node
	parent *node
}
type binarysearchTree struct {
	root *node
}

// var k int
func (tree *binarysearchTree) insert(k int) {
	if k == 5 {
		return
	}
	var val int
	fmt.Println("enter the value to insert")
	fmt.Scan(&val)
	newnode := new(node)
	newnode.data = val
	// newnode.left=nil
	// newnode.right=nil
	if tree.root == nil {
		tree.root = newnode
		return

	}
	temp := tree.root
	for temp != nil {
		if newnode.data < temp.data {
			if temp.left == nil {
				temp.left = newnode
				newnode.parent = temp
				return
			}
			temp = temp.left

		} else if newnode.data > temp.data {
			if temp.right == nil {
				temp.right = newnode
				newnode.parent = temp
				return
			}
			temp = temp.right
		}

	}

}

func (tree *binarysearchTree) inorder() {
	//fmt.Print("inorder:  ")
	tree.inordersearch(tree.root)
}
func (tree *binarysearchTree) inordersearch(temp *node) {
	if temp != nil {
		tree.inordersearch(temp.left)
		fmt.Println(temp.data)
		tree.inordersearch(temp.right)
	}

}

func (tree *binarysearchTree) preorder() {
	tree.preordersearch(tree.root)
}
func (tree *binarysearchTree) preordersearch(temp *node) {
	if temp != nil {
		fmt.Println(temp.data)
		tree.preordersearch(temp.left)

		tree.preordersearch(temp.right)
	}
}

func main() {
	n := node{}
	fmt.Println(n)
	tree := binarysearchTree{}
	fmt.Println(tree.root)
	var opt int
	for opt < 444 {
		fmt.Println("--------------------\n 1 for adding into tree  \n 2for display \n 3 for preorder display\n ----------------\n ")
		fmt.Scan(&opt)
		switch opt {
		case 1:

			tree.insert(1)

		case 2:
			tree.inorder()
		case 3:
			tree.preorder()

		}

	}

}
