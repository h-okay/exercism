package darts

import (
	"math"
)

const (
	outerRadius  = 10
	middleRadius = 5
	innerRadius  = 1
)

func Score(x, y float64) int {
	distanceToCenter := math.Sqrt(x*x + y*y)
	switch {
	case distanceToCenter > outerRadius:
		return 0
	case distanceToCenter > middleRadius:
		return 1
	case distanceToCenter > innerRadius:
		return 5
	default:
		return 10
	}
}
