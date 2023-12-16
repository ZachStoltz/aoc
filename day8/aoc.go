package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	digit := partOneAndTwo()

	fmt.Printf("Result: %d\n", digit)
}

func partOneAndTwo() int {
	file, err := os.ReadFile("/Users/zachstoltz/develop/aoc/day8/p1")
	if err != nil {
			log.Fatal(err)
	}


	contents := string(file)
	lines := strings.Split(contents, "\n\n")
	itr := 0
	nodes := make(map[string]map[string]string)
	steps := strings.Split(lines[0], "")
	rest := strings.Split(lines[1], "\n")

	for _, val := range rest {
		fmt.Println(val)
		parts := strings.Split(val, " = ")
		steps := strings.Split(parts[1], ", ")

		nodes[parts[0]] = make(map[string]string)
		nodes[parts[0]]["L"] = steps[0][1:]
		nodes[parts[0]]["R"] = steps[1][:len(steps[1]) -1]
	}

	curr := "AAA"

	return walk(itr, curr, steps, nodes)
}

func walk (itr int, curr string, steps []string, nodes map[string]map[string]string) int {
	if (curr == "ZZZ") {
		return itr
	}

	step := steps[itr]
	next := nodes[curr][step]
	itr += 1

	// fmt.Printf("step %s next %s nextItr %d\n", step, next, itr)

	if (itr > len(steps) - 1) {
		steps = slices.Insert(steps, len(steps), steps...)
	}

	return walk(itr, next, steps, nodes)
}