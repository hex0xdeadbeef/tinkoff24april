package main

import "fmt"

type empty struct{}

func main() {
	var (
		n    int
		nums []int
	)

	fmt.Scan(&n)

	nums = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(solutionA(n, nums))
}

func solutionA(size int, marks []int) int {
	const (
		greatMark = 5

		weekSize = 7
	)

	if size < weekSize {
		return -1
	}

	var (
		goodMarks = map[int]empty{4: {}, 5: {}}

		l, r int

		badCount, goodCount, greatCurCount int

		greatMaxCount = -1

		greatCheck = func() {
			if goodCount == weekSize {
				greatMaxCount = max(greatMaxCount, greatCurCount)
			}
		}

		subtractPrev = func(ind int) {
			switch _, ok := goodMarks[marks[ind]]; {
			case ok:
				if marks[ind] == greatMark {
					greatCurCount--
				}
				goodCount--

			default:
				badCount--

			}
		}

		incrementNext = func(ind int) {
			switch _, ok := goodMarks[marks[ind]]; {
			case ok:
				if marks[ind] == greatMark {
					greatCurCount++
				}
				goodCount++

			default:
				badCount++

			}
		}
	)

	for r = 0; r < weekSize; r++ {
		incrementNext(r)
	}

	for i := 0; i < size-weekSize; i++ {

		greatCheck()

		subtractPrev(l)
		l++

		incrementNext(r)
		r++

	}

	greatCheck()

	return greatMaxCount

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
