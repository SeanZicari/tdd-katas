package string_calculator_kata

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Add(numbers string) (int, error) {
	delimiter, numbers := deriveDelimiters(numbers)
	return sumNumbers(delimiter.Split(numbers, -1))
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

	return sum, errIfNegatives(negatives, err)
}

func errIfNegatives(negatives []string, err error) error {
	if len(negatives) > 0 {
		err = fmt.Errorf("negative numbers not allowed: %s", strings.Join(negatives, ", "))
	}
	return err
}

func deriveDelimiters(numbers string) (*regexp.Regexp, string) {
	var delimiters []string

	if strings.HasPrefix(numbers, "//") {
		pieces := strings.Split(numbers, "\n")

		re := regexp.MustCompile(`[^/\[\]]+`)
		delimiters = re.FindAllString(pieces[0], -1)
		numbers = pieces[1]
	} else {
		delimiters = []string{"\n", ","}
	}

	return regexp.MustCompile(strings.Join(escapeRegexSpecialChars(delimiters), "|")), numbers
}

func escapeRegexSpecialChars(delimiters []string) []string {
	var escapedDelimiters []string

	for _, delim := range delimiters {
		escapedDelimiters = append(escapedDelimiters, regexp.QuoteMeta(delim))
	}

	return escapedDelimiters
}
