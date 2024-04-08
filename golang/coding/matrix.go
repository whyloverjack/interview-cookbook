package main

import (
	"fmt"
)

// 正向螺旋矩阵
func spiralMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	top, bottom, left, right := 0, n-1, 0, n-1
	direction := 0 // 0 - right, 1 - down, 2 - left, 3 - up

	for i := 1; i <= n*n; i++ {
		switch direction {
		case 0:
			if left <= right {
				matrix[top][left] = i
				left++
			}
		case 1:
			if top <= bottom {
				matrix[bottom][left] = i
				bottom--
			}
		case 2:
			if right >= left {
				matrix[bottom][right] = i
				right--
			}
		case 3:
			if bottom >= top {
				matrix[top][right] = i
				top++
			}
		}

		// 更新边界并改变移动方向
		switch {
		case left > right:
			direction = (direction + 1) % 4
			left--
			right++
		case top > bottom:
			direction = (direction + 1) % 4
			top--
			bottom++
		default:
			if direction == 0 {
				left++
			} else if direction == 1 {
				bottom--
			} else if direction == 2 {
				right--
			} else {
				top++
			}
		}
	}

	return matrix
}

// 反向螺旋矩阵
func reverseSpiralMatrix(n int) [][]int {
	spiral := spiralMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			spiral[i][j], spiral[n-j-1][i], spiral[n-i-1][n-j-1], spiral[j][n-i-1] =
				spiral[n-j-1][i], spiral[n-i-1][n-j-1], spiral[j][n-i-1], spiral[i][j]
		}
	}
	return spiral
}

func main() {
	n := 5
	fmt.Println("Forward Spiral Matrix:")
	for _, row := range spiralMatrix(n) {
		fmt.Println(row)
	}

	fmt.Println("\nReverse Spiral Matrix:")
	for _, row := range reverseSpiralMatrix(n) {
		fmt.Println(row)
	}
}
