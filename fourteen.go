package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"time"
)

type Robot struct {
	locX  int
	locY  int
	veloX int
	veloY int
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func pprint(x int, y int, s *[]*Robot) {
	m := make([][]int, y)
	for i := range m {
		m[i] = make([]int, x)
		for j := range m[i] {
			m[i][j] = 0
		}
	}

	for _, r := range *s {
		// fmt.Println(r.locX, r.locY, r.veloX, r.veloY)
		m[r.locY][r.locX] += 1
	}

	for i := range m {
		fmt.Println("")
		for j := range m[i] {
			if m[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("X")
			}
		}
	}
	fmt.Println("")

}

func getRobots() []*Robot {
	tiles := make([]*Robot, 0)

	file, err := os.Open("fourteen.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	locPattern := regexp.MustCompile(`p=-?\d+,-?\d+`)
	velPattern := regexp.MustCompile(`v=-?\d+,-?\d+`)

	line := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		locString := locPattern.FindString(scanner.Text())
		velString := velPattern.FindString(scanner.Text())

		var lx, ly int
		var vx, vy int

		if _, err := fmt.Sscanf(locString, "p=%d,%d", &lx, &ly); err == nil {
			if _, err := fmt.Sscanf(velString, "v=%d,%d", &vx, &vy); err == nil {
				tiles = append(tiles, &Robot{locX: lx,
					locY:  ly,
					veloX: vx,
					veloY: vy})
			}
		}
		line++
	}
	return tiles
}

func move(x int, y int, s *[]*Robot) int {
	t := 0
	min := math.MaxInt
	i := 0

	for t < 10000 {
		for _, r := range *s {
			newLocX := (r.locX + r.veloX) % x
			newLocY := (r.locY + r.veloY) % y
			if newLocX < 0 {
				newLocX += x
			}
			if newLocY < 0 {
				newLocY += y
			}
			r.locX = newLocX
			r.locY = newLocY
		}
		t += 1
		c := count(x, y, s)
		if c < min {
			pprint(x, y, s)
			min = c
			i = t
		}
	}
	return i
}

func count(x int, y int, s *[]*Robot) int {
	s1, s2, s3, s4 := 0, 0, 0, 0
	for _, r := range *s {
		if r.locX < int(x/2) && r.locY < int(y/2) {
			s1 += 1
		}
		if r.locX < int(x/2) && r.locY > int(y/2) {
			s2 += 1
		}
		if r.locX > int(x/2) && r.locY < int(y/2) {
			s3 += 1
		}
		if r.locX > int(x/2) && r.locY > int(y/2) {
			s4 += 1
		}
	}
	return s1 * s2 * s3 * s4
}

func main() {
	defer timer("main")()
	tiles := getRobots()

	pprint(101, 103, &tiles)
	tree := move(101, 103, &tiles)

	fmt.Println(tree)
}
