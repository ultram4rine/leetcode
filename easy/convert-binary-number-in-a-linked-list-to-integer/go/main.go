package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

func getDecimalValue(head *ListNode) int {
	var res = head.Val

	for head.Next != nil {
		res = (res << 1) | head.Next.Val
		head = head.Next
	}

	return res
}
