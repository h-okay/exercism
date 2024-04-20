package lsproduct

import (
	"errors"
	"fmt"
	"slices"
	"unicode"
)

var ErrInvalidInput = errors.New("digits input must only contain digits")
var ErrNegativeSpan = errors.New("span must not be negative")

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span <= 0 {
		return 0, ErrNegativeSpan
	}

	chunks, err := getChunks(digits, span)
	if err != nil {
		return 0, err
	}
	fmt.Println(chunks)

	products := make([]int64, 0, len(chunks))
	for _, c := range chunks {
		products = append(products, getProduct(c))
	}
	fmt.Println(products)

	if len(products) > 0 {
		return slices.Max(products), nil
	}

	return 0, errors.New("something went wrong")
}

func getChunks(s string, n int) ([]string, error) {
	chunks := make([]string, 0)
	i, j := 0, 0
	for i < len(s)-1 {
		if !unicode.IsDigit(rune(s[i])) {
			return nil, ErrInvalidInput
		}
		if (j+1)%n == 0 {
			chunks = append(chunks, s[i:j])
			i++
			j = i + 1
			continue
		}
		j++
	}

	return chunks, nil
}

func getProduct(c string) int64 {
	var product int64 = 1
	for _, c := range c {
		product *= int64(c - '0')
	}
	return product
}
