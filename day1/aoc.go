package main

import (
	"bufio"
	"errors"
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

	text := scanner.Text()

	fmt.Println(text)

	chars := strings.Split(text, "")
	firstNumString := getFirstNum(0, 0, "", "", chars)
	lastNumStartPos := len(chars) - 1
	lastNumString := getLastNum(lastNumStartPos, lastNumStartPos, "", "", chars)

	// NOTE: Def should handle error.... oh well
	digit, _ := getDigit(firstNumString + lastNumString)

	info := fmt.Sprintf("first %s <-> second %s [%d] \n", firstNumString, lastNumString, digit)

	fmt.Print(info)

	return solveDayOne(value + digit, scanner.Scan(), scanner)
}

func getDigitFromString(word string) (int, error) {
	switch word {
		case "one": return 1, nil;
		case "two": return 2, nil;
		case "three": return 3,nil;
		case "four": return 4,nil;
		case "five": return 5, nil;
		case "six": return 6, nil;
		case "seven": return 7, nil;
		case "eight": return 8, nil;
		case "nine": return 9, nil;
		default: return 0, errors.New("Fail");
	}
}

func getDigit(char string) (int, error) {
	return strconv.Atoi(char)
}

func getFirstNum(idx int, currPos int, currValue string, value string, chars []string) string {
	maxLengthOfCharacterCount := 5

	// GoLang N00B.... should prob return tuple int, error
	if currPos >= len(chars) {
		return value
	}

	char := chars[currPos]
	str := currValue + char
	digit, err := getDigit(char)
	dig2, err2 := getDigitFromString(str)
	nextIdx := currPos + 1

	if (err == nil && len(str) < maxLengthOfCharacterCount) {
		return char
	}

	if (err2 == nil) { return strconv.Itoa(dig2) }

	if (len(str) >= maxLengthOfCharacterCount) {
		strt := idx + 1

		if (err == nil) {
			return getFirstNum(strt, strt, value, char, chars)
		}


		return getFirstNum(strt, strt, "", value, chars)
	}


	if (err == nil) { return getFirstNum(idx, nextIdx, str, strconv.Itoa(digit) , chars) }

	return getFirstNum(idx, nextIdx, str, value, chars)
}

func getLastNum(idx int, currPos int, currValue string, value string, chars []string) string {
	maxLengthOfCharacterCount := 5

	char := chars[currPos]
	str :=  char + currValue
	digit, err := getDigit(char)
	nextIdx := currPos - 1
	dig2, err2 := getDigitFromString(str)

	if (err == nil && len(str) < maxLengthOfCharacterCount) {
		return char
	}

	if (err2 == nil) { return strconv.Itoa(dig2) }

	if (len(str) >= maxLengthOfCharacterCount) {
		strt := idx - 1

		if (err == nil) {
			return getLastNum(strt, strt, value, char, chars)
		}

		return getLastNum(strt, strt, "", value, chars)
	}

	if (err == nil) {
		return getLastNum(idx, nextIdx, str, strconv.Itoa(digit) , chars)
	}

	return getLastNum(idx, nextIdx, str, value, chars)
}

func main() {
    file, err := os.Open("/Users/zachstoltz/develop/aoc/day1/data-p2")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

		digit := solveDayOne(0, scanner.Scan(), scanner)

		finalPrint := fmt.Sprintf("Final Number %d\n", digit)

		fmt.Print(finalPrint)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}