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
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day3/data-p1")
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

func isDot(char string) bool {
	return char == "."
}

func isCharNumber(char string) bool {
	_, err := strconv.Atoi(char)

	return err == nil
}

func get(x int, y int, dir []int, contents [][]string) (string,bool) {
	xDir := dir[0]
	yDir := dir[1]
	dy := y + yDir
	dx := x + xDir

	// fmt.Println(fmt.Sprintf("xDir %d yDir %d x %d y %d dx %d dy %d", xDir, yDir, x,y,dx,dy))

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


func partTwo() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day3/data-p2")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	digit := solveDayThreePartTwo(0, scanner.Scan(), scanner)

	finalPrint := fmt.Sprintf("Part two %d\n", digit)

	fmt.Print(finalPrint)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}

func solveDayThreePartTwo(sum int, hasMoreLines bool, scanner *bufio.Scanner) int {
	if (!hasMoreLines) {
		return sum
	}

	text := scanner.Text()
	// rollupByColor := make(map[string]int)

	fmt.Println(text)

	return solveDayThreePartTwo(0, scanner.Scan(), scanner)
}