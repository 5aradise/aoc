package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/5aradise/aoc"
)

func main() {
	data, err := os.ReadFile(aoc.InputFile)
	if err != nil {
		panic(err)
	}

	grid := getInput(data)

	fa := forkliftAccess(grid)
	ffa := fullForkliftAccess(grid)

	fmt.Println(aoc.FormatAnswers(fa, ffa))
}

func getInput(data []byte) aoc.Map[bool] {
	lines := bytes.Split(data, []byte{'\n'})
	grid := aoc.NewMap[bool](len(lines[0]), len(lines))
	for i, line := range lines {
		for j, c := range line {
			if c == '@' {
				grid.Set(aoc.Vec2{X: j, Y: i}, true)
			}
		}
	}
	return grid
}

func forkliftAccess(grid aoc.Map[bool]) int {
	var n int
	for pos, v := range grid.Positions() {
		if v && neighbors(pos, grid) < 4 {
			n++
		}
	}
	return n
}

func fullForkliftAccess(grid aoc.Map[bool]) int {
	var (
		n    = 0
		curr = -1
	)
	for curr != 0 {
		curr = 0
		for pos, v := range grid.Positions() {
			if v && neighbors(pos, grid) < 4 {
				grid.Set(pos, false)
				curr++
			}
		}
		n += curr
	}
	return n
}

func neighbors(pos aoc.Vec2, grid aoc.Map[bool]) int {
	var n int
	t, b := grid.IsInY(pos.Y-1), grid.IsInY(pos.Y+1)
	l, r := grid.IsInX(pos.X-1), grid.IsInX(pos.X+1)
	if t {
		if l {
			if grid.Get(pos.Add(aoc.Vec2{X: -1, Y: -1})) {
				n++
			}
		}
		if grid.Get(pos.Add(aoc.Vec2{X: 0, Y: -1})) {
			n++
		}
		if r {
			if grid.Get(pos.Add(aoc.Vec2{X: +1, Y: -1})) {
				n++
			}
		}
	}
	{
		if l {
			if grid.Get(pos.Add(aoc.Vec2{X: -1, Y: 0})) {
				n++
			}
		}
		if r {
			if grid.Get(pos.Add(aoc.Vec2{X: +1, Y: 0})) {
				n++
			}
		}
	}
	if b {
		if l {
			if grid.Get(pos.Add(aoc.Vec2{X: -1, Y: +1})) {
				n++
			}
		}
		if grid.Get(pos.Add(aoc.Vec2{X: 0, Y: +1})) {
			n++
		}
		if r {
			if grid.Get(pos.Add(aoc.Vec2{X: +1, Y: +1})) {
				n++
			}
		}
	}
	return n
}
