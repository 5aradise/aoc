package aoc

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

type Vec2 struct {
	X, Y int
}

func (v Vec2) Add(dv Vec2) Vec2 {
	return Vec2{v.X + dv.X, v.Y + dv.Y}
}

func (v Vec2) Sub(dv Vec2) Vec2 {
	return Vec2{v.X - dv.X, v.Y - dv.Y}
}

func (v Vec2) RotateRight() Vec2 {
	return Vec2{-v.Y, v.X}
}

func (v Vec2) RotateLeft() Vec2 {
	return Vec2{v.Y, -v.X}
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Has(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) Delete(value T) {
	delete(s, value)
}
