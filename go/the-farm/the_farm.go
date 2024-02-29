package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	amnt, err := fc.FodderAmount(numberOfCows)
	if err != nil {
		return 0, err
	}

	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return amnt * factor / float64(numberOfCows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows > 0 {
		return DivideFood(fc, numberOfCows)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	numberOfCows int
	message      string
}

func (icw InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", icw.numberOfCows, icw.message)
}

func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows < 0 {
		return &InvalidCowsError{
			numberOfCows: numberOfCows,
			message:      "there are no negative cows",
		}
	}

	if numberOfCows == 0 {
		return &InvalidCowsError{
			numberOfCows: numberOfCows,
			message:      "no cows don't need food",
		}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
