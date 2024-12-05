package main

import (
	"aoc24-pleg/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("../input.txt")
	utils.Panic_fn(err)
	var arr1 []int
	map2 := make(map[int]int)
	for _, str := range strings.Split(string(dat), "\n") {
		splitted := strings.Split(str, " ")
		convertAndAppend(splitted[0], &arr1)
		convertAndCount(splitted[len(splitted)-1], &map2)
	}
	var total int
	for _, v := range arr1 {
		total += v * map2[v]
	}
	fmt.Println(total)
}

func convertAndAppend(number string, arr *[]int) {
	f1, err := strconv.Atoi(number)
	utils.Panic_fn(err)
	*arr = append(*arr, f1)
}

func convertAndCount(number string, m *map[int]int) {
	i1, err := strconv.Atoi(number)
	utils.Panic_fn(err)
	val, exists := (*m)[i1]
	if exists {
		(*m)[i1] = val + 1
	} else {
		(*m)[i1] = 1
	}
}
