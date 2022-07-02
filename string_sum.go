package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

const customErr = "error is: %w"

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func inputSplit(input string) (operands []string) {
	var str string
	for i, r := range input {
		if i != 0 && (r == '-' || r == '+') {
			operands = append(operands, str)
			str = string(r)
		} else {
			str += string(r)
		}
	}
	operands = append(operands, str)
	return
}

func StringSum(input string) (output string, err error) {

	newInput := strings.ReplaceAll(input, " ", "")
	if len(newInput) == 0 {
		return "", fmt.Errorf(customErr, errorEmptyInput)
	}
	operands := inputSplit(newInput)

	if len(operands) != 2 {
		return "", fmt.Errorf(customErr, errorNotTwoOperands)
	}
	firstNum, err := strconv.Atoi(operands[0])

	if err != nil {
		return "", fmt.Errorf(customErr, err)
	}

	secondNum, err := strconv.Atoi(operands[1])

	if err != nil {
		return "", fmt.Errorf(customErr, err)
	}
	output = strconv.Itoa(firstNum + secondNum)
	return
}
