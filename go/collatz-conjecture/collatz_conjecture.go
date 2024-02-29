package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("invalid number %d", n)
	}

	var steps int
	for n != 1 {
		if n%2 == 0 {
			n /= 2
			steps++
			continue
		}
		n *= 3
		n++
		steps++

	}
	return steps, nil
}
