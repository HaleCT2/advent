package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func padding(n float64) float64 {
	return math.Pow10(int(math.Log10(n) + 1))
}

func evaluate(p float64, n []float64) bool {
	if len(n) == 2 {
		// if (p/n[1]) == n[0] || p-n[1] == n[0] {
		if (p/n[1]) == n[0] || p-n[1] == n[0] || (p-n[1])/padding(n[1]) == n[0] {
			return true
		} else {
			return false
		}
	}

	// return evaluate(p/n[len(n)-1], n[:len(n)-1]) || evaluate(p-n[len(n)-1], n[:len(n)-1])
	return evaluate(p/n[len(n)-1], n[:len(n)-1]) || evaluate(p-n[len(n)-1], n[:len(n)-1]) || evaluate((p-n[len(n)-1])/padding(n[len(n)-1]), n[:len(n)-1])
}

func main() {
	defer timer("main")()
	input := make(map[float64][]float64)

	file, err := os.Open("seven.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		product, err := strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}
		input[float64(product)] = make([]float64, 0)
		nums := strings.Split(line[1][1:], " ")
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			input[float64(product)] = append(input[float64(product)], float64(n))
		}
	}

	sum := 0.0
	for key, value := range input {
		if evaluate(key, value) {
			sum = sum + key
		}
	}

	fmt.Println(int(sum))
}
