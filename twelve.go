package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type Coord struct {
	kind      string
	x         int
	y         int
	edges     int
	neighbors []*Coord
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func strPoint(p *Coord) string {
	return strings.Join([]string{fmt.Sprint(p.x), fmt.Sprint(p.y)}, ",")
}

// func pprint(s *[][]*Coord) {
// 	for x := range *s {
// 		for _, c := range (*s)[x] {
// 			fmt.Print(c.kind)
// 		}
// 		fmt.Print("\n")
// 	}
// }

func setupMap() [][]*Coord {
	input := make([][]string, 0)
	farm := make([][]*Coord, 0)
	file, err := os.Open("twelve.txt")
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

	for x := range input {
		var row []*Coord
		for y, c := range input[x] {
			newCoord := Coord{
				kind:      c,
				x:         x,
				y:         y,
				edges:     0,
				neighbors: nil,
			}
			row = append(row, &newCoord)
		}
		farm = append(farm, row)
	}

	setNeighbors(&farm)

	return farm
}

func setNeighbors(farm *[][]*Coord) {
	for x, row := range *farm {
		for y := range row {
			point := (*farm)[x][y]
			point.neighbors = getNeighbors(point, farm)
		}
	}
}

func getNeighbors(p *Coord, farm *[][]*Coord) []*Coord {
	sizeX := float64(len((*farm)[0]))
	sizeY := float64(len(*farm))
	var neighbors []*Coord

	for row := math.Max(float64(p.x-1), 0); row < math.Min(float64(p.x+2), sizeX); row++ {
		for col := math.Max(float64(p.y-1), 0); col < math.Min(float64(p.y+2), sizeY); col++ {
			point := (*farm)[int(row)][int(col)]
			if point.kind == p.kind && point != p && (int(row) == p.x || int(col) == p.y) {
				neighbors = append(neighbors, point)
			}
		}
	}
	p.edges = 4 - len(neighbors)
	return neighbors
}

func calcSize(farm *[][]*Coord) int {
	visited := map[string]bool{}

	m := make(map[string]int)

	for x, row := range *farm {
		for y := range row {
			point := (*farm)[x][y]
			if !visited[strPoint(point)] {
				// fmt.Println(strPoint(point))
				visited[strPoint(point)] = true
				size := 1
				perimeter := point.edges
				q := []*Coord{point}

				for len(q) > 0 {
					point := q[0]
					q = q[1:]

					for _, neighbor := range point.neighbors {
						if !visited[strPoint(neighbor)] {
							visited[strPoint(neighbor)] = true
							size += 1
							perimeter += neighbor.edges
							q = append(q, neighbor)
						}
					}
				}

				m[point.kind] += (size * perimeter)
			}
		}
	}

	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	defer timer("main")()
	farm := setupMap()
	// pprint(&farm)
	fmt.Println(calcSize(&farm))
}
