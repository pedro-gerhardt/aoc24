package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("122" < "13")
	data, err := os.ReadFile("../input.txt")
	panic_fn(err)

	total := 0
	for _, report := range strings.Split(string(data), "\n") {
		levels := strings.Split(report, " ")

		val0, err := strconv.ParseFloat(levels[0], 64)
		panic_fn(err)
		val1, err := strconv.ParseFloat(levels[1], 64)
		panic_fn(err)
		asc := val0 < val1

		skip := 0
		unsafe := -1
		for idx_base := range levels[1:] {
			idx := idx_base + 1

			val, err := strconv.ParseFloat(levels[idx], 64)
			panic_fn(err)
			val2, err := strconv.ParseFloat(levels[idx-1], 64)
			panic_fn(err)

			abs := math.Abs(val - val2)
			adj := abs > 0 && abs < 4
			all_inc := asc && val > val2
			all_dec := !asc && val < val2
			if !(adj && (all_inc || all_dec)) {
				unsafe = idx
				break
			}
		}
		if unsafe != -1 {
			fmt.Println("---------------------------------")
			fmt.Println(levels)
			levelsCopy := make([]string, len(levels))
			copy(levelsCopy, levels)
			levels = slices.Delete(levels, unsafe, unsafe+1)
			fmt.Println(unsafe)
			fmt.Println(levels)
			for idx_base := range levels[1:] {
				idx := idx_base + 1

				val, err := strconv.ParseFloat(levels[idx], 64)
				panic_fn(err)
				val2, err := strconv.ParseFloat(levels[idx-1], 64)
				panic_fn(err)

				abs := math.Abs(val - val2)
				adj := abs > 0 && abs < 4
				all_inc := asc && val > val2
				all_dec := !asc && val < val2
				if !(adj && (all_inc || all_dec)) {
					skip += 1
					break
				}
			}
			levelsCopy = slices.Delete(levelsCopy, unsafe-1, unsafe)
			for idx_base := range levelsCopy[1:] {
				idx := idx_base + 1

				val, err := strconv.ParseFloat(levelsCopy[idx], 64)
				panic_fn(err)
				val2, err := strconv.ParseFloat(levelsCopy[idx-1], 64)
				panic_fn(err)

				abs := math.Abs(val - val2)
				adj := abs > 0 && abs < 4
				all_inc := asc && val > val2
				all_dec := !asc && val < val2
				if !(adj && (all_inc || all_dec)) {
					skip += 1
					break
				}
			}
			fmt.Println(skip)
		}
		if skip < 2 {
			total += 1
		}
	}

	fmt.Println(total)
}

func panic_fn(err error) {
	if err != nil {
		panic(err)
	}
}
