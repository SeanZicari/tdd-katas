package string_calculator_kata

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd_EmptyStringReturnsZero(t *testing.T) {
	assert := assert2.New(t)

	res, _ := Add("")

	assert.Equal(0, res)
}

func TestAdd_OneNumberReturnsSameNumber(t *testing.T) {
	assert := assert2.New(t)

	res, _ := Add("1")

	assert.Equal(1, res)
}

func TestAdd_HandlesAnyAmountOfNumbers(t *testing.T) {
	assert := assert2.New(t)

	res, _ := Add("1,2,3,4,5")

	assert.Equal(15, res)
}

func TestAdd_CanUseNewlineAsSeparator(t *testing.T) {
	assert := assert2.New(t)

	res, _ := Add("1\n2,3,4,5")

	assert.Equal(15, res)
}

func TestAdd_TwoNumbersReturnsSum(t *testing.T) {
	assert := assert2.New(t)

	res, _ := Add("1,2")

	assert.Equal(3, res)
}

func TestAdd_CustomDelimiterSupported(t *testing.T) {
	assert := assert2.New(t)

	res, _ := Add("//;\n1;2")

	assert.Equal(3, res)
}

func TestAdd_NegativeNumberReturnsError(t *testing.T) {
	assert := assert2.New(t)

	_, err := Add("-1,2")

	assert.NotNil(err)
	//assert.Equal("negative numbers not allowed: -1", err.Error())
}
