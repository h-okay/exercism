package romannumerals

import (
	"bytes"
	"fmt"
)

type Conversion struct {
	value int
	roman string
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", fmt.Errorf("number %d is not valid", input)
	}

	lookup := []Conversion{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	r := bytes.NewBufferString("")
	for _, conversion := range lookup {
		for input >= conversion.value {
			r.WriteString(conversion.roman)
			input -= conversion.value
		}
	}

	return r.String(), nil

}
