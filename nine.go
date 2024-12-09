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
	// sb := make([]int, 0)

	input, err := os.ReadFile("nine.txt")
	if err != nil {
		fmt.Println(err)
	}

	intInput := make([]int, 0)

	// k := 0
	for i := range input {
		size, err := strconv.Atoi(string(input[i]))
		if err != nil {
			panic(err)
		}
		// if i%2 == 0 || i == 0 {

		// 	for j := 0; j < size; j = j + 1 {
		// 		sb = append(sb, k)
		// 	}
		// 	k++
		// } else {
		// 	for j := 0; j < size; j = j + 1 {
		// 		sb = append(sb, -1)
		// 	}
		// }
		intInput = append(intInput, size)
	}

	place := 0
	s := 0
	end := len(intInput) - 1
	for n := range intInput {
		if n <= end {
			if n%2 == 0 || n == 0 {
				if intInput[n] >= 0 {
					for i := 0; i < intInput[n]; i++ {
						s = s + (int(n/2) * place)
						place++
					}
				} else {
					for j := 0; j > intInput[n]; j-- {
						place++
					}
				}
			} else {
				temp := intInput[n]
				tempEnd := end
				for temp > 0 && tempEnd > n {
					if intInput[tempEnd] <= temp && intInput[tempEnd] > 0 {
						temp = temp - intInput[tempEnd]
						for i := 0; i < intInput[tempEnd]; i++ {
							s = s + (int(tempEnd/2) * place)
							place++
						}
						intInput[tempEnd] = -1 * intInput[tempEnd]
					}
					tempEnd = tempEnd - 2
				}
				if temp > 0 {
					for j := temp; j > 0; j-- {
						place++
					}
				}
			}
		}
	}

	// sum := 0
	// end := len(sb) - 1
	// for b := 0; b < len(sb); b++ {
	// 	if sb[b] == -1 {

	// 		for sb[end] == -1 {
	// 			end = end - 1
	// 		}
	// 		sb[b] = sb[end]
	// 		end = end - 1
	// 	}
	// 	if b <= end {
	// 		sum = sum + (b * sb[b])
	// 	}
	// }

	fmt.Println(s)
	// fmt.Println(sum)
}
