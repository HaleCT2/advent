package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"time"
	"unicode/utf8"
)

type PointHeap []*Point

func (h PointHeap) Len() int            { return len(h) }
func (h PointHeap) Less(i, j int) bool  { return h[i].value < h[j].value }
func (h PointHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PointHeap) Push(x interface{}) { *h = append(*h, x.(*Point)) }

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()

	input := make([]Coord, 0)

	file, err := os.Open("eighteen.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	line := 0
	// num := 1024
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x, y int
		if _, err := fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y); err == nil {
			input = append(input, Coord{x, y})
		}
		line++
	}

	maze := setupMaze(71, 71)
	start := generatePoint(&maze, 0, 0)
	end := generatePoint(&maze, 70, 70)

	for _, c := range input {
		fmt.Println(c.x, c.y)
		maze[c.y][c.x].wall = true
		resetMaze(&maze)
		path := shortestPath(start, end)
		print(&maze, &path)
	}
}

type Point struct {
	x          int
	y          int
	wall       bool
	value      int
	neighbours []*Point
	prev       *Point
}

type Coord struct {
	x int
	y int
}

func shortestPath(start *Point, end *Point) []*Point {
	visited := map[string]bool{}
	start.value = 0
	q := &PointHeap{start}

	for q.Len() > 0 {
		point := heap.Pop(q).(*Point)

		if point == end {
			break
		}

		for _, neighbour := range point.neighbours {
			if !visited[strPoint(neighbour)] {
				visited[strPoint(neighbour)] = true
				heap.Push(q, neighbour)
				neighbour.value = point.value + 1
				neighbour.prev = point
			}
		}
	}

	path := []*Point{}
	point := end
	for point != start {
		path = append([]*Point{point}, path...)
		point = point.prev
	}

	fmt.Println(end.value)
	return path
}

func resetMaze(maze *[][]*Point) {
	for j, row := range *maze {
		for i := range row {
			point := (*maze)[j][i]
			if !point.wall {
				point.value = len(*maze) * len((*maze)[0])
				point.prev = nil
				point.neighbours = nil
			}
		}
	}
	setNeighbours(maze)
}

func setupMaze(sizeX int, sizeY int) [][]*Point {
	var maze [][]*Point

	for j := 0; j < sizeY; j++ {
		var row []*Point
		for i := 0; i < sizeX; i++ {
			newPoint := Point{
				x:     i,
				y:     j,
				wall:  false,
				value: sizeX * sizeY,
				prev:  nil,
			}
			row = append(row, &newPoint)

		}
		maze = append(maze, row)
	}

	setNeighbours(&maze)
	return maze
}

func setNeighbours(maze *[][]*Point) {
	for j, row := range *maze {
		for i := range row {
			point := (*maze)[j][i]
			point.neighbours = getPointNeighbours(point, maze)
		}
	}
}

func getPointNeighbours(p *Point, maze *[][]*Point) []*Point {
	sizeX := float64(len((*maze)[0]))
	sizeY := float64(len(*maze))
	var neighbours []*Point

	for row := math.Max(float64(p.x-1), 0); row < math.Min(float64(p.x+2), sizeX); row++ {
		for col := math.Max(float64(p.y-1), 0); col < math.Min(float64(p.y+2), sizeY); col++ {
			point := (*maze)[int(col)][int(row)]
			if !point.wall && point != p && (int(row) == p.x || int(col) == p.y) {
				neighbours = append(neighbours, point)
			}
		}
	}
	return neighbours
}

func print(maze *[][]*Point, path *[]*Point) {
	for j := 0; j < len(*maze); j++ {
		row := ""
		for i := 0; i < len((*maze)[0]); i++ {
			point := (*maze)[j][i]
			icon, _ := utf8.DecodeRune([]byte{0xE2, 0xAC, 0x9C})
			if point.wall {
				icon, _ = utf8.DecodeRune([]byte{0xE2, 0xAC, 0x9B})
			}
			if path != nil {
				if slices.Contains(*path, point) {
					icon, _ = utf8.DecodeRune([]byte{0xE2, 0xAD, 0x95})
				}
				if point == (*path)[0].prev {
					icon, _ = utf8.DecodeRune([]byte{0xE2, 0xAD, 0x90})
				}
				if point == (*path)[len(*path)-1] {
					icon, _ = utf8.DecodeRune([]byte{0xF0, 0x9F, 0x8E, 0xAF})
				}
			}
			row += string(icon)
		}
		fmt.Println(row)
	}
}

func generatePoint(maze *[][]*Point, x int, y int) *Point {
	return (*maze)[y][x]
}

func strPoint(p *Point) string {
	return strings.Join([]string{fmt.Sprint(p.x), fmt.Sprint(p.y)}, ",")
}
