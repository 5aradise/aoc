package aoc

func Sum[N Number](sl []N) N {
	var sum N
	for _, v := range sl {
		sum += v
	}
	return sum
}
