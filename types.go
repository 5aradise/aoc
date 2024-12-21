package aoc

import (
	"iter"
)

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

type Map[T any] struct {
	data   []T
	length int
}

func NewMap[T any](length, height int) Map[T] {
	return Map[T]{
		data:   make([]T, length*height),
		length: length,
	}
}

func NewMapFromSlice[T any](sl []T, length int) Map[T] {
	return Map[T]{
		data:   sl,
		length: length,
	}
}

func (m Map[T]) Get(pos Vec2) T {
	return m.data[pos.Y*m.length+pos.X]
}

func (m Map[T]) Set(pos Vec2, v T) {
	m.data[pos.Y*m.length+pos.X] = v
}

func (m Map[T]) IsIn(pos Vec2) bool {
	return m.IsInX(pos.X) && m.IsInY(pos.Y)
}

func (m Map[T]) IsInX(x int) bool {
	return 0 <= x && x < m.length
}

func (m Map[T]) IsInY(y int) bool {
	return 0 <= y && y < len(m.data)/m.length
}

func (m Map[T]) Positions() iter.Seq2[Vec2, T] {
	return func(yield func(Vec2, T) bool) {
		for i, v := range m.data {
			if !yield(Vec2{i % m.length, i / m.length}, v) {
				return
			}
		}
	}
}

func (m Map[T]) Copy() Map[T] {
	cp := make([]T, len(m.data))
	copy(cp, m.data)
	return NewMapFromSlice(cp, m.length)
}
