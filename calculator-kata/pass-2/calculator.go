package calculator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Add(s string) (int, error) {
	var negatives []string
	sum := 0
	delimiters := []string{","}

	for _, line := range strings.Split(s, "\n") {
		if strings.HasPrefix(line, "//") {
			delimiters = deriveCustomDelimiters(line)
			continue
		}
		negatives, sum = accumulate(delimiters, line, negatives, sum)
	}

	return sum, errIfNegatives(negatives)
}

func accumulate(delimiters []string, line string, negatives []string, sum int) ([]string, int) {
	for _, numStr := range regexp.MustCompile(strings.Join(delimiters, "|")).Split(line, -1) {
		num, _ := strconv.Atoi(numStr)
		if num > 1000 {
			continue
		}
		if num < 0 {
			negatives = append(negatives, numStr)
			continue
		}
		sum += num
	}
	return negatives, sum
}

func errIfNegatives(negatives []string) error {
	var err error

	if len(negatives) > 0 {
		err = fmt.Errorf("negatives not allowed: %s", strings.Join(negatives, ", "))
	}

	return err
}

func deriveCustomDelimiters(line string) []string {
	var delimiters []string

	for _, match := range regexp.MustCompile(`[^/\[\]]+`).FindAllString(line, -1) {
		delimiters = append(delimiters, regexp.QuoteMeta(match))
	}

	return delimiters
}