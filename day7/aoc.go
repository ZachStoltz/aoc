package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
	bid int
	rank int
	cards []string
}

func getRank(cards []string, cardValues map[string]int) int {
	cardMap := make(map[string]int)
	first := cards[0]

	cardMap[first] = 1
	highCard := ""
	highCardVal := 0

	for _, card := range cards[1:] {
		item, exists := cardMap[card]

		if (exists) {
			cardMap[card] = item + 1
		} else {
			cardMap[card] = 1
		}
	}

	for k, v := range cardMap {
		if k == "J" {
			continue
		}

		if (v > highCardVal) {
			highCard = k
			highCardVal = v
		}
	}

	wildcard, hasWildcard := cardMap["J"]

	if hasWildcard && wildcard == 5 {
		return 6
	}

	score := 0

	for card, total := range cardMap {
		if (card == "J") {
			continue
		}
		if (card == highCard && hasWildcard) {
			total += wildcard
		}

		if (total == 4 || total == 5) {
			return total + 1
		}
	}

	if (score != 0) {
		return score
	}

	isPair := false

	for card, total := range cardMap {
		if (card == "J") {
			continue
		}
		if (card == highCard && hasWildcard) {
			total += wildcard
		}
		if (total == 3 || total == 2) {
			score += total
		}

		if (score == 5) {
			score = 4
			break;
		}

		if (score == 4) {
			score = 2
			isPair = false
			break
		}

		if (score == 2) {
			isPair = true
		}
	}

	if (score > 2) {
		return score
	}

	if (isPair) {
		return 1
	}

	return score
}

func getHand(line string, cardValues map[string]int) *Hand {
	vals := strings.Split(line, " ")
	cardString := vals[0]
	cards := strings.Split(cardString, "")
	bid, _ := strconv.Atoi(vals[1])

	rank := getRank(cards, cardValues)

	hand := Hand{ bid, rank, cards }

	return &hand
}

func getNumValue(card string, cardValues map[string]int) int {
	numVal, exists := cardValues[card]

	if (exists) {
		return numVal
	}

	num, _ := strconv.Atoi(card)

	return num
}

func solve(hasMoreLines bool, scanner *bufio.Scanner, values []*Hand, cardValues map[string]int) int {
	if (hasMoreLines) {
		hand := getHand(scanner.Text(), cardValues)
		values = append(values, hand)

		return solve(scanner.Scan(), scanner, values, cardValues)
	}

	for i := 0; i < len(values) - 1; i++ {
		swapped := false
		for j := 0; j < len(values) - i - 1; j++ {
			curHand := values[j]
			nextHand := values[j + 1]

			if (curHand.rank == nextHand.rank) {
				for num := 0; num < len(curHand.cards); num++ {
					c := curHand.cards[num]
					n := nextHand.cards[num]

					if (c == n) {
						continue;
					}

					cN := getNumValue(c, cardValues)
					nN := getNumValue(n, cardValues)

					if (cN == nN) {
						continue
					}

					if (cN > nN) {
						values[j] = nextHand
						values[j + 1] = curHand
						swapped = true
						break
					}
					break
				}
			}

			if (curHand.rank > nextHand.rank) {
					values[j] = nextHand
					values[j + 1] = curHand
					swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	sum := 0
	for idx, hand := range values {
		sum += ((idx + 1) * hand.bid)

		// if (hand.rank == 1) {
		// 	fmt.Printf("%s %d\n", strings.Join(hand.cards, ""), hand.bid)
		// }

		// fmt.Println(*hand)
	}


	return sum
}

func main() {
	partOneAndTwo()
}

func partOneAndTwo() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day7/p1")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	cardValues := make(map[string]int)

	cardValues["T"] = 10
	cardValues["J"] = 1
	cardValues["Q"] = 12
	cardValues["K"] = 13
	cardValues["A"] = 14

	scanner := bufio.NewScanner(file)
	digit := solve(scanner.Scan(), scanner, []*Hand{}, cardValues)

	fmt.Printf("Result: %d\n", digit)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}
