package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

type Node struct {
	Value    *bool
	Operator *string
	Left     *Node
	Right    *Node
}

type Node2 struct {
	Name     string
	Children []*Node2
}

func (n *Node) Eval() bool {
	if n.Value != nil {
		return *n.Value
	}
	switch *n.Operator {
	case "AND":
		return n.Left.Eval() && n.Right.Eval()
	case "OR":
		return n.Left.Eval() || n.Right.Eval()
	}
	return false //TODO: error handling
}

func main() {
	data, _ := os.ReadFile("2.json")
	fmt.Println(string(data))

	tree := Node{}
	json.Unmarshal(data, &tree)
	fmt.Println(tree)
	spew.Dump(tree)
	fmt.Println(tree.Eval())
}
