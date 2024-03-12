package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	output := make(map[string]int)
	for point, arr := range in {
		for _, value := range arr {
			output[strings.ToLower(value)] = point
		}
	}
	return output
}
