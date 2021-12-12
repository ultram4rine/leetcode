package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	l1.Display()

	var l2 = removeNthFromEnd(l1, 5)
	l2.Display()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		res = head
		m   = make(map[int]*ListNode)
	)

	var i = 0
	for res != nil {
		i++
		m[i] = res

		res = res.Next
	}

	if i == n {
		head = head.Next
	} else {
		m[i-n+1] = nil
		m[i-n].Next = m[i-n+2]
	}

	return head
}

func (l *ListNode) Display() {
	list := l
	for list != nil {
		fmt.Printf("%+v -> ", list.Val)
		list = list.Next
	}
	fmt.Println()
}
