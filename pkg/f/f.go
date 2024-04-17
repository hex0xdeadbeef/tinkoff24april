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

	solutionF(paths)
}

type Move struct {
	dx, dy int
}

type Coordinate struct {
	x, y          int
	state         string
	curMovesCount int
}

const (
	placeholder   = "."
	start, finish = "S", "F"
	knight, king  = "K", "G"

	pad = 1

	inf = 1e5
)

var (
	knightMoves, kingMoves []Move
)

func init() {
	kingMoves = []Move{
		{dx: 0, dy: 1},
		{dx: 1, dy: 1},
		{dx: 1, dy: 0},
		{dx: 1, dy: -1},
		{dx: 0, dy: -1},
		{dx: -1, dy: -1},
		{dx: -1, dy: 0},
		{dx: -1, dy: 1},
	}

	knightMoves = []Move{
		{dx: 1, dy: 2},
		{dx: -1, dy: 2},

		{dx: 1, dy: -2},
		{dx: -1, dy: -2},

		{dx: 2, dy: -1},
		{dx: -2, dy: -1},

		{dx: 2, dy: 1},
		{dx: -2, dy: 1},
	}

}

func solutionF(rows [][]string) {

	var (
		height, width      = len(rows), len(rows[0])
		initialX, initialY = findCell(rows, start)

		movesCount = createPolygon(height, width)

		queue = make([]Coordinate, 0, height*width)
	)

	queue = append(queue, Coordinate{x: initialX, y: initialY, state: knight, curMovesCount: 0})

	var (
		top Coordinate
	)

	for len(queue) != 0 {
		top = queue[0]
		queue = queue[1:]

		if top.curMovesCount >= movesCount[top.x][top.y] {
			continue
		} else {
			movesCount[top.x][top.y] = top.curMovesCount
		}

		top.curMovesCount++

		if deskCoordinateState := rows[top.x][top.y]; deskCoordinateState != placeholder && deskCoordinateState != start {
			top.state = deskCoordinateState
		}

		var (
			moveOptions []Move

			isWithinBounds = func(x, dx, y, dy int) bool {
				return (x+dx < height && y+dy < width) && (x+dx >= 0 && y+dy >= 0)
			}

			newCoordinate Coordinate
		)

		switch top.state {
		case knight:
			moveOptions = knightMoves
		case king:
			moveOptions = kingMoves
		}

		for _, move := range moveOptions {
			if isWithinBounds(top.x, move.dx, top.y, move.dy) {
				newCoordinate = Coordinate{x: top.x + move.dx, y: top.y + move.dy, state: top.state, curMovesCount: top.curMovesCount}
				queue = append(queue, newCoordinate)
			}

		}
	}

	var (
		finishX, finishY = findCell(rows, finish)
	)

	if finishVal := movesCount[finishX][finishY]; finishVal != inf {
		fmt.Println(finishVal)
	} else {
		fmt.Println(-1)
	}
}

func findCell(rows [][]string, state string) (int, int) {
	var (
		height, width = len(rows), len(rows[0])
	)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if rows[i][j] == state {
				return i, j
			}
		}
	}

	panic("UNREACHABLE")
}

func createPolygon(heigth, width int) [][]int {
	var (
		polygon [][]int = make([][]int, heigth)
	)

	for i := 0; i < heigth; i++ {
		polygon[i] = make([]int, width)
		for j := 0; j < width; j++ {
			polygon[i][j] = inf
		}
	}

	return polygon
}
