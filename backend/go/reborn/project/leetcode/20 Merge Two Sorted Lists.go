package main

import "fmt"

/*
将两个排序链表合并为一个新的排序链表

样例 1:
输入:  null,  0->3->3->null
输出: 0->3->3->null

样例2:
输入: 1->3->8->11->15->null, 2->null
输出: 1->2->3->8->11->15->null

样例3:
输入：1->2->4->null, 1->3->4->null
输出：1->1->2->3->4->4->null
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func Merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}

		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	}

	if l2 != nil {
		node.Next = l2
	}

	return head
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	fmt.Printf("1->2->4->null, 1->3->4->null, return: %+v\n", Merge(l1, l2))
}
