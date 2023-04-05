

## N 皇后问题
N 皇后问题是一个经典的数学和计算机科学问题，它要求在一个 N x N 的棋盘上放置 N 个皇后，使得每个皇后都不能互相攻击（即不在同一行、同一列或同一斜线上）。  
简单来说，就是要在一个棋盘上摆放 N 个棋子（皇后），使得每行、每列和每条对角线上都只有一个棋子。这个问题挑战了人们的思维能力和计算能力，在计算机科学中也有很多研究和应用。

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    n := 8 // 设置 N 的值为 8，即 8 皇后问题
    queens := make([]int, n) // 初始化一个长度为 n 的整数切片，用来存放每行皇后的列号
    solveNQueens(queens, 0, n) // 调用 solveNQueens 函数解决 N 皇后问题
}

// 解决 N 皇后问题
// queens：存放每行皇后的列号
// row：当前行号
// n：皇后的数量
func solveNQueens(queens []int, row int, n int) {
    // 如果已经找到了一组解，就输出这组解并返回
    if row == n {
        printQueens(queens)
        return
    }

    // 依次尝试当前行的每一列
    for i := 0; i < n; i++ {
        if isValid(queens, row, i) { // 判断当前位置是否可以放置皇后
            queens[row] = i // 放置皇后
            solveNQueens(queens, row+1, n) // 递归调用下一行
        }
    }
}

// 判断当前位置是否可以放置皇后
// queens：存放每行皇后的列号
// row：当前行号
// col：当前列号
func isValid(queens []int, row int, col int) bool {
    // 遍历前面已经放置的皇后
    for i := 0; i < row; i++ {
        // 判断当前位置是否和前面已经放置的皇后在同一列或同一对角线上
        if queens[i] == col || math.Abs(float64(queens[i]-col)) == float64(row-i) {
            return false // 如果是，则返回 false，表示这个位置不能放置皇后
        }
    }
    return true // 否则返回 true，表示这个位置可以放置皇后
}

// 输出一组解
// queens：存放每行皇后的列号
func printQueens(queens []int) {
    n := len(queens)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if queens[i] == j { // 如果当前位置放置了皇后，就输出 Q
                fmt.Print("Q ")
            } else { // 否则输出 .
                fmt.Print(". ")
            }
        }
        fmt.Println() // 换行
    }
    fmt.Println() // 输出一个空行，方便查看
}
```
