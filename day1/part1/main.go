package main

import (
	"aoc24-pleg/utils"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	utils.Panic_fn(err)
	var arr1 []float64
	var arr2 []float64
	for _, str := range strings.Split(string(dat), "\n") {
		splitted := strings.Split(str, " ")
		convertAndAppend(splitted[0], &arr1)
		convertAndAppend(splitted[len(splitted)-1], &arr2)
	}
	slices.Sort(arr1)
	slices.Sort(arr2)
	var total float64
	for i := range arr1 {
		total += math.Abs((float64)(arr1[i]) - arr2[i])
	}
	fmt.Println(int(total))
}

func convertAndAppend(number string, arr *[]float64) {
	f1, err := strconv.ParseFloat(number, 64)
	utils.Panic_fn(err)
	*arr = append(*arr, f1)
}
