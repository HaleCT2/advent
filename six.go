package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Coord struct {
	x, y int
}

type Gen struct {
	s [][]string
	d string
	i int
}

func pprint(s [][]string) {
	for i := range s {
		fmt.Println(s[i])
	}
	fmt.Println("")
}

func popSlice() [][]string {
	content, err := os.ReadFile("six.txt")
	if err != nil {
		fmt.Println(err)
	}

	sliceData := strings.Split(string(content), "\n")
	slice := make([][]string, len(sliceData))

	for i := range slice {
		slice[i] = strings.Split(sliceData[i], "")
	}

	return slice
}

func findGuard(g Gen) Coord {
	guard := Coord{}
	for i := range g.s {
		for j := range g.s[i] {
			if g.s[i][j] == "^" {
				guard = Coord{i, j}
			} else if g.s[i][j] == ">" {
				guard = Coord{i, j}
			} else if g.s[i][j] == "<" {
				guard = Coord{i, j}
			} else if g.s[i][j] == "v" {
				guard = Coord{i, j}
			}
		}
	}

	return guard
}

func checkPos(g Gen, c Coord) bool {
	clear := true
	if g.s[c.x][c.y] == "#" {
		clear = false
	}
	return clear
}

func checkEscape(g Gen) bool {
	escaped := false
	if findGuard(g).x == 0 || findGuard(g).y == 0 {
		escaped = true
	} else if findGuard(g).x == len(g.s)-1 || findGuard(g).y == len(g.s)-1 {
		escaped = true
	}

	return escaped
}

func movePos(g Gen) Gen {
	next := Gen{g.s, g.d, g.i}
	guardPos := findGuard(g)

	switch g.d {
	case "N":
		newPos := Coord{guardPos.x - 1, guardPos.y}
		if checkPos(g, newPos) {
			if next.s[newPos.x][newPos.y] == "." {
				next.i++
			}
			next.s[newPos.x][newPos.y] = "^"
			next.s[guardPos.x][guardPos.y] = "|"
		} else {
			next.d = "E"
			newPos := Coord{guardPos.x, guardPos.y + 1}
			if next.s[newPos.x][newPos.y] == "." {
				next.i++
			}
			next.s[newPos.x][newPos.y] = ">"
			next.s[guardPos.x][guardPos.y] = "+"
		}
	case "E":
		newPos := Coord{guardPos.x, guardPos.y + 1}
		if checkPos(g, newPos) {
			if next.s[newPos.x][newPos.y] == "." {
				next.i++
			}
			next.s[newPos.x][newPos.y] = ">"
			next.s[guardPos.x][guardPos.y] = "-"
		} else {
			next.d = "S"
			newPos := Coord{guardPos.x + 1, guardPos.y}
			if next.s[newPos.x][newPos.y] == "." {
				next.i++
			}
			next.s[newPos.x][newPos.y] = "v"
			next.s[guardPos.x][guardPos.y] = "+"
		}
	case "S":
		newPos := Coord{guardPos.x + 1, guardPos.y}
		if checkPos(g, newPos) {
			if next.s[newPos.x][newPos.y] == "." {
				next.i++
			}
			next.s[newPos.x][newPos.y] = "v"
			next.s[guardPos.x][guardPos.y] = "|"
		} else {
			next.d = "W"
			newPos := Coord{guardPos.x, guardPos.y - 1}
			if next.s[newPos.x][newPos.y] == "." {
				next.i++
			}
			next.s[newPos.x][newPos.y] = "<"
			next.s[guardPos.x][guardPos.y] = "+"
		}
	case "W":
		newPos := Coord{guardPos.x, guardPos.y - 1}
		if checkPos(g, newPos) {
			if next.s[newPos.x][newPos.y] == "." {
				next.i++
			}
			next.s[newPos.x][newPos.y] = "<"
			next.s[guardPos.x][guardPos.y] = "-"
		} else {
			next.d = "N"
			newPos := Coord{guardPos.x - 1, guardPos.y}
			if next.s[newPos.x][newPos.y] == "." {
				next.i++
			}
			next.s[newPos.x][newPos.y] = "^"
			next.s[guardPos.x][guardPos.y] = "+"
		}
	default:
		fmt.Println("The guard is lost!")
	}

	return next
}

func checkTurns(g Gen) bool {
	surrounded := false
	guardPos := findGuard(g)

	left := make([]string, len(g.s[guardPos.x][:guardPos.y]))
	copy(left, g.s[guardPos.x][:guardPos.y])
	slices.Reverse(left)
	if slices.Index(left, "#") != -1 {
		left = left[:slices.Index(left, "#")]
	}
	slices.Reverse(left)

	right := make([]string, len(g.s[guardPos.x][guardPos.y:len(g.s[guardPos.x])]))
	copy(right, g.s[guardPos.x][guardPos.y:len(g.s[guardPos.x])])
	if slices.Index(right, "#") != -1 {
		right = right[:slices.Index(right, "#")]
	}

	up := make([]string, guardPos.x)
	for i := range up {
		up[i] = g.s[i][guardPos.y]
	}
	slices.Reverse(up)
	if slices.Index(up, "#") != -1 {
		up = up[:slices.Index(up, "#")]
	}
	slices.Reverse(up)

	down := make([]string, len(g.s[guardPos.x])-guardPos.x)
	for j := range down {
		down[j] = g.s[j+guardPos.x][guardPos.y]
	}
	if slices.Index(down, "#") != -1 {
		down = down[:slices.Index(down, "#")]
	}

	if slices.Contains(left, "+") &&
		slices.Contains(right, "+") &&
		slices.Contains(up, "+") &&
		slices.Contains(down, "+") {
		surrounded = true
	}

	// fmt.Println(left)
	// fmt.Println(right)
	// fmt.Println(up)
	// fmt.Println(down)

	return surrounded
}

func checkLoop(g Gen) bool {
	stuck := false
	guardPos := findGuard(g)

	for i := range g.s {
		for j := range g.s[i] {
			if g.s[i][j] == "+" &&
				g.s[i][guardPos.y] == "+" &&
				g.s[guardPos.x][j] == "+" {

				switch g.d {
				case "N":
					newPos := Coord{guardPos.x - 1, guardPos.y}
					if !checkPos(g, newPos) {
						stuck = true
					}
				case "E":
					newPos := Coord{guardPos.x, guardPos.y + 1}
					if !checkPos(g, newPos) {
						stuck = true
					}
				case "S":
					newPos := Coord{guardPos.x + 1, guardPos.y}
					if !checkPos(g, newPos) {
						stuck = true
					}
				case "W":
					newPos := Coord{guardPos.x, guardPos.y - 1}
					if !checkPos(g, newPos) {
						stuck = true
					}
				default:
					fmt.Println("The guard is lost!")
				}

				if checkTurns(g) {
					stuck = true
				}
			}
		}
	}

	return stuck
}

func main() {
	blocks := 0
	currentGen := Gen{popSlice(), "N", 1}

	for !checkEscape(currentGen) {
		currentGen = movePos(currentGen)
	}

	path := currentGen.s

	for i := range currentGen.s {
		for j := range currentGen.s[i] {
			if path[i][j] != "." {
				currentGen = Gen{popSlice(), "N", 1}
				currentGen.s[i][j] = "#"
				for !checkEscape(currentGen) && !checkLoop(currentGen) {
					currentGen = movePos(currentGen)
				}

				if checkLoop(currentGen) {
					blocks++
					fmt.Println(blocks)
					pprint(currentGen.s)
				}
			}
		}
	}
	fmt.Println(blocks)
}
