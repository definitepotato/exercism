package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	Err  error
	Cows int
}

func (ice *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %v", ice.Cows, ice.Err)
}

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	fodderAmount, err := fc.FodderAmount(cows)
	if err != nil {
		return 0.0, err
	}

	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	result := float64(fodderAmount / float64(cows) * fatteningFactor)
	return result, err
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		result, err := DivideFood(fc, cows)
		return result, err
	}

	return 0.0, errors.New("invalid number of cows")
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		ice := &InvalidCowsError{
			Cows: cows,
			Err:  errors.New("there are no negative cows"),
		}
		return ice
	}

	if cows == 0 {
		ice := &InvalidCowsError{
			Cows: cows,
			Err:  errors.New("no cows don't need food"),
		}
		return ice
	}

	return nil
}
