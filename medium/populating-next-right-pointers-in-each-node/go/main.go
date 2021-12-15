package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	var root = &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7},
		},
	}
	connect(root)

	fmt.Println(root.Left.Next.Val, root.Left.Left.Next.Val, root.Left.Right.Next.Val, root.Right.Left.Next.Val)
}

func connect(root *Node) *Node {
	if root.Left == nil && root.Right == nil {
		return root
	}
	connectLevel([]*Node{root.Left, root.Right})
	return root
}

func connectLevel(nodes []*Node) {
	var nextNodes = make([]*Node, 0, len(nodes)*2)
	for i := range nodes {
		if i != len(nodes)-1 {
			nodes[i].Next = nodes[i+1]
		}
		if nodes[i].Left != nil {
			nextNodes = append(nextNodes, nodes[i].Left, nodes[i].Right)
		}
	}
	if len(nextNodes) == 0 {
		return
	}
	connectLevel(nextNodes)
}
