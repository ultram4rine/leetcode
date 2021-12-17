package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var tail = reverseList(head.Next)
	head.Next.Next = head

	head.Next = nil

	return tail
}
