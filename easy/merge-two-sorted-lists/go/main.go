package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var res *ListNode = new(ListNode)

	if list1.Val < list2.Val {
		res = list1
		res.Next = mergeTwoLists(list1.Next, list2)
	} else {
		res = list2
		res.Next = mergeTwoLists(list1, list2.Next)
	}

	return res
}
