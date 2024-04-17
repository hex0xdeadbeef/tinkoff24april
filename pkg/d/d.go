package main

import "fmt"

const (
	vectorsFmt = "%d %d %d %d\n"
)

func main() {
	var (
		n         int
		direction rune
	)

	fmt.Scanf("%d %c", &n, &direction)

	solutionD(n, byte(direction))
}

func solutionD(size int, mod byte) {
	const (
		even, oddReminder = 2, 1

		left, right byte = 'L', 'R'

		reverseLines, revereColumns = true, false
	)

	switch size % even {
	case oddReminder:
		fmt.Println(square(size))
	default:
		fmt.Println(square(size) + size/even)
	}

	switch mod {
	case left:
		reverseMatrixLinesOrColumns(size, reverseLines)
		transpose(size)

	case right:
		reverseMatrixLinesOrColumns(size, revereColumns)
		transpose(size)
	}
}

// reverseMatrixLinesOrColumns reverses lines' or columnns' elems
// put mod == true if you want to reverse lines, otherwise put false
func reverseMatrixLinesOrColumns(size int, mod bool) {
	const (
		pad         = 1
		sizeDivider = 2
	)

	var (
		pivotIndex = size / sizeDivider
	)

	for i := 0; i < size; i++ {
		for j := 0; j < pivotIndex; j++ {
			switch mod {
			case true:
				fmt.Printf(vectorsFmt, i, j, i, size-j-pad)

			case false:
				fmt.Printf(vectorsFmt, j, i, size-j-pad, i)

			}
		}
	}
}

func transpose(size int) {
	var (
		layerPad int
	)

	for i := 0; i < size; i++ {
		for j := layerPad; j < size; j++ {
			fmt.Printf(vectorsFmt, i, j, j, i)
		}

		layerPad++
	}

}

func square(a int) int {
	return a * a
}
