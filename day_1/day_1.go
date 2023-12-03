package main

import (
	file_utils "aoc2023"
	"strings"
	"unicode"
)

func getSingle2DigitNumberFromFirstAndLastDigit(input string) int {
	var i, j int
	twoDigitNumber := 0

	for i = 0; i < len(input); i++ {
		inputChar := rune(input[i])
		if unicode.IsDigit(inputChar) {
			twoDigitNumber += int(inputChar-'0') * 10
			break
		}
	}

	for j = len(input) - 1; j >= i; j-- {
		inputChar := rune(input[j])
		if unicode.IsDigit(inputChar) {
			twoDigitNumber += int(inputChar - '0')
			break
		}
	}

	return twoDigitNumber
}

func sumOfCollaborationValues(lines []string) int {
	totalCollaborationValue := 0

	for _, eachLine := range lines {
		totalCollaborationValue += getSingle2DigitNumberFromFirstAndLastDigit(eachLine)
	}

	return totalCollaborationValue
}

var wordConsideredAsDigits []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func getFirstOccurrenceWordDigit(input string) int {
	digit := -1
	lowestIdx := len(input)
	for i, eachNum := range wordConsideredAsDigits {
		idx := strings.Index(input, eachNum)
		if idx > -1 {
			if idx < lowestIdx {
				lowestIdx = idx
				digit = i + 1
			}
		}
		if lowestIdx == 0 {
			break
		}
	}
	return digit
}

func getLastOccurrenceWordDigit(input string) int {
	digit := -1
	highestIdx := -1
	for i, eachNum := range wordConsideredAsDigits {
		idx := strings.LastIndex(input, eachNum)
		if idx > -1 {
			if idx > highestIdx {
				highestIdx = idx
				digit = i + 1
			}
		}
		if highestIdx == len(input)-1 {
			break
		}
	}
	return digit
}

func getCorrectedSingle2DigitNumberFromFirstAndLastDigit(input string) int {
	var i, j int
	twoDigitNumber := 0

	for i = 0; i < len(input); i++ {
		inputChar := rune(input[i])
		if unicode.IsDigit(inputChar) {
			digit := getFirstOccurrenceWordDigit(input[0:i])
			if digit != -1 {
				twoDigitNumber += digit * 10
			} else {
				twoDigitNumber += int(inputChar-'0') * 10
			}
			break
		}
	}

	// found 1st digit
	if twoDigitNumber > 9 {
		for j = len(input) - 1; j >= i; j-- {
			inputChar := rune(input[j])
			if unicode.IsDigit(inputChar) {
				digit := getLastOccurrenceWordDigit(input[j:])
				if digit != -1 {
					twoDigitNumber += digit
				} else {
					twoDigitNumber += int(inputChar - '0')
				}
				break
			}
		}
	} else {
		firstDigit := getFirstOccurrenceWordDigit(input)
		lastDigit := getLastOccurrenceWordDigit(input)
		twoDigitNumber += firstDigit*10 + lastDigit
	}

	return twoDigitNumber
}

func correctedSumOfCollaborationValues(lines []string) int {
	totalCollaborationValue := 0

	for _, eachLine := range lines {
		collaborationValue := getCorrectedSingle2DigitNumberFromFirstAndLastDigit(eachLine)
		totalCollaborationValue += collaborationValue
	}

	return totalCollaborationValue
}

func main() {
	lines := file_utils.ReadFile("./input.txt")

	println("sumOfCollaborationValues :- ", sumOfCollaborationValues(lines))

	println("correctedSumOfCollaborationValues :- ", correctedSumOfCollaborationValues(lines))
}
