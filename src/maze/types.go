package maze

type Point struct {
	x, y int
}

type Maze struct {
	Grid   [][]int
	Width  int
	Height int
}
