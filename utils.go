package aoc

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

func Sum[N Number](sl []N) N {
	var sum N
	for _, v := range sl {
		sum += v
	}
	return sum
}
