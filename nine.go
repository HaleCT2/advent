package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()
	sb := make([]int, 0)

	input, err := os.ReadFile("nine.txt")
	if err != nil {
		fmt.Println(err)
	}

	k := 0
	for i := range input {
		size, err := strconv.Atoi(string(input[i]))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 || i == 0 {

			for j := 0; j < size; j = j + 1 {
				sb = append(sb, k)
			}
			k++
		} else {
			for j := 0; j < size; j = j + 1 {
				sb = append(sb, -1)
			}
		}
	}

	sum := 0
	end := len(sb) - 1
	for b := 0; b < len(sb); b++ {
		if sb[b] == -1 {

			for sb[end] == -1 {
				end = end - 1
			}
			sb[b] = sb[end]
			end = end - 1
		}
		if b <= end {
			sum = sum + (b * sb[b])
		}
	}
	fmt.Println(sum)
}
