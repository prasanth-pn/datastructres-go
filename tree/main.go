package main

import "fmt"

type node struct{
	data int
	left,right *node
}
type bst struct{
	root *node
}
func (tree *bst) insert(data int){
	// var data int 
	// fmt.Println("enter the values to insert")
	// fmt.Scan(&data)
	newnode:=new(node)
	newnode.data=data
	if tree.root==nil{
tree.root=newnode
	}else{
		temp:=tree.root
		for {
			if data >temp.data{
				if temp.left==nil{
		           temp.left=newnode
				   break

			}else{
				temp=temp.left
			}
		}else{
			if data<temp.data{
				if temp.right==nil{
					temp.right=newnode
					break
				}else{
					temp=temp.right

				}
			}
		}
	}
	}

}
func (tree *bst)display(){
	var opt int=0
	for opt<6{
		fmt.Println("--------------\n enter the option \n 1for insert\n 2 for post order\n 3 for preorder \n-4 fordelete \n 5  for inorder\n -------------")
		fmt.Scan(&opt)
		switch opt {
		case 5:
			fmt.Println("DISPLAYING IN INORDER")
			inorder(tree.root)
		case  2:
			postorder(tree.root)
		case 3:
			preorder(tree.root)
		case 4:
			delete(tree)
		case 1:
			
			tree.insert(435)
			tree.insert(145)
			tree.insert(457)
			tree.insert(45)
			tree.insert(445)
			tree.insert(4325)
			tree.insert(2345)
			tree.insert(1145)
			tree.insert(4115)
			tree.insert(1245)

			

			


			
		}
	}

}
func inorder( root *node){
	
	if root!=nil{
		inorder(root.right)
		fmt.Println(root.data)
		inorder(root.left)
		
	}

}
func preorder(root *node){
	if root!=nil{
		fmt.Println(root.data)
		postorder(root.left)
		postorder(root.right)
	}
}
func postorder(root *node){
	if root!=nil{
		postorder(root.left)
		postorder(root.right)
		fmt.Println(root.data)
		
	}
}
//delete the values
func delete(tree *bst){
	fmt.Println("enter the value to delete")
	var num int
	fmt.Scan(&num)
	tree.remove(num, tree.root,nil)
	}

	func (tree *bst) remove(data int,currentnode *node,parentnode *node){
for currentnode!=nil{
	if data<currentnode.data{
		parentnode=currentnode
		currentnode=currentnode.left  
	}else if data>currentnode.data{
		parentnode=currentnode
		currentnode=currentnode.right
		//fmt.Println("something wrong------------------")

	}else{
		if currentnode.left!=nil&&currentnode.right!=nil{
			currentnode.data=getminimum(currentnode.right)
			tree.remove(currentnode.data,currentnode.right,currentnode)
		}else{
			if parentnode==nil{
				if currentnode.left==nil{
					tree.root=currentnode.right
				}else{
					tree.root=currentnode.left
				}
			}else{
				if parentnode.left==currentnode{
					if parentnode.right==nil{
						parentnode.left=currentnode.left
					}else{
						parentnode.left=currentnode.right
					}
				}else {
					if parentnode.right==nil{
						parentnode.right=currentnode.left
					}else{
						parentnode.right=currentnode.right
					}
				}
			}
		}
		break
	}
}
	}
	func getminimum(currentnode *node)int{
		if currentnode.left==nil{
			return currentnode.data
		}
		return getminimum(currentnode.left)
	}

func main(){
	tree:=bst{}
	//fmt.Println(tree)
	tree.display()

}