package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	/* var l1 = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	var l2 = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}} */
	/* var l1 = &ListNode{Val: 0}
	var l2 = &ListNode{Val: 0} */
	/* var l1 = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	var l2 = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}} */
	var l1 = &ListNode{Val: 5, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5}}}
	var l2 = &ListNode{Val: 5, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5}}}

	l1.Display()
	l2.Display()

	var l3 = addTwoNumbers(l1, l2)
	l3.Display()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		res   *ListNode = nil // result.
		i1              = l1
		i2              = l2
		shift           = false
	)

	for {
		var (
			cur1 int
			cur2 int
		)

		if i1 == nil && i2 == nil {
			break
		} else if i1 != nil && i2 == nil {
			cur1 = i1.Val
			i1 = i1.Next
		} else if i1 == nil && i2 != nil {
			cur2 = i2.Val
			i2 = i2.Next
		} else {
			cur1 = i1.Val
			cur2 = i2.Val
			i1 = i1.Next
			i2 = i2.Next
		}

		var sum = cur1 + cur2
		if shift {
			sum += 1
		}
		if sum >= 10 {
			shift = true
		} else {
			shift = false
		}

		if res == nil {
			res = &ListNode{Val: sum % 10}
		} else {
			var last = res
			for last.Next != nil {
				last = last.Next
			}
			last.Next = &ListNode{Val: sum % 10}
		}
	}

	if shift {
		var last = res
		for last.Next != nil {
			last = last.Next
		}
		last.Next = &ListNode{Val: 1}
	}

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
