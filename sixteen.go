package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
	"time"
)

type PointHeap []Point

func (h PointHeap) Len() int            { return len(h) }
func (h PointHeap) Less(i, j int) bool  { return h[i].value < h[j].value }
func (h PointHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PointHeap) Push(x interface{}) { *h = append(*h, x.(Point)) }

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

	input := make([][]string, 0)

	file, err := os.Open("sixteen.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	line := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		input = append(input, chars)
		line++
	}

	maze := setupMaze(len(input[0]), (len(input)), &input)
	start := generatePoint(&maze, &input, "S")
	end := generatePoint(&maze, &input, "E")

	fmt.Println(shortestPath(&maze, start, end))
}

type Point struct {
	x         int
	y         int
	wall      bool
	value     int
	direction string
	prev      *Point
}

func turnLeft(s string) string {
	switch s {
	case "⏩":
		return "🔼"
	case "🔼":
		return "⏪"
	case "⏪":
		return "🔽"
	case "🔽":
		return "⏩"
	default:
		return ""
	}
}

func turnRight(s string) string {
	switch s {
	case "⏩":
		return "🔽"
	case "🔼":
		return "⏩"
	case "⏪":
		return "🔼"
	case "🔽":
		return "⏪"
	default:
		return ""
	}
}

func shortestPath(maze *[][]Point, start Point, end Point) int {
	visited := map[string]bool{}
	start.value = 0
	start.direction = "⏩"
	q := &PointHeap{start}
	for q.Len() > 0 {
		point := heap.Pop(q).(Point)
		// fmt.Println(strPoint(point), point.value, point.direction)

		if strPoint(point) == strPoint(end) {
			return point.value
		}

		neighbour := getNeighbour(point, maze)
		if strPoint(neighbour) != "-1,-1" {
			neighbour.value = point.value + 1
			neighbour.prev = &point
			heap.Push(q, neighbour)
		}

		if !visited[strPoint(point)] {
			heap.Push(q, Point{
				x:         point.x,
				y:         point.y,
				wall:      false,
				value:     point.value + 1000,
				direction: turnLeft(point.direction),
				prev:      point.prev,
			})
			heap.Push(q, Point{
				x:         point.x,
				y:         point.y,
				wall:      false,
				value:     point.value + 1000,
				direction: turnRight(point.direction),
				prev:      point.prev,
			})
		}
		visited[strPoint(point)] = true
	}

	return -1
}

func setupMaze(sizeX int, sizeY int, input *[][]string) [][]Point {
	var maze [][]Point

	for j := 0; j < sizeY; j++ {
		var row []Point
		for i := 0; i < sizeX; i++ {
			if (*input)[j][i] == "#" {
				newPoint := Point{
					x:     i,
					y:     j,
					wall:  true,
					value: sizeX * sizeY,
					prev:  nil,
				}
				row = append(row, newPoint)
			} else {
				newPoint := Point{
					x:         i,
					y:         j,
					wall:      false,
					direction: "⏩",
					value:     sizeX * sizeY,
					prev:      nil,
				}
				row = append(row, newPoint)
			}
		}
		maze = append(maze, row)
	}
	return maze
}

func getNeighbour(p Point, maze *[][]Point) Point {
	var posX int
	var posY int

	switch p.direction {
	case "⏩":
		posX = p.x + 1
		posY = p.y
	case "🔼":
		posX = p.x
		posY = p.y - 1
	case "⏪":
		posX = p.x - 1
		posY = p.y
	case "🔽":
		posX = p.x
		posY = p.y + 1
	default:
		posX = -1
		posY = -1
	}

	if posX > 0 && posX < len((*maze)[0])-1 && posY > 0 && posY < len((*maze))-1 {
		point := (*maze)[posY][posX]
		if !point.wall && point != p {
			point.direction = p.direction
			return point
		}
	}
	return Point{
		x:     -1,
		y:     -1,
		wall:  true,
		value: -1,
		prev:  nil,
	}
}

func generatePoint(maze *[][]Point, input *[][]string, s string) Point {
	sizeX := len((*maze)[0])
	sizeY := len(*maze)
	var point Point

	for j := 0; j < sizeY; j++ {
		for i := 0; i < sizeX; i++ {
			if (*input)[j][i] == s {
				point = (*maze)[j][i]
			}
		}
	}
	return point
}

func strPoint(p Point) string {
	return strings.Join([]string{fmt.Sprint(p.x), fmt.Sprint(p.y)}, ",")
}
