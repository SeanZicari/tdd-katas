package calculator

import (
	"errors"
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

