package main

import (
	"fmt"
	"strings"
)

const DEBUG bool = true

type Node struct {
	char  string
	left  *Node
	right *Node
}

func createNode(char string) *Node {
	return &Node{
		char:  char,
		left:  nil,
		right: nil,
	}
}

// Preorder traversal: Root -> Left -> Right
func preorder(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%s ", node.char)
	preorder(node.left)
	preorder(node.right)
}


func decodeHuffman(encodedText string, codingTable map[string]string) string {
	var decoded strings.Builder
	box := ""
	for i := 0 ; i < len(encodedText) ; i++{
		box += 

	}
	return decoded.String()
}



