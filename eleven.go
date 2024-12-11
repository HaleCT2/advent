package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func length(i int) int {
	return (int)(math.Log10(float64(i)) + 1)
}

func leftSplit(i int) int {
	return int(i / int((math.Pow10(length(i) / 2))))
}

func rightSplit(i int) int {
	return int(i % int((math.Pow10(length(i) / 2))))
}

func blink(stones *[]int) {
	var wg sync.WaitGroup
	c := make(chan int)
	for s := 0; s < len(*stones); s++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i == 0 {
				(*stones)[s] = 1
			} else if length(i)%2 == 1 {
				(*stones)[s] *= 2024
			} else {
				wg.Add(1)
				(*stones)[s] = rightSplit(i)
				c <- leftSplit(i)
			}
		}((*stones)[s])
	}

	go func() {
		for t := range c {
			*stones = append(*stones, t)
			wg.Done()
		}
	}()

	wg.Wait()
}

func main() {
	defer timer("main")()

	content, err := os.ReadFile("eleven.txt")
	if err != nil {
		fmt.Println(err)
	}

	sContent := strings.Split(strings.Trim(string(content), "\n"), " ")
	stones := make([]int, 0)

	for s := range sContent {
		num, err := strconv.Atoi(sContent[s])
		if err != nil {
			panic(err)
		}
		stones = append(stones, num)
	}

	blinks := 25
	for i := 0; i < blinks; i++ {
		fmt.Println(i, len(stones))
		blink(&stones)
	}
	fmt.Println(blinks, len(stones))
}
