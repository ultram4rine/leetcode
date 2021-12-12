package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	l1.Display()

	var l2 = middleNode(l1)
	l2.Display()
}

func middleNode(head *ListNode) *ListNode {
	var (
		res   *ListNode = nil
		nodes []*ListNode
		l     int
	)

	for head != nil {
		l++

		nodes = append(nodes, head)

		head = head.Next
	}

	res = nodes[l/2]

	return res
}

func (l *ListNode) Display() {
	list := l
	for list != nil {
		fmt.Printf("%+v -> ", list.Val)
		list = list.Next
	}
	fmt.Println()
}
