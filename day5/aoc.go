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
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day5/p1")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	digit := solveDayFivePartOne(scanner.Scan(), nil, scanner, []int{}, [][]int{}, [][]int{}, [][]int{}, [][]int{}, [][]int{}, [][]int{}, [][]int{})
	finalPrint := fmt.Sprintf("Part One %d\n", digit)

	fmt.Print(finalPrint)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}
// dest (value) source (key) range
func solveDayFivePartOne(hasMoreLines bool, curProcessType *int, scanner *bufio.Scanner, seeds []int, seedToSoil [][]int, soilToFert [][]int, fertToWater [][]int, waterToLight [][]int, lightToTemp [][]int, tempToHumidity [][]int, humidityToLocation [][]int) int {
	if (hasMoreLines) {
		line := scanner.Text()

		if (len(line) == 0) {
			curProcessType = nil

			return solveDayFivePartOne(scanner.Scan(), curProcessType, scanner, seeds, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation)
		}

		if (curProcessType == nil) {
			seedRowRegex := regexp.MustCompile(`seeds:\s`)
			seedToSoilRegex := regexp.MustCompile(`seed-to-soil map:`)
			soilToFertRegex := regexp.MustCompile(`soil-to-fertilizer map:`)
			fertToWaterRegex := regexp.MustCompile(`fertilizer-to-water map:`)
			waterToLightRegex := regexp.MustCompile(`water-to-light map:`)
			lightToTempRegex := regexp.MustCompile(`light-to-temperature map:`)
			tempToHumidityRegex := regexp.MustCompile(`temperature-to-humidity map:`)
			humidityToLocationRegex := regexp.MustCompile(`humidity-to-location map:`)

			if (seedRowRegex.MatchString(line)) {
				items := strings.Split(regexp.MustCompile(`seeds:\s`).ReplaceAllString(line, ""), " ")

				for _, item := range items {
					seed, _ := strconv.Atoi(item)

					seeds = append(seeds, seed)
				}

				return solveDayFivePartOne(scanner.Scan(), curProcessType, scanner, seeds, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation)
			}

			if (seedToSoilRegex.MatchString(line)) {
				val := 0
				curProcessType = &val

 				return solveDayFivePartOne(scanner.Scan(), curProcessType, scanner, seeds, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation)
			}

			if (soilToFertRegex.MatchString(line)) {
				val := 1
				curProcessType = &val

 				return solveDayFivePartOne(scanner.Scan(), curProcessType, scanner, seeds, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation)
			}

			if (fertToWaterRegex.MatchString(line)) {
				val := 2
				curProcessType = &val

 				return solveDayFivePartOne(scanner.Scan(), curProcessType, scanner, seeds, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation)
			}

			if (waterToLightRegex.MatchString(line)) {
				val := 3
				curProcessType = &val

 				return solveDayFivePartOne(scanner.Scan(), curProcessType, scanner, seeds, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation)
			}

			if (lightToTempRegex.MatchString(line)) {
				val := 4
				curProcessType = &val

 				return solveDayFivePartOne(scanner.Scan(), curProcessType, scanner, seeds, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation)
			}

			if (tempToHumidityRegex.MatchString(line)) {
				val := 5
				curProcessType = &val

 				return solveDayFivePartOne(scanner.Scan(), curProcessType, scanner, seeds, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation)
			}

			if (humidityToLocationRegex.MatchString(line)) {
				val := 6
				curProcessType = &val

 				return solveDayFivePartOne(scanner.Scan(), curProcessType, scanner, seeds, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation)
			}
		}

		switch *curProcessType {
		case 0:
			seedToSoil = processRow(line, seedToSoil)
			break
		case 1:
			soilToFert = processRow(line, soilToFert)
			break
		case 2:
			fertToWater = processRow(line, fertToWater)
			break
		case 3:
			waterToLight = processRow(line, waterToLight)
			break
		case 4:
			lightToTemp = processRow(line, lightToTemp)
			break
		case 5:
			tempToHumidity = processRow(line, tempToHumidity)
			break
		case 6:
			humidityToLocation = processRow(line, humidityToLocation)
			break
		}

		return solveDayFivePartOne(scanner.Scan(), curProcessType, scanner, seeds, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation)
	}

	lowestLocationNumber := -1

	for _, seed := range seeds {
		locationNum := processSeed(seed, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation)

		if (lowestLocationNumber == -1) {
			lowestLocationNumber = locationNum
			continue
		}

		if locationNum < lowestLocationNumber {
			lowestLocationNumber = locationNum
		}
	}

	return lowestLocationNumber
}

func processRow(text string, almanacTypeLookup [][]int) [][]int  {
	parts := strings.Split(text, " ")
	src, _ := strconv.Atoi(parts[1])
	dest, _ := strconv.Atoi(parts[0])
	rangeStop, _ := strconv.Atoi(parts[2])

	return append(almanacTypeLookup, []int{src, dest, rangeStop})
}

func getDestination(input int, lookups [][]int) int {
	out := 0

	for _, lookup := range lookups {
		src := lookup[0]
		dest := lookup[1]
		rangeStop := lookup[2]
		lastSrcNumber := src + rangeStop
		diff := input - src

		fmt.Printf("input %d src %d dest %d range %d\n", input, src, dest, rangeStop)

		if ((input >= src && input <= lastSrcNumber) && diff <= rangeStop) {
			srcDiff := input - src

			out = dest + srcDiff
		}

		if (out == 0) {
			out = input
		}
	}
	return out
}

func processSeed(seed int, seedToSoil [][]int, soilToFert [][]int, fertToWater [][]int, waterToLight [][]int, lightToTemp [][]int, tempToHumidity [][]int, humidityToLocation [][]int) int {
	soil := getDestination(seed, seedToSoil)
	fert := getDestination(soil, soilToFert)
	water := getDestination(fert, fertToWater)
	light := getDestination(water, waterToLight)
	temp := getDestination(light, lightToTemp)
	humidity := getDestination(temp, tempToHumidity)
	loc := getDestination(humidity, humidityToLocation)

	if (loc == 0) {
		loc = seed
	}

	return loc
}



func partTwo() {
	file, err := os.Open("/Users/zachstoltz/develop/aoc/day5/p2")
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
		// text := scanner.Text()


		return solveDayFourPartTwo(contents, scanner.Scan(), scanner)
	}

	return 0;
}
