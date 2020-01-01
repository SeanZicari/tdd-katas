package string_calculator_kata

import (
	"fmt"
	"strconv"
	"strings"
)

var delimiters = []rune{'\n', ','}

func Add(numbers string) (int, error) {
	numbers = deriveDelimiter(numbers)
	numSlice := strings.FieldsFunc(numbers, customDelimiter)
	return sumNumbers(numSlice)
}

func sumNumbers(nums []string) (int, error) {
	sum := 0
	for _, numStr := range nums {
		num, _ := strconv.Atoi(numStr)
		if num < 0 {
			return 0, fmt.Errorf("negative numbers not allowed: %d", num)
		}
		sum += num
	}
	return sum, nil
}

func deriveDelimiter(numbers string) string {
	if strings.HasPrefix(numbers, "//") {
		pieces := strings.Split(numbers, "\n")
		delimiter := rune(pieces[0][len(pieces[0])-1])
		delimiters = []rune{delimiter}
		return pieces[1]
	}
	return numbers
}

func customDelimiter(r rune) bool {
	for _, d := range delimiters {
		if r == d {
			return true
		}
	}
	return false
}