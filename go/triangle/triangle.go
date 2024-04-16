package triangle

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

func KindFromSides(a, b, c float64) Kind {
	if !isValid(a, b, c) {
		return NaT
	}

	if isEquilateral(a, b, c) {
		return Equ
	}

	if isIsosceles(a, b, c) {
		return Iso
	}

	if isScalene(a, b, c) {
		return Sca
	}

	return NaT
}

func isValid(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if (a+b) >= c && (a+c) >= b && (c+b) >= a {
		return true
	}
	return false
}

func isEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

func isIsosceles(a, b, c float64) bool {
	return a == b || b == c || a == c
}

func isScalene(a, b, c float64) bool {
	return a != b && b != c
}
