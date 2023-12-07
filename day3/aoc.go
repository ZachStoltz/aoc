package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day3/data-test")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	dirs := [][]int{{-1, 0}, {-1,-1}, {0, -1}, {1,-1}, {0, 1}, {1, 1}, {1, 0}, {-1, 1}}
	scanner := bufio.NewScanner(file)
	digit := solveDayThreePartOne([][]string{}, dirs, scanner.Scan(), scanner)
	finalPrint := fmt.Sprintf("Part One %d\n", digit)

	fmt.Print(finalPrint)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}

func solveDayThreePartOne(matrix [][]string, dirs [][]int, hasMoreLines bool, scanner *bufio.Scanner) int {
	if (hasMoreLines) {
		split := strings.Split(scanner.Text(), "")

		fmt.Println(split)

		matrix = append(matrix, split)

		return solveDayThreePartOne(matrix, dirs, scanner.Scan(), scanner)
	}

	sum := 0

	for y, row := range matrix {
		isNumber := false;
    currentNumber := "";
    check := true

		for x, char := range row {
			isNumber = isCharNumber(char)

			if (!isNumber && !check) {
				fmt.Println(fmt.Sprintf("current Number -> %s", currentNumber))

				curr, _ := strconv.Atoi(currentNumber)

				sum += curr
			}

			if (!isNumber) {
				currentNumber = ""
				check = true
			}

			if (isNumber && check) {
				doesPass := false

				for _, dir := range dirs {
					char, hasValue := get(x, y, dir, matrix)

					doesPass = doesPass || hasValue && !isDot(char) && !isCharNumber(char)
				}

				if (doesPass) {
					check = false
				}
			}

			if (isNumber) {
				charNum, _ := get(x, y, []int{0,0}, matrix)
				currentNumber += charNum
			}
		}
		if (isNumber && !check) {
			fmt.Println(fmt.Sprintf("current Number -> %s", currentNumber))

			curr, _ := strconv.Atoi(currentNumber)

			sum += curr
		}
	}

	return sum;
}

func partTwo() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day3/data-p2")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	dirs := [][]int{{-1, 0}, {-1,-1}, {0, -1}, {1,-1}, {0, 1}, {1, 1}, {1, 0}, {-1, 1}}
	scanner := bufio.NewScanner(file)
	digit := solveDayThreePartTwo([][]string{}, dirs, scanner.Scan(), scanner)
	finalPrint := fmt.Sprintf("Part Two %d\n", digit)

	fmt.Print(finalPrint)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}

func solveDayThreePartTwo(contents [][]string, dirs [][]int, hasMoreLines bool, scanner *bufio.Scanner) int {
	if (hasMoreLines) {
		split := strings.Split(scanner.Text(), "")

		fmt.Println(split)

		contents = append(contents, split)

		return solveDayThreePartTwo(contents, dirs, scanner.Scan(), scanner)
	}

	sum := 0

	for y, row := range contents {
		isGear := false;

		for x, char := range row {
			isGear = isCharGear(char)

			if (isGear) {
				gearVals := []int{}

				currentSearchedIndiciesByRowIdx := make(map[int][]int)

				for _, dir := range dirs {
					dy := y + dir[1]
					dx := x + dir[0]
					dLeftx := dx - 1
					dRightx := dx + 1
					shouldContinue := false

					checkForRow := currentSearchedIndiciesByRowIdx[dy]

					for _, val := range checkForRow {
						if (val == dx) {
							shouldContinue = true
							break
						}
					}

					if (shouldContinue) {
						continue
					}

					pChar, isValid := get(x, y, dir, contents)

					if (!isValid || isDot(pChar)) { continue }

					isNumber := isCharNumber(pChar)

					if (!isNumber) { continue	}

					currentSearchedIndiciesByRowIdx[dy] = append(checkForRow, dy)

					potentialNumberChars := pChar

					potentialNumberChars = gatherLeftNumberChars(potentialNumberChars, dLeftx, x, dy, contents, currentSearchedIndiciesByRowIdx)

					fmt.Println(fmt.Sprintf("potentialNumberChars BEFORE gatherRight -> %s", potentialNumberChars))

					potentialNumberChars = gatherRightNumberChars(potentialNumberChars, dRightx, x, dy, contents, currentSearchedIndiciesByRowIdx)

					fmt.Println(fmt.Sprintf("potentialNumberChars -> %s", potentialNumberChars))

					num, _ := strconv.Atoi(potentialNumberChars)

					gearVals = append(gearVals, num)
				}

				gearVal := 1

				if (len(gearVals) == 1) {
					continue
				}

				for _, val := range gearVals {
					gearVal *= val
				}

				sum += gearVal
			}
		}
	}

	return sum;
}

func gatherLeftNumberChars(thingToMutate string, iterator int, x int, y int, contents [][]string, currentSearchedIndiciesByRowIdx map[int][]int) string {
	row := contents[y]

	if (iterator < 0) {
		return thingToMutate
	}

	currentSearchedIndiciesByRowIdx[y] = append(currentSearchedIndiciesByRowIdx[y], iterator)

	char := row[iterator]

	if (isDot(char) || isCharGear(char)) {
		return thingToMutate
	}

	thingToMutate = char + thingToMutate
	iterator -= 1

	return gatherLeftNumberChars(thingToMutate, iterator, x, y, contents, currentSearchedIndiciesByRowIdx)
}

func gatherRightNumberChars(thingToMutate string, iterator int, x int, y int, contents [][]string, currentSearchedIndiciesByRowIdx map[int][]int) string {
	row := contents[y]

	if (iterator > len(row) - 1) {
		return thingToMutate
	}

	currentSearchedIndiciesByRowIdx[y] = append(currentSearchedIndiciesByRowIdx[y], iterator)


	char, _ := get(iterator, y, []int{0,0}, contents)

	if (isDot(char) || isCharGear((char))) {
		return thingToMutate
	}

	thingToMutate += char
	iterator += 1

	return gatherRightNumberChars(thingToMutate, iterator, x, y, contents, currentSearchedIndiciesByRowIdx)
}

func isDot(char string) bool {
	return char == "."
}

func isCharNumber(char string) bool {
	_, err := strconv.Atoi(char)

	return err == nil
}

func isCharGear(char string) bool {
	return char == "*"
}

func get(x int, y int, dir []int, contents [][]string) (string,bool) {
	xDir := dir[0]
	yDir := dir[1]
	dy := y + yDir
	dx := x + xDir

	if (dy < 0 || dy > len(contents) - 1) {
		return "" , false
	}

	chars := contents[dy]

	if (chars == nil) {
		return "", false
	}

	if (dx < 0 || dx > len(chars) - 1) {
		return "", false
	}

	return chars[dx], true;
}


