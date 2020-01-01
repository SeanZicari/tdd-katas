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
	var negatives []string
	var err error

	sum := 0
	for _, numStr := range nums {
		num, _ := strconv.Atoi(numStr)
		if num < 0 {
			negatives = append(negatives, numStr)
		}
		if num > 1000 {
			continue
		}
		sum += num
	}

	if len(negatives) > 0 {
		err = fmt.Errorf("negative numbers not allowed: %s", strings.Join(negatives, ", "))
	}

	return sum, err
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
