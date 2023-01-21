package main

import "fmt"

/*
给一棵二叉搜索树以及一个整数 n, 如果在树中找到和为 n 的两个数字则返回true

样例:
输入:
n = 3

二叉搜索树如下：
    4
   / \
  2   5
 / \
1   3
输出:
true

提示:
使用广度优先搜索(BFS)遍历整个树
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	return helper(root, root, k)
}

func helper(root, searchRoot *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	return (root.Val*2 != k && findNode(searchRoot, k-root.Val)) ||
		helper(root.Left, searchRoot, k) ||
		helper(root.Right, searchRoot, k)
}

func findNode(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	}

	if root.Val < target {
		return findNode(root.Right, target)
	}

	return findNode(root.Left, target)
}

func main() {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	target := 3
	fmt.Printf("return: %t\n", findTarget(tree, target))
}
