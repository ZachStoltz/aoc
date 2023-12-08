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
	partOne()
	// partTwo()
}

func partOne() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day4/p1")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	digit := solveDayFourPartOne(0, scanner.Scan(), scanner)
	finalPrint := fmt.Sprintf("Part One %d\n", digit)

	fmt.Print(finalPrint)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}

func solveDayFourPartOne(sum int, hasMoreLines bool, scanner *bufio.Scanner) int {
	if (!hasMoreLines) { return sum }

		text := scanner.Text()
		text = regexp.MustCompile	(`\s\s`).ReplaceAllString(text, " ")
		text = regexp.MustCompile(`Card\s\d+:\s`).ReplaceAllString(text, "")
		split := strings.Split(text, " | ")
		card := make(map[int]bool)
		cardAmt := 0

		for _, val := range strings.Split(split[0], " ") {
			num, _ := strconv.Atoi(val)

			card[num] = true
		}

		for _, val := range strings.Split(split[1], " ") {
			num, _ := strconv.Atoi(val)

			if card[num] {
				if cardAmt <= 1 {
					cardAmt += 1
					continue
				}

				cardAmt *= 2
			}
		}

		sum += cardAmt

		return solveDayFourPartOne(sum, scanner.Scan(), scanner)
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
