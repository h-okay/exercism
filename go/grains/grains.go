package grains

import "fmt"

// Square calculates the number of grains on a specific square of a chessboard.
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("square doesn't exist")
	}
	return 1 << uint64(number-1), nil
}

// Total calculates the total number of grains on the chessboard.
func Total() uint64 {
	return (1 << 64) - 1
}
