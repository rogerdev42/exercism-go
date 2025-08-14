package thefarm

import (
    "errors"
	"fmt"
)

type InvalidCowsError struct {
    message string
    numCows int
}

func (e InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
} 

func DivideFood(fc FodderCalculator, numCows int) (float64, error){
    fa, err := fc.FodderAmount(numCows)
    if err != nil {
        return 0, err
    }
    ff, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    return fa / float64(numCows) * ff, nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error){
	if numCows <= 0 {
        return 0, errors.New("invalid number of cows")
    }
    return DivideFood(fc, numCows)
}

func ValidateNumberOfCows(num int) error {
    switch {
        case num == 0: return &InvalidCowsError{numCows: num, message: "no cows don't need food"}
        case num < 0: return &InvalidCowsError{numCows: num, message: "there are no negative cows"}
        default: return nil
    }
}

