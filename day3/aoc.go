package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day3/data-p1")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	digit := solveDayTwoPartOne(0, scanner.Scan(), scanner)

	finalPrint := fmt.Sprintf("Part One %d\n", digit)

	fmt.Print(finalPrint)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}

func solveDayTwoPartOne(sum int, hasMoreLines bool, scanner *bufio.Scanner) int {
	if (!hasMoreLines) {
		return sum
	}

	text := scanner.Text()

	fmt.Println(text)

	return solveDayTwoPartOne(0, scanner.Scan(), scanner);
}

func partTwo() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day3/data-p2")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	digit := solveDayTwoPartTwo(0, scanner.Scan(), scanner)

	finalPrint := fmt.Sprintf("Part two %d\n", digit)

	fmt.Print(finalPrint)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}

func solveDayTwoPartTwo(sum int, hasMoreLines bool, scanner *bufio.Scanner) int {
	if (!hasMoreLines) {
		return sum
	}

	text := scanner.Text()
	// rollupByColor := make(map[string]int)

	fmt.Println(text)

	return solveDayTwoPartTwo(0, scanner.Scan(), scanner)
}