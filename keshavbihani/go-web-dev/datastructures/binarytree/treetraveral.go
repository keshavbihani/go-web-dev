package main

import "fmt"

type node struct{
	val int 
	left *node
	right *node
}

func Inordertraversal(root *node,f func(int)){
	
	if root==nil{
		return 
	}

	f((*root).val) // or f(root.val)
	Inordertraversal(root.left,f)
	Inordertraversal(root.right,f)
}

func main(){
	root := node{
		1,
		nil,
		nil,
	}
	root.left = &node{
		2,
		nil,
		nil,
	}
	root.left.left = &node{
		4,
		nil,
		nil,
	}
	root.right = &node{
		3,
		nil,
		nil,
	}
	arr:= []int{}
	Inordertraversal(&root,func(x int){
		arr = append(arr, x)
	})
	fmt.Println(arr)
}