package main

import "fmt"

func multiplySquareMatrices(m [][]int, n [][]int) [][]int {
	result := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			for k := 0; k < len(m[i]); k++ {
				result[i][j] += m[i][k] * n[k][j]
			}
		}
	}

	return result
}

func main() {
	matrix1 := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	matrix2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(multiplySquareMatrices(matrix1, matrix2))
}
