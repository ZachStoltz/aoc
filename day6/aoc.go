package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partOneAndTwo()
}

func partOneAndTwo() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day6/p1")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	digit := solve(scanner.Scan(), scanner, [][]int{})
	finalPrint := fmt.Sprintf("Result: %d\n", digit)

	fmt.Print(finalPrint)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}

func solve(hasMoreLines bool, scanner *bufio.Scanner, values [][]int) int {
	if (hasMoreLines) {
		line := scanner.Text()
		line = regexp.MustCompile(`\w+:.\s+`).ReplaceAllString(line, "")
		line = regexp.MustCompile(`\s+`).ReplaceAllString(line, " ")

		numValues := strings.Split(line, " ")
		nums := []int{}

		for _, val := range numValues {
			num, _ := strconv.Atoi(val)
			nums = append(nums, num)
		}

		values = append(values, nums)

		return solve(scanner.Scan(), scanner, values)
	}

	total := 1

	raceTimes := values[0]
	records := values[1]

		for idx, time := range raceTimes {
			record := records[idx]
			winners := 0

			for i := 0; i <= time; i++ {
				score := i * (time - i)

				if (score > record) {
					winners += 1
				}
			}

			if (winners != 0) {
				total *= winners
			}
		}

	return total
}
