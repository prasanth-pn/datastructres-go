package main

import  "fmt"




type node struct{
	data int 
	left *node
	right *node

}
type tree struct{
	root *node
}
//insert tree
func (list *tree) insert(val int){
	newnode:=new(node)
	newnode.data =val
	// newnode.left=nil
	// newnode.right=nil
	if list.root==nil{
		list.root=newnode
	}else{
		temp:=list.root
		for{
		if val<temp.data{
			if temp.left==nil{
				temp.left=newnode
				break
			}else{
				temp=temp.left
			}

	}else{
		if val>temp.data{
			if temp.right==nil{
				temp.left.right=newnode
				break
			}else{
				temp=temp.right
			}

		}
	}
	}
	}


}



func main (){
	list:=tree{}
	opt:=0

	for opt<8{
		fmt.Println("----------------------\n enter the choice\n 1 for insert\n 2 for inorder \n 3 for preorder \n 4 for post order \n 5 for search \n 6 for delete \n-----------------\n ")
		fmt.Scan(&opt)

		switch opt{
		case 1:
	list.insert(678)
	list.insert(34)
	list.insert(67)
	list.insert(78)
	list.insert(648)
	list.insert(378)
	list.insert(58)
		}
	}
	
	fmt.Println("program sucessfully completed ")

}