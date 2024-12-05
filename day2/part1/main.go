package main

import (
	"aoc24-pleg/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("122" < "13")
	data, err := os.ReadFile("../input.txt")
	utils.Panic_fn(err)

	total := 0
	for _, report := range strings.Split(string(data), "\n") {
		levels := strings.Split(report, " ")

		val0, err := strconv.ParseFloat(levels[0], 64)
		utils.Panic_fn(err)
		val1, err := strconv.ParseFloat(levels[1], 64)
		utils.Panic_fn(err)
		asc := val0 < val1

		skip := false
		for idx_base := range levels[1:] {
			idx := idx_base + 1

			val, err := strconv.ParseFloat(levels[idx], 64)
			utils.Panic_fn(err)
			val2, err := strconv.ParseFloat(levels[idx-1], 64)
			utils.Panic_fn(err)

			abs := math.Abs(val - val2)
			adj := abs > 0 && abs < 4
			all_inc := asc && val > val2
			all_dec := !asc && val < val2
			if !(adj && (all_inc || all_dec)) {
				skip = true
				break
			}
		}
		if !skip {
			total += 1
		}
	}

	fmt.Println(total)
}
