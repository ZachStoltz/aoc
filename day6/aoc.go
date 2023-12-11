package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	digit := solve(scanner.Scan(), scanner, []int{})
	finalPrint := fmt.Sprintf("Result: %d\n", digit)

	fmt.Print(finalPrint)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}

func solve(hasMoreLines bool, scanner *bufio.Scanner, values []int) int {
	if (hasMoreLines) {
		line := scanner.Text()
		line = regexp.MustCompile(`\w+:.\s+`).ReplaceAllString(line, "")
		line = regexp.MustCompile(`\s+`).ReplaceAllString(line, "")

			num, _ := strconv.Atoi(line)

		values = append(values, num)

		return solve(scanner.Scan(), scanner, values)
	}

	time := values[0]
	record := values[1]

	winners := 0

	for i := 0; i <= time; i++ {
		score := i * (time - i)

		if (score > record) {
			winners += 1
		}
	}

	return winners
}
