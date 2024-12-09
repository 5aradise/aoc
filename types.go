package aoc

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

type Set[T comparable] map[T]bool

func (s Set[T]) Add(value T) {
	s[value] = true
}

func (s Set[T]) Has(value T) bool {
	return s[value]
}

func (s Set[T]) Delete(value T) {
	delete(s, value)
}
