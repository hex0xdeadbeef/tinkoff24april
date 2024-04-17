package main

import "fmt"

func main() {
	var (
		n, m   int
		matrix [][]uint64
	)

	fmt.Scanf("%d %d", &n, &m)

	matrix = make([][]uint64, n)
	for i := 0; i < n; i++ {
		row := make([]uint64, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&row[j])
		}

		matrix[i] = row
	}

	solutionB(matrix)
}

func solutionB(matrix [][]uint64) {
	for i := 0; i < len(matrix[0]); i++ {
		for j := len(matrix) - 1; j >= 0; j-- {
			fmt.Printf("%d ", matrix[j][i])
		}

		fmt.Println()
	}
}
