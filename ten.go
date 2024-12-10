package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Coord struct {
	x, y int
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func findTrailheads(m [][]int) []Coord {
	trailheads := make([]Coord, 0)
	for x := range m {
		for y := range m[x] {
			if m[x][y] == 0 {
				head := Coord{x, y}
				trailheads = append(trailheads, head)
			}
		}
	}
	return trailheads
}

func removeDuplicates(coords []Coord) []Coord {
	allKeys := make(map[Coord]bool)
	list := []Coord{}
	for _, item := range coords {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func inbounds(c Coord, l int) bool {
	if c.x < l && c.y < l && c.x >= 0 && c.y >= 0 {
		return true
	} else {
		return false
	}
}

func evaluate(c Coord, m [][]int, d int, u *[]Coord) int {
	if m[c.x][c.y] == 9 {
		*u = append(*u, Coord{c.x, c.y})
		return 1
	}

	switch true {
	case inbounds(Coord{c.x + 1, c.y}, len(m)) && m[c.x+1][c.y]-1 == m[c.x][c.y] && d < 1:
		return evaluate(Coord{c.x + 1, c.y}, m, 0, u) + evaluate(Coord{c.x, c.y}, m, 1, u)
	case inbounds(Coord{c.x - 1, c.y}, len(m)) && m[c.x-1][c.y]-1 == m[c.x][c.y] && d < 2:
		return evaluate(Coord{c.x - 1, c.y}, m, 0, u) + evaluate(Coord{c.x, c.y}, m, 2, u)
	case inbounds(Coord{c.x, c.y + 1}, len(m)) && m[c.x][c.y+1]-1 == m[c.x][c.y] && d < 3:
		return evaluate(Coord{c.x, c.y + 1}, m, 0, u) + evaluate(Coord{c.x, c.y}, m, 3, u)
	case inbounds(Coord{c.x, c.y - 1}, len(m)) && m[c.x][c.y-1]-1 == m[c.x][c.y]:
		return evaluate(Coord{c.x, c.y - 1}, m, 0, u)
	default:
		return 0
	}
}

func main() {
	defer timer("main")()
	input := make([][]int, 0)

	file, err := os.Open("ten.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	line := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "")
		input = append(input, make([]int, 0))
		input[line] = make([]int, len(nums))

		for num := range nums {
			n, err := strconv.Atoi(nums[num])
			if err != nil {
				panic(err)
			}
			input[line][num] = n
		}
		line++
	}

	trailheads := findTrailheads(input)

	sum := 0
	for c := range trailheads {
		uniq := make([]Coord, 0)
		// evaluate(Coord{trailheads[c].x, trailheads[c].y}, input, 0, &uniq)
		// sum = sum + len(removeDuplicates(uniq))
		sum = sum + evaluate(Coord{trailheads[c].x, trailheads[c].y}, input, 0, &uniq)
	}

	fmt.Println(sum)
}
