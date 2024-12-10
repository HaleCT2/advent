package main

import (
	"bufio"
	"fmt"
	"os"
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

// (3,4) (5,5) -> (7,6), (1,3) ->

func findNodes(m []Coord, s int) []Coord {
	n := []Coord{}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if i != j {
				x, y := m[i].x-m[j].x, m[i].y-m[j].y
				nC := Coord{m[j].x - x, m[j].y - y}
				for nC.x < s && nC.y < s && nC.x >= 0 && nC.y >= 0 {
					n = append(n, nC)
					nC = Coord{nC.x - x, nC.y - y}
				}
				nC = Coord{m[j].x + x, m[j].y + y}
				for nC.x < s && nC.y < s && nC.x >= 0 && nC.y >= 0 {
					n = append(n, nC)
					nC = Coord{nC.x + x, nC.y + y}
				}
			}
		}
	}
	return n
}

func main() {
	defer timer("main")()
	input := make([][]string, 0)

	file, err := os.Open("eight.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ""))
	}

	m := make(map[string][]Coord)

	for i := range input {
		for j := range input[i] {
			if input[i][j] != "." {
				if len(m[input[i][j]]) == 0 {
					m[input[i][j]] = []Coord{}
				}
				m[input[i][j]] = append(m[input[i][j]], Coord{i, j})
			}
		}
	}

	n := make(map[string][]Coord)

	for key, value := range m {
		if len(value) >= 2 {
			n[key] = findNodes(value, len(input))
		}
	}

	processed := make(map[Coord]struct{})
	uniq := make([]Coord, 0)

	for _, v := range n {
		for _, value := range v {
			if _, ok := processed[value]; ok {
				continue
			}

			uniq = append(uniq, value)

			processed[value] = struct{}{}
		}
	}

	// fmt.Println(len(input))
	// fmt.Println(m)
	// fmt.Println(n)
	fmt.Println(len(uniq))
}
