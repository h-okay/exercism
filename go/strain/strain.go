package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[K any](collection []K, predicate func(K) bool) []K {
	var output []K
	for _, v := range collection {
		if predicate(v) {
			output = append(output, v)
		}
	}
	return output
}

func Discard[K any](collection []K, predicate func(K) bool) []K {
	var output []K
	for _, v := range collection {
		if !predicate(v) {
			output = append(output, v)
		}
	}
	return output
}

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
