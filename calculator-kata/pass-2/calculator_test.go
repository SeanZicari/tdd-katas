package calculator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestAdd_EmptyStringReturnsZero(t *testing.T) {
	wanted := 0

	if got, _ := Add(""); got != wanted {
		t.Errorf("Got %d but wanted %d", got, wanted)
	}
}

func TestAdd_SingleNumberReturnsSameNumber(t *testing.T) {
	wanted := 13

	if got, _ := Add("13"); got != wanted {
		t.Errorf("Got %d but wanted %d", got, wanted)
	}
}

func TestAdd_TwoNumbersReturnsSum(t *testing.T) {
	wanted := 20

	if got, _ := Add("13,7"); got != wanted {
		t.Errorf("Got %d but wanted %d", got, wanted)
	}
}

func TestAdd_MultipleNumbersReturnsSum(t *testing.T) {
	wanted := 400

	if got, _ := Add("13,7,4,6,2,2,2,2,2,360"); got != wanted {
		t.Errorf("Got %d but wanted %d", got, wanted)
	}
}

func TestAdd_NewlinesOk(t *testing.T) {
	wanted := 21

	if got, _ := Add("1\n2,3\n4,5,6"); got != wanted {
		t.Errorf("Got %d but wanted %d", got, wanted)
	}
}

func TestAdd_CustomDelimiterOk(t *testing.T) {
	wanted := 3

	if got, _ := Add("//;\n1;2"); got != wanted {
		t.Errorf("Got %d but wanted %d", got, wanted)
	}
}

func TestAdd_NegativeNumberProducesError(t *testing.T) {
	wanted := errors.New("negatives not allowed: -1")

	_, err := Add("-1,2")

	if got := err.Error(); got != wanted.Error() {
		t.Errorf("Got %s but wanted %s", got, wanted)
	}
}

func TestAdd_AllNegativeNumbersReturnedInErrorMsg(t *testing.T) {
	wanted := errors.New("negatives not allowed: -1, -3")

	_, err := Add("-1,2,-3")

	if got := err.Error(); got != wanted.Error() {
		t.Errorf("Got %s but wanted %s", got, wanted)
	}
}

func TestAdd_LargeNumbersIgnored(t *testing.T) {
	wanted := 4

	if got, _ := Add("2,1001,2"); got != wanted {
		t.Errorf("Got %d but wanted %d", got, wanted)
	}
}

func TestAdd_CustomSizeDelimiterOk(t *testing.T) {
	wanted := 4

	if got, _ := Add("//[***]\n2***1001***2"); got != wanted {
		t.Errorf("Got %d but wanted %d", got, wanted)
	}
}

func TestAdd_MultipleCustomDelimitersOk(t *testing.T) {
	wanted := 4

	if got, _ := Add("//[*--*][_]\n2*--*1001_2"); got != wanted {
		t.Errorf("Got %d but wanted %d", got, wanted)
	}
}

func Add(s string) (int, error) {
	var err error
	var negatives []string
	var delimiters []string
	sum := 0

	for _, line := range strings.Split(s, "\n") {
		if strings.HasPrefix(line, "//") {
			for _, match := range regexp.MustCompile(`[^/\[\]]+`).FindAllString(line, -1) {
				delimiters = append(delimiters, regexp.QuoteMeta(match))
			}
			//continue  - Test passes without this, but line is later split for no reason
		} else {
			delimiters = append(delimiters, ",")
		}
		for _, numStr := range regexp.MustCompile(strings.Join(delimiters, "|")).Split(line, -1) {
			num, _ := strconv.Atoi(numStr)
			if num > 1000 {
				continue
			}
			if num < 0 {
				negatives = append(negatives, numStr)
				//continue  - Works without but makes no sense to sum when we know we've failed
			}
			sum += num
		}
	}

	if len(negatives) > 0 {
		err = fmt.Errorf("negatives not allowed: %s", strings.Join(negatives, ", "))
	}

	return sum, err
}