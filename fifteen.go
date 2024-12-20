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
	x    int
	y    int
	box  bool
	bot  bool
	wall bool
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func pprint(x int, y int, f *map[string]*Coord) {
	m := make([][]string, y)
	for i := range m {
		m[i] = make([]string, x)
		for j := range m[i] {
			m[i][j] = "."
		}
	}

	for _, c := range *f {
		if c.bot {
			m[c.x][c.y] = "@"
		} else if c.box {
			m[c.x][c.y] = "0"
		} else if c.wall {
			m[c.x][c.y] = "#"
		}
	}

	for i := range m {
		fmt.Println("")
		for j := range m[i] {
			fmt.Print(m[i][j])
		}
	}
	fmt.Println("")

}

func next(x int, y int, r *Coord, m string, f *map[string]*Coord) {
	switch m {
	case "^":
		for i := r.x - 1; i > 0; i-- {
			if _, ok := (*f)[strInts(i, r.y)]; ok && (*f)[strInts(i, r.y)].box {
			} else if _, ok := (*f)[strInts(i, r.y)]; ok && !(*f)[strInts(i, r.y)].wall {
				(*f)[strInts(i, r.y)].box = true
				(*f)[strInts(r.x-1, r.y)].box = false
				r.bot = false
				(*f)[strInts(r.x-1, r.y)].bot = true
				break
			} else {
				break
			}
		}
	case ">":
		for j := r.y + 1; j < y; j++ {
			if _, ok := (*f)[strInts(r.x, j)]; ok && (*f)[strInts(r.x, j)].box {
			} else if _, ok := (*f)[strInts(r.x, j)]; ok && !(*f)[strInts(r.x, j)].wall {
				(*f)[strInts(r.x, j)].box = true
				(*f)[strInts(r.x, r.y+1)].box = false
				r.bot = false
				(*f)[strInts(r.x, r.y+1)].bot = true
				break
			} else {
				break
			}
		}
	case "v":
		for i := r.x + 1; i < x; i++ {
			if _, ok := (*f)[strInts(i, r.y)]; ok && (*f)[strInts(i, r.y)].box {
			} else if _, ok := (*f)[strInts(i, r.y)]; ok && !(*f)[strInts(i, r.y)].wall {
				(*f)[strInts(i, r.y)].box = true
				r.bot = false
				(*f)[strInts(r.x+1, r.y)].box = false
				(*f)[strInts(r.x+1, r.y)].bot = true
				break
			} else {
				break
			}
		}
	case "<":
		for j := r.y - 1; j > 0; j-- {
			if _, ok := (*f)[strInts(r.x, j)]; ok && (*f)[strInts(r.x, j)].box {
			} else if _, ok := (*f)[strInts(r.x, j)]; ok && !(*f)[strInts(r.x, j)].wall {
				(*f)[strInts(r.x, j)].box = true
				r.bot = false
				(*f)[strInts(r.x, r.y-1)].box = false
				(*f)[strInts(r.x, r.y-1)].bot = true
				break
			} else {
				break
			}
		}
	}
}

func move(x int, y int, d *[]string, f *map[string]*Coord) {
	for len(*d) > 0 {
		m := (*d)[0]
		*d = (*d)[1:]

		for _, c := range *f {
			if c.bot {
				next(x, y, c, m, f)
				break
			}
		}
	}

}

func count(f *map[string]*Coord) int {
	sum := 0
	for _, v := range *f {
		if v.box {
			sum += (100 * (v.x)) + (v.y)
		}
	}
	return sum
}

func strInts(i int, j int) string {
	return strings.Join([]string{strconv.Itoa(i), strconv.Itoa(j)}, ",")

}

func main() {
	defer timer("main")()

	input := make([][]string, 0)
	directions := make([]string, 0)
	floor := make(map[string]*Coord)

	file, err := os.Open("fifteen.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	newline := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			newline = true
		}

		if !newline {
			chars := strings.Split(scanner.Text(), "")
			input = append(input, chars)
		} else {
			chars := strings.Split(scanner.Text(), "")
			directions = append(directions, chars...)
		}
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			switch input[i][j] {
			case "O":
				floor[strInts(i, j)] = &Coord{x: i,
					y:    j,
					box:  true,
					bot:  false,
					wall: false}
			case "@":
				floor[strInts(i, j)] = &Coord{x: i,
					y:    j,
					box:  false,
					bot:  true,
					wall: false}
			case "#":
				floor[strInts(i, j)] = &Coord{x: i,
					y:    j,
					box:  false,
					bot:  false,
					wall: true}
			default:
				floor[strInts(i, j)] = &Coord{x: i,
					y:    j,
					box:  false,
					bot:  false,
					wall: false}
			}
		}
	}
	pprint(50, 50, &floor)
	move(50, 50, &directions, &floor)
	pprint(50, 50, &floor)
	fmt.Println(count(&floor))
}
