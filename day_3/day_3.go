package main

import (
	file_utils "aoc2023"
	"regexp"
	"strconv"
	"unicode"
)

type ConnectedNumber struct {
	num         int
	isConnected bool
	startIdx    int
	endIdx      int
}

type ConnectedSymbol struct {
	idx int
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

// func extractNumbers(s string) ([]int, error) {
// 	var numbers []int
// 	re := regexp.MustCompile(`\d+`) // Regular expression to match sequences of digits

// 	matches := re.FindAllStringIndex(s, -1) // Find all matches
// 	for _, match := range matches {
// 		num, err := strconv.Atoi(s[match[0]:match[1]])
// 		if err != nil {
// 			return nil, err // Handle conversion error
// 		}
// 		numbers = append(numbers, num)
// 	}

// 	return numbers, nil
// }

func parseAndGetConnectedNumbers(line string) []ConnectedNumber {
	var listOfNum []ConnectedNumber

	re := regexp.MustCompile(`\d+`) // Regular expression to match sequences of digits

	matches := re.FindAllStringIndex(line, -1) // Find all matches
	for _, match := range matches {
		start, end := match[0], match[1]
		num, _ := strconv.Atoi(line[start:end])
		isConnected := false
		if start-1 > 0 {
			isConnected = isSymbol(rune(line[start-1]))
		}
		if end < len(line) {
			isConnected = isConnected || isSymbol(rune(line[end-1]))
		}
		listOfNum = append(listOfNum, ConnectedNumber{num: num, startIdx: start, endIdx: end - 1, isConnected: isConnected})
	}

	return listOfNum
}

func parseAndGetConnectedSymbols(line string) []ConnectedSymbol {
	var listOfSymbol []ConnectedSymbol

	for i := 0; i < len(line); i++ {
		if isSymbol(rune(line[i])) {
			listOfSymbol = append(listOfSymbol, ConnectedSymbol{idx: i})
		}
	}

	return listOfSymbol
}

func parseAndCheckConnected(listOfPrevNum *[]ConnectedNumber, listOfPrevSymbol *[]ConnectedSymbol, currLine string) ([]ConnectedNumber, []ConnectedSymbol) {
	var listOfCurrNum []ConnectedNumber = parseAndGetConnectedNumbers(currLine)
	var listOfCurrSymbol []ConnectedSymbol = parseAndGetConnectedSymbols(currLine)

	for i, eachNum := range listOfCurrNum {
		for _, eachSymbol := range listOfCurrSymbol {
			if eachSymbol.idx >= eachNum.startIdx-1 && eachSymbol.idx <= eachNum.endIdx+1 {
				listOfCurrNum[i].isConnected = true
				break
			}
		}
		for _, eachSymbol := range *listOfPrevSymbol {
			if eachSymbol.idx >= eachNum.startIdx-1 && eachSymbol.idx <= eachNum.endIdx+1 {
				listOfCurrNum[i].isConnected = true
				break
			}
		}
	}

	for i, eachNum := range *listOfPrevNum {
		for _, eachSymbol := range listOfCurrSymbol {
			if eachSymbol.idx >= eachNum.startIdx-1 && eachSymbol.idx <= eachNum.endIdx+1 {
				(*listOfPrevNum)[i].isConnected = true
				break
			}
		}
	}

	return listOfCurrNum, listOfCurrSymbol
}

func partOne(lines []string) int {
	output := 0

	allNums := []ConnectedNumber{}
	prevNums := []ConnectedNumber{}
	prevSymbols := []ConnectedSymbol{}

	for _, eachLine := range lines {
		newNums, newSymbols := parseAndCheckConnected(&prevNums, &prevSymbols, eachLine)
		allNums = append(allNums, prevNums...)
		for _, eachNum := range prevNums {
			println(eachNum.num, eachNum.isConnected, eachNum.startIdx, eachNum.endIdx)
		}
		prevNums, prevSymbols = newNums, newSymbols
		for _, eachSymb := range prevSymbols {
			println(eachSymb.idx)
		}
	}

	allNums = append(allNums, prevNums...)

	for _, eachNum := range allNums {
		// println(eachNum.num, eachNum.isConnected, eachNum.startIdx, eachNum.endIdx)
		if eachNum.isConnected {
			output += eachNum.num
		}
	}

	return output
}

func partTwo(lines []string) int {
	output := 0

	return output
}

func main() {
	lines := file_utils.ReadFile("input.txt")

	println("partOne :- ", partOne(lines))
	// println("partTwo :- ", partTwo(lines))

	// for _, e := range lines {
	// 	println(e)
	// 	println("nums")
	// 	for _, i := range parseAndGetConnectedNumbers(e) {
	// 		println(i.num)
	// 	}
	// 	println("syms")
	// 	for _, i := range parseAndGetConnectedSymbols(e) {
	// 		println(i.idx)
	// 	}
	// 	println()
	// }

}
