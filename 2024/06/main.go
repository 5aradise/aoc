package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/5aradise/aoc"
)

func main() {
	m1 := NewMap(getInput(aoc.InputFile))
	m1.Show()
	m2 := m1.Copy()
	fmt.Println("Count visited positions:")
	drawRoute(m1, getGuardPos(m1))
	fmt.Println(m1.CountVisitedPos())
	fmt.Println("Count new obstructions:")
	fmt.Println(countNewObs(m2, getGuardPos(m2)))
}

type mark byte

const (
	empty      mark = '.'
	obs        mark = '#'
	visitedObs mark = 'O'
	visited    mark = 'X'
	visited2   mark = '2'
	visited3   mark = '3'
	guard      mark = '^'
)

type Map [][]mark

func NewMap(bs [][]byte) Map {
	m := make(Map, len(bs))
	for i, row := range bs {
		mr := make([]mark, len(row))
		for i, v := range row {
			mr[i] = mark(v)
		}
		m[i] = mr
	}
	return m
}

func (m Map) Copy() Map {
	cp := make(Map, len(m))
	for i, row := range m {
		cp[i] = make([]mark, len(row))
		copy(cp[i], row)
	}
	return cp
}

func (m Map) IsIn(pos aoc.Vec2) bool {
	return (0 <= pos.X && pos.X <= len(m[0])-1) && (0 <= pos.Y && pos.Y <= len(m)-1)
}

func (m Map) Get(pos aoc.Vec2) mark {
	return m[pos.Y][pos.X]
}

func (m Map) Set(pos aoc.Vec2, value mark) {
	m[pos.Y][pos.X] = value
}

func (m Map) CountVisitedPos() int {
	count := 0
	for _, row := range m {
		for _, pos := range row {
			if pos == visited {
				count++
			}
		}
	}
	return count
}

func (m Map) IsLoop(g, d aoc.Vec2) bool {
	isRotated := false
	for {
		if !isRotated {
			switch m.Get(g) {
			case visited:
				m.Set(g, visited2)
			case visited2:
				m.Set(g, visited3)
			case visited3:
				return true
			default:
				m.Set(g, visited)
			}
		}
		g = g.Add(d)
		isRotated = false
		if !m.IsIn(g) {
			return false
		}
		if m.Get(g) == obs {
			g = g.Sub(d)
			d = d.RotateRight()
			isRotated = true
		}
	}
}

func (m Map) Show() {
	for _, row := range m {
		fmt.Println(string(row))
	}
}

func getInput(inputFile string) [][]byte {
	data, _ := os.ReadFile(inputFile)
	return bytes.Split(data, []byte{'\n'})
}

func getGuardPos(mp Map) aoc.Vec2 {
	for i, row := range mp {
		for j, r := range row {
			if r == guard {
				return aoc.Vec2{X: j, Y: i}
			}
		}
	}
	return aoc.Vec2{}
}

func drawRoute(m Map, g aoc.Vec2) {
	d := aoc.Vec2{X: 0, Y: -1}
	for {
		m.Set(g, visited)
		g = g.Add(d)
		if !m.IsIn(g) {
			return
		}
		if m.Get(g) == obs {
			g = g.Sub(d)
			d = d.RotateRight()
		}
	}
}

func countNewObs(m Map, g aoc.Vec2) int {
	count := 0
	newObs := make(aoc.Set[aoc.Vec2])
	d := aoc.Vec2{X: 0, Y: -1}
	for {
		g = g.Add(d)
		if !m.IsIn(g) {
			return count
		}
		if m.Get(g) == obs {
			g = g.Sub(d)
			d = d.RotateRight()
		} else if !newObs.Has(g) {
			mc := m.Copy()
			mc.Set(g, obs)
			newObs.Add(g)
			if mc.IsLoop(g.Sub(d), d) {
				count++
			}
		}
	}
}
