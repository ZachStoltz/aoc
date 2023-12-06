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
	partTwo()
}

func partOne() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day2/data-p1")
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
	redMax := 12
	greenMax := 13
	blueMax := 14

	if (!hasMoreLines) {
		return sum
	}

	text := scanner.Text()
	gameId, passes := validGame(text, redMax, greenMax, blueMax)
	nextSum := sum

	if (passes) {
		nextSum = sum + gameId
	}

	return solveDayTwoPartOne(nextSum, scanner.Scan(), scanner);
}

func validGame(text string, redMax int, greenMax int, blueMax int) (int, bool) {
	pattern := regexp.MustCompile(`Game (\d+):\s`)
	match := pattern.FindStringSubmatch(text)
	captureGroups := match[1:]
	gameId, err := strconv.Atoi(captureGroups[0])
	subText := pattern.ReplaceAllString(text, "")
	groups := strings.Split(subText, ";")
	digitAndColorPattern := regexp.MustCompile(`(\d+)\s+(red|blue|green)`)
	validGame := true

	for _, group := range groups {
		matches := digitAndColorPattern.FindAllStringSubmatch(group, -1)

		if len(matches) > 0 {
			for _, match := range matches {
				digit, _ := strconv.Atoi(match[1])
				color := match[2]

				if (color == "red") {
					validGame = digit <= redMax
				}
				if (color == "green") {
					validGame = digit <= greenMax
				}
				if (color == "blue") {
					validGame = digit <= blueMax
				}

				if (!validGame) { break;}
			}
		}

		if (!validGame) { break; }
	}

	return gameId, err == nil && validGame
}

func partTwo() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day2/data-p2")
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
	digitAndColorPattern := regexp.MustCompile(`\d+\s(?:blue|green|red)`)

	text := scanner.Text()
	rollupByColor := make(map[string]int)

	groups := digitAndColorPattern.FindAllStringSubmatch(text, -1)

	for _, group := range groups {
		split := strings.Split(group[0], " ")
		digit, _ := strconv.Atoi(split[0])
		color:= split[1]

		if (rollupByColor[color] == 0 || digit > rollupByColor[color]) {
			rollupByColor[color] = digit
		}
	}

	power := 1

	for _, total := range rollupByColor {
		power *= total
	}

	return solveDayTwoPartTwo(sum + power, scanner.Scan(), scanner)
}