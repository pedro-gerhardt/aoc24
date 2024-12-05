package main

import (
	"aoc24-pleg/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	utils.Panic_fn(err)
	re, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	utils.Panic_fn(err)
	strs := string(data)
	fmt.Println(strs)
	total := 0
	for _, com := range re.FindAllString(strs, -1) {
		fmt.Println(com)
		sNums := strings.Split(com[4:len(com)-1], ",")
		v1, err := strconv.Atoi(sNums[0])
		utils.Panic_fn(err)
		v2, err := strconv.Atoi(sNums[1])
		utils.Panic_fn(err)
		total += v1 * v2
	}
	fmt.Println(total)
}
