package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	panic_fn(err)
	re, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	panic_fn(err)
	strs := string(data)
	fmt.Println(strs)
	total := 0
	for _, com := range re.FindAllString(strs, -1) {
		fmt.Println(com)
		sNums := strings.Split(com[4:len(com)-1], ",")
		v1, err := strconv.Atoi(sNums[0])
		panic_fn(err)
		v2, err := strconv.Atoi(sNums[1])
		panic_fn(err)
		total += v1 * v2
	}
	fmt.Println(total)
}

func panic_fn(err error) {
	if err != nil {
		panic(err)
	}
}
