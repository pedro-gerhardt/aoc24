package main

import (
	"aoc24-pleg/utils"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type DOs struct {
	index int
	do    bool
}

func main() {
	data, err := os.ReadFile("input.txt")
	utils.Panic_fn(err)
	re, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	utils.Panic_fn(err)
	redo, err := regexp.Compile(`do\(\)`)
	utils.Panic_fn(err)
	redont, err := regexp.Compile(`don't\(\)`)
	utils.Panic_fn(err)

	strs := string(data)
	i_redo := redo.FindAllStringIndex(strs, -1)
	i_redont := redont.FindAllStringIndex(strs, -1)

	total := 0
	list := []DOs{{0, true}}
	var maior int
	pivot := 0
	if len(i_redo) > len(i_redont) {
		maior = len(i_redo)
	} else {
		maior = len(i_redont)
	}
	for i := range maior {
		if i < len(i_redo) {
			list = append(list, DOs{i_redo[i][0], true})
		}
		if i < len(i_redont) {
			list = append(list, DOs{i_redont[i][0], false})
		}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].index < list[j].index
	})
	fmt.Println(list)
	for _, idxs := range re.FindAllStringIndex(strs, -1) {
		for {
			if pivot == len(list)-1 || list[pivot+1].index > idxs[0] {
				break
			}
			pivot++
		}
		if list[pivot].do {
			com := strs[idxs[0]:idxs[1]]
			fmt.Println(com)
			sNums := strings.Split(com[4:len(com)-1], ",")
			v1, err := strconv.Atoi(sNums[0])
			utils.Panic_fn(err)
			v2, err := strconv.Atoi(sNums[1])
			utils.Panic_fn(err)
			total += v1 * v2
		}
	}
	fmt.Println(total)
}
