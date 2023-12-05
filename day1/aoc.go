package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)



func solveDayOne(value int, hasMoreLines bool, scanner *bufio.Scanner) int {
	if (!hasMoreLines) {
		return value
	}

	var chars = strings.Split(scanner.Text(), "")
	var firstNumString = getFirstNum(0, "", chars)
	var lastNumString = getLastNum(len(chars) - 1, "", chars)

	// NOTE: Def should handle error.... oh well
	digit, _ := getDigit(firstNumString + lastNumString)

	fmt.Printf("first num %s <-> second num %s num [%d] \n", firstNumString, lastNumString, digit)

	return solveDayOne(value + digit, scanner.Scan(), scanner)
}

func getDigit(char string) (int, error) {
	return strconv.Atoi(char)
}

func getFirstNum(idx int, value string, chars []string) string {
	// GoLang N00B.... should prob return tuple int, error
	if idx > len(chars) {
		return value
	}

	str := chars[idx]

	_, err := getDigit(str)

	if (err != nil) {
		i := idx + 1
		return getFirstNum(i, "", chars)
	}

	return str
}

func getLastNum(idx int, value string, chars []string) string {
	// GoLang N00B.... should prob return tuple int, error
	if idx < 0 {
		return value
	}

	str := chars[idx]

	_, err := getDigit(str)

	if (err != nil) {
		i := idx - 1
		return getLastNum(i, "", chars)
	}

	return str
}

func main() {
    file, err := os.Open("/Users/zachstoltz/develop/aoc/day1/data-test")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

		digit := solveDayOne(0, scanner.Scan(), scanner)

		fmt.Printf("Final Number %d\n", digit)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}