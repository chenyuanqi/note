package main

import "fmt"

/*
n皇后问题是将n个皇后放置在n×n的棋盘上，皇后彼此之间不能相互攻击(任意两个皇后不能位于同一行，同一列，同一斜线)。

给定一个整数n，返回所有不同的N皇后问题的解决方案。

每个解决方案包含一个明确的N皇后放置布局，其中 “Q” 和 “X” 分别表示一个女王和一个空位置。

样例
例1:
输入:1
输出:
   [["Q"]]

例2:
输入:4
输出:
[
  ["XQXX",
   "XXXQ",
   "QXXX",
   "XXQX"
  ],
  ["XXQX",
   "QXXX",
   "XXXQ",
   "XQXX"
  ]
]

挑战
你能否不使用递归完成？
*/
func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}

	cols := make([]bool, n)
	d1 := make([]bool, 2*n)
	d2 := make([]bool, 2*n)

	board := make([]string, n)

	res := [][]string{}

	dfs(0, cols, d1, d2, board, &res)

	return res
}

func dfs(r int, cols, d1, d2 []bool, board []string, res *[][]string) {

	if r == len(board) {
		tmp := make([]string, len(board))
		copy(tmp, board)
		*res = append(*res, tmp)
		return
	}

	n := len(board)

	for c := 0; c < len(board); c++ {
		id1 := r - c + n
		id2 := 2*n - r - c - 1
		if !cols[c] && !d1[id1] && !d2[id2] {
			b := make([]byte, n)
			for i := range b {
				b[i] = 'x'
			}
			b[c] = 'Q'
			board[r] = string(b)
			cols[c], d1[id1], d2[id2] = true, true, true

			dfs(r+1, cols, d1, d2, board, res)

			cols[c], d1[id1], d2[id2] = false, false, false
		}
	}
}

func main() {
	fmt.Printf("1 queens: %+v\n", solveNQueens(1))
	fmt.Printf("4 queens: %+v\n", solveNQueens(4))
}
