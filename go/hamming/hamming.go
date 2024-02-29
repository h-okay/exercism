package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("sequences have different lenthgs, a: %d b: %d", len(a), len(b))
	}

	var count int
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
