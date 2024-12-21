package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/5aradise/aoc"
)

func main() {
	m := getInput(aoc.InputFile)
	m2 := m.Copy()
	fmt.Println("Total price:")
	fmt.Println(calcTotalPrice(m, getP))
	fmt.Println("Total price with bulk discount:")
	fmt.Println(calcTotalPrice(m2, getN))
}

func getInput(inputFile string) aoc.Map[byte] {
	data, _ := os.ReadFile(inputFile)
	sep := []byte{'\n'}
	return aoc.NewMapFromSlice(bytes.ReplaceAll(data, sep, nil), bytes.Index(data, sep))
}

func calcTotalPrice(m aoc.Map[byte], getV func(mark byte, start aoc.Vec2, m aoc.Map[byte], visited aoc.Set[aoc.Vec2]) int) int {
	total := 0
	for pos, v := range m.Positions() {
		if v != 0 {
			visited := make(aoc.Set[aoc.Vec2])
			v := getV(m.Get(pos), pos, m, visited)
			total += len(visited) * v
		}
	}
	return total
}

func getP(mark byte, curr aoc.Vec2, m aoc.Map[byte], visited aoc.Set[aoc.Vec2]) int {
	if visited.Has(curr) {
		return 0
	}
	visited.Add(curr)
	m.Set(curr, 0)
	p := 0
	sides := []aoc.Vec2{
		curr.Add(aoc.Vec2{X: -1, Y: 0}), curr.Add(aoc.Vec2{X: 1, Y: 0}),
		curr.Add(aoc.Vec2{X: 0, Y: -1}), curr.Add(aoc.Vec2{X: 0, Y: 1}),
	}
	for _, side := range sides {
		if m.IsIn(side) {
			if m.Get(side) == mark {
				p += getP(mark, side, m, visited)
			} else {
				if !visited.Has(side) {
					p++
				}
			}
		} else {
			p++
		}
	}
	return p
}

func getN(mark byte, curr aoc.Vec2, m aoc.Map[byte], visited aoc.Set[aoc.Vec2]) int {
	if visited.Has(curr) {
		return 0
	}
	visited.Add(curr)
	m.Set(curr, 0)
	n := 0
	sides := []aoc.Vec2{
		curr.Add(aoc.Vec2{X: -1, Y: 0}), curr.Add(aoc.Vec2{X: -1, Y: -1}), curr.Add(aoc.Vec2{X: 0, Y: -1}),
		curr.Add(aoc.Vec2{X: 1, Y: -1}), curr.Add(aoc.Vec2{X: 1, Y: 0}), curr.Add(aoc.Vec2{X: 1, Y: 1}),
		curr.Add(aoc.Vec2{X: 0, Y: 1}), curr.Add(aoc.Vec2{X: -1, Y: 1}),
	}
	isNeigh := make([]bool, 8)
	for i, side := range sides {
		isNeigh[i] = m.IsIn(side) && (m.Get(side) == mark || visited.Has(side))
	}
	for i, side := range sides {
		if i%2 == 0 {
			if isNeigh[i] {
				n += getN(mark, side, m, visited)
			}
		} else {
			if !isNeigh[i-1] && !isNeigh[(i+1)%8] || isNeigh[i-1] && !isNeigh[i] && isNeigh[(i+1)%8] {
				n += 1
			}
		}
	}
	return n
}
