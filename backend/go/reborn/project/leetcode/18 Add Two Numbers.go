package main

import "fmt"

/*
你有两个用链表代表的整数，其中每个节点包含一个数字。
数字存储按照在原来整数中相反的顺序，使得第一个数字位于链表的开头。
写出一个函数将两个整数相加，用链表形式返回和。

样例 1:
输入: 7->1->6->null, 5->9->2->null
输出: 2->1->9->null
样例解释: 617 + 295 = 912, 912 转换成链表:  2->1->9->null

样例 2:
输入:  3->1->5->null, 5->9->2->null
输出: 8->0->8->null
样例解释: 513 + 295 = 808, 808 转换成链表: 8->0->8->null

挑战
链表逆序存储，例如：(3→1→5)+(5→9→2)=9→0→7
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resPre := &ListNode{}
	cur := resPre
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}

func main() {
	listNode1 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 6,
			},
		},
	}

	listNode2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 2,
			},
		},
	}

	fmt.Printf("7->1->6->null, 5->9->2->null result: %+v\n", addTwoNumbers(listNode1, listNode2))
}
