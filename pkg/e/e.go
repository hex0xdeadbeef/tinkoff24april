package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		n int

		curStr string
		paths  [][]string
	)

	fmt.Scanln(&n)

	paths = make([][]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&curStr)
		paths[i] = strings.Split(curStr, "")
	}

	solutionE(paths)
}

func solutionE(rows [][]string) {
	const (
		columnsNumber       = 3
		left, middle, right = 0, 1, 2
		pad                 = 1

		bush, mushroom, grass          = "W", "C", "."
		bushVal, mushroomVal, grassVal = -1e5, 1, 0
	)

	var (
		maxCount int

		prevRow = []int{grassVal, grassVal, grassVal}
		curRow  = make([]int, columnsNumber)

		fillNextRow = func(nums []int, obstacles []string) {
			for i, obstacle := range obstacles {
				switch obstacle {
				case mushroom:
					nums[i] = mushroomVal
				case grass:
					nums[i] = grassVal
				case bush:
					nums[i] = bushVal
				}
			}
		}

		updateRow = func(prev, cur []int) bool {

			var (
				bushesCount int
			)
			for i := 0; i < columnsNumber; i++ {

				if cur[i] == bushVal {
					bushesCount++
					if bushesCount == columnsNumber {
						return false
					}
					continue
				}

				switch i {
				case left:
					cur[i] = max(max(cur[i], cur[i]+prev[i]), max(cur[i], cur[i]+prev[i+pad]))
				case middle:
					cur[i] = max(
						max(max(cur[i], cur[i]+prev[i]), max(cur[i], cur[i]+prev[i+pad])),
						max(max(cur[i], cur[i]+prev[i-pad]), max(cur[i], cur[i]+prev[i])),
					)
				case right:
					cur[i] = max(max(cur[i], cur[i]+prev[i-pad]), max(cur[i], cur[i]+prev[i]))
				}

				maxCount = max(maxCount, cur[i])

			}

			return true
		}
	)

	// DP
	for _, nextRow := range rows {
		fillNextRow(curRow, nextRow)
		if !updateRow(prevRow, curRow) {
			break
		}

		prevRow, curRow = curRow, prevRow
		curRow = make([]int, columnsNumber)
	}

	fmt.Print(maxCount)

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
