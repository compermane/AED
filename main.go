package main

import (
	"fmt"

	"github.com/compermane/AED/AVL"
)

func main() {
	root := AVL.CreateNode(3)
	fmt.Println(root.Inorder())
}
