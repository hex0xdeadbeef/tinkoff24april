package main

import (
	"fmt"
	"sort"
	"strings"
)

type Node struct {
	Val      string
	Children []*Node
}

type Nodes []*Node

func (a Nodes) Len() int           { return len(a) }
func (a Nodes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Nodes) Less(i, j int) bool { return a[i].Val < a[j].Val }

func main() {
	var (
		n int

		curStr string
		paths  []string
	)

	fmt.Scanln(&n)

	paths = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&curStr)
		paths[i] = curStr
	}

	solutionC(paths)
}

func solutionC(paths []string) {

	const (
		splitter       = "/"
		indentationStr = ""

		initialInd = 0
	)

	var (
		root *Node = &Node{}

		add      func(path []string, ind int, node *Node)
		traverse func(node *Node, indentation int)
	)

	add = func(path []string, ind int, node *Node) {

		var (
			nextNode *Node
		)

		if ind == len(path) {
			return
		}

		for _, node := range node.Children {
			if path[ind] == node.Val {
				nextNode = node
				break
			}
		}

		if nextNode == nil {
			nextNode = &Node{Val: path[ind], Children: []*Node{}}
			node.Children = append(node.Children, nextNode)
		}

		add(path, ind+1, nextNode)
		return
	}

	traverse = func(node *Node, indentation int) {
		fmt.Printf("%*s%s\n", indentation, indentationStr, node.Val)

		sort.Sort(Nodes(node.Children))

		for _, child := range node.Children {
			traverse(child, indentation+2)
		}

		indentation -= 2
	}

	for _, path := range paths {
		parts := strings.Split(path, splitter)

		add(parts, initialInd, root)
	}

	traverse(root.Children[0], 0)

}
