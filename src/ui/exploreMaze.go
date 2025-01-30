package ui

import (
	"GoMaze/src/maze"
	"sync"
	"time"
)

func ExploreMaze(currentMaze *maze.Maze, updateChan chan<- [][]int) {
	var wg sync.WaitGroup

	wg.Add(1)
	exploreMaze(currentMaze, 1, 1, &wg, updateChan)
	wg.Wait()
}

func exploreMaze(maze *maze.Maze, x, y int, wg *sync.WaitGroup, updateChan chan<- [][]int) {
	defer wg.Done()

	maze.Grid[x][y] = 2
	updateChan <- maze.Grid

	time.Sleep(100 * time.Millisecond)

	directions := []struct{ dx, dy int }{
		{0, -1},
		{-1, 0},
		{1, 0},
		{0, 1},
	}

	for _, dir := range directions {
		nx, ny := x+dir.dx, y+dir.dy
		if nx >= 0 && nx < len(maze.Grid) && ny >= 0 && ny < len(maze.Grid[0]) && maze.Grid[nx][ny] == 0 {
			wg.Add(1)
			go exploreMaze(maze, nx, ny, wg, updateChan)
		}
	}
}
