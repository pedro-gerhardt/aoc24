package main

import (
	"aoc24-pleg/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	data, err := os.ReadFile("../input.txt")
	utils.Panic_fn(err)

	rules := make(map[string]mapset.Set[string])
	isRule := true
	soma := 0
	for _, line := range strings.Split(string(data), "\n") {
		if isRule {
			if line == "" {
				isRule = false
			} else {
				splitted := strings.Split(line, "|")
				if rules[splitted[0]] == nil {
					rules[splitted[0]] = mapset.NewSet[string]()
				}
				rules[splitted[0]].Add(splitted[1])
			}
		} else {
			valido := true
			splitted := strings.Split(line, ",")
			slices.Reverse(splitted)
			for i, s := range splitted {
				for _, n := range splitted[i:] {
					if rules[s] != nil && rules[s].ContainsOne(n) {
						valido = false
						break
					}
				}
				if !valido {
					break
				}
			}
			if valido {
				v, err := strconv.Atoi(splitted[len(splitted)/2])
				utils.Panic_fn(err)
				soma += v
			}
		}
	}

	fmt.Println(soma)
}
