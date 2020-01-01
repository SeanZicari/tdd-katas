package string_calculator_kata

import (
	"fmt"
	"strconv"
	"strings"
)

func Add(numbers string) (int, error) {
	delimiters, numbers := deriveDelimiter(numbers)
	numSlice := strings.FieldsFunc(numbers, setDelimiters(delimiters))
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

func deriveDelimiter(numbers string) ([]rune, string) {
	var delimiters []rune

	if strings.HasPrefix(numbers, "//") {
		pieces := strings.Split(numbers, "\n")
		delimiter := rune(pieces[0][len(pieces[0])-1])
		delimiters = []rune{delimiter}
		numbers = pieces[1]
	} else {
		delimiters = []rune{'\n', ','}
	}
	return delimiters, numbers
}

func setDelimiters(delimiters []rune) func(r rune) bool {
	return func(r rune) bool {
		for _, d := range delimiters {
			if r == d {
				return true
			}
		}
		return false
	}
}
