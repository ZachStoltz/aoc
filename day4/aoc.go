package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// partOne()
	// partTwo()
}

func partOne() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day4/test")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	digit := solveDayFourPartOne([][]string{}, scanner.Scan(), scanner)
	finalPrint := fmt.Sprintf("Part One %d\n", digit)

	fmt.Print(finalPrint)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}

func solveDayFourPartOne(matrix [][]string, hasMoreLines bool, scanner *bufio.Scanner) int {
	if (hasMoreLines) {
		split := strings.Split(scanner.Text(), "")

		matrix = append(matrix, split)

		return solveDayFourPartOne(matrix, scanner.Scan(), scanner)
	}

	return 0;
}

func partTwo() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day4/p2")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	digit := solveDayFourPartTwo([][]string{}, scanner.Scan(), scanner)
	finalPrint := fmt.Sprintf("Part Two %d\n", digit)

	fmt.Print(finalPrint)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}

func solveDayFourPartTwo(contents [][]string, hasMoreLines bool, scanner *bufio.Scanner) int {
	if (hasMoreLines) {
		split := strings.Split(scanner.Text(), "")

		fmt.Println(split)

		contents = append(contents, split)

		return solveDayFourPartTwo(contents, scanner.Scan(), scanner)
	}

	return 0
}
