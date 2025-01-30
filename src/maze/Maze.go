package maze

import (
	"fmt"
	"math/rand"
)

func NewMaze(width, height int) *Maze {
	maze := &Maze{
		Grid:   make([][]int, height),
		Width:  width,
		Height: height,
	}

	for i := range maze.Grid {
		maze.Grid[i] = make([]int, width)
		for j := range maze.Grid[i] {
			maze.Grid[i][j] = 1
		}
	}

	return maze
}

func (m *Maze) Generate(width, height int) {
	start := Point{x: width, y: height}

	m.Grid[start.y][start.x] = 0
	m.generateMaze(start)
}

func (m *Maze) generateMaze(current Point) {
	directions := []Point{
		{2, 0},
		{-2, 0},
		{0, 2},
		{0, -2},
	}

	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	for _, dir := range directions {
		next := Point{current.x + dir.x, current.y + dir.y}

		if next.x > 0 && next.x < m.Width-1 && next.y > 0 && next.y < m.Height-1 && m.Grid[next.y][next.x] == 1 {

			m.Grid[current.y+dir.y/2][current.x+dir.x/2] = 0
			m.Grid[next.y][next.x] = 0

			m.generateMaze(next)
		}
	}
}

func (m *Maze) Print() {
	for _, row := range m.Grid {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

func (m *Maze) Reset() {
	for y := range m.Grid {
		for x := range m.Grid[y] {
			if m.Grid[y][x] == 2 {
				m.Grid[y][x] = 0
			}
		}
	}
}
