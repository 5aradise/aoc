package aoc

import "math"

func Sum[N Number](sl []N) N {
	var sum N
	for _, v := range sl {
		sum += v
	}
	return sum
}

var Epsilon = 0.0001

func IsInteger(f float64) bool {
	return math.Abs(f-math.Round(f)) < Epsilon
}
