package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/5aradise/aoc"
)

func main() {
	r, m, r2, m2, moves := getInput(aoc.InputFile)
	fmt.Println("Sum of all boxes' GPS coordinates:")
	moveRobot(r, m, moves)
	fmt.Println(calcBoxesGPSSum(m, box))
	fmt.Println("Sum of all boxes' GPS coordinates (twice as wide):")
	moveRobot2(r2, m2, moves)
	fmt.Println(calcBoxesGPSSum(m2, boxL))
}

type robot struct {
	pos aoc.Vec2
}

const (
	up    byte = '^'
	down  byte = 'v'
	left  byte = '<'
	right byte = '>'

	robotIcon  byte = '@'
	empty      byte = '.'
	box        byte = 'O'
	boxL, boxR byte = '[', ']'
	wall       byte = '#'
)

func getInput(inputFile string) (r *robot, m aoc.Map[byte], r2 *robot, m2 aoc.Map[byte], moves []byte) {
	data, _ := os.ReadFile(inputFile)
	parts := bytes.Split(data, []byte("\n\n"))
	sep := []byte("\n")
	linedMap := bytes.ReplaceAll(parts[0], sep, nil)
	length := bytes.Index(parts[0], sep)
	doubled := doubleLine(linedMap)
	m = aoc.NewMapFromSlice(linedMap, length)
	for pos, v := range m.Positions() {
		if v == robotIcon {
			r = &robot{pos}
			m.Set(pos, empty)
			break
		}
	}
	m2 = aoc.NewMapFromSlice(doubled, length*2)
	r2 = &robot{r.pos}
	r2.pos.X *= 2
	m2.Set(r2.pos, empty)
	m2.Set(r2.pos.Add(aoc.Vec2{X: 1, Y: 0}), empty)
	rows := bytes.Split(parts[1], []byte("\n"))
	for _, row := range rows {
		moves = append(moves, row...)
	}
	return
}

func doubleLine(line []byte) []byte {
	double := make([]byte, len(line)*2)
	for i, v := range line {
		if v == box {
			double[2*i] = boxL
			double[2*i+1] = boxR
		} else {
			double[2*i] = v
			double[2*i+1] = v
		}
	}
	return double
}

func (r *robot) Move(m aoc.Map[byte], dir aoc.Vec2) {
	next := r.pos.Add(dir)
	curr := next
	for m.IsIn(curr) && m.Get(curr) == box {
		curr = curr.Add(dir)
	}
	if !m.IsIn(curr) || m.Get(curr) == wall {
		return
	}
	if next != curr {
		m.Set(curr, box)
		m.Set(next, empty)
	}
	r.pos = next
}

func moveRobot(r *robot, m aoc.Map[byte], moves []byte) {
	for _, move := range moves {
		switch move {
		case up:
			r.Move(m, aoc.Vec2{X: 0, Y: -1})
		case down:
			r.Move(m, aoc.Vec2{X: 0, Y: 1})
		case left:
			r.Move(m, aoc.Vec2{X: -1, Y: 0})
		case right:
			r.Move(m, aoc.Vec2{X: 1, Y: 0})
		}
	}
}

func moveVert(start aoc.Vec2, m aoc.Map[byte], dir aoc.Vec2, toMove map[aoc.Vec2]bool) bool {
	canMove, visited := toMove[start]
	if visited {
		return canMove
	}

	next := start.Add(dir)
	if !m.IsIn(next) {
		toMove[start] = false
		return false
	}
	nextV := m.Get(next)
	if nextV == wall {
		toMove[start] = false
		return false
	}
	if nextV == empty {
		toMove[start] = true
		return true
	}

	var nextL, nextR aoc.Vec2
	if nextV == boxL {
		nextL = next
		nextR = next.Add(aoc.Vec2{X: 1, Y: 0})
	} else {
		nextL = next.Add(aoc.Vec2{X: -1, Y: 0})
		nextR = next
	}
	canMove = moveVert(nextL, m, dir, toMove) && moveVert(nextR, m, dir, toMove)
	toMove[start] = canMove
	return canMove
}

func moveAllBlocks(m aoc.Map[byte], canMoves map[aoc.Vec2]bool, dir aoc.Vec2) {
	toMoves := make(map[aoc.Vec2]byte)
	for toMove := range canMoves {
		toMoves[toMove] = m.Get(toMove)
		m.Set(toMove, empty)
	}
	for toMove := range canMoves {
		m.Set(toMove.Add(dir), toMoves[toMove])
	}
}

func moveL(start aoc.Vec2, m aoc.Map[byte]) bool {
	dir := aoc.Vec2{X: -1, Y: 0}
	dir2 := aoc.Vec2{X: -2, Y: 0}
	next := start.Add(dir)
	curr := next
	for m.IsIn(curr) && m.Get(curr) == boxR {
		curr = curr.Add(dir2)
	}
	if !m.IsIn(curr) || m.Get(curr) == wall {
		return false
	}
	if next != curr {
		curr = next
		m.Set(curr, empty)
		curr = curr.Add(dir)
		for m.Get(curr) != empty {
			if m.Get(curr) == boxL {
				m.Set(curr, boxR)
			} else {
				m.Set(curr, boxL)
			}
			curr = curr.Add(dir)
		}
		m.Set(curr, boxL)
	}
	return true
}

func moveR(start aoc.Vec2, m aoc.Map[byte]) bool {
	dir := aoc.Vec2{X: 1, Y: 0}
	dir2 := aoc.Vec2{X: 2, Y: 0}
	next := start.Add(dir)
	curr := next
	for m.IsIn(curr) && m.Get(curr) == boxL {
		curr = curr.Add(dir2)
	}
	if !m.IsIn(curr) || m.Get(curr) == wall {
		return false
	}
	if next != curr {
		curr = next
		m.Set(curr, empty)
		curr = curr.Add(dir)
		for m.Get(curr) != empty {
			if m.Get(curr) == boxL {
				m.Set(curr, boxR)
			} else {
				m.Set(curr, boxL)
			}
			curr = curr.Add(dir)
		}
		m.Set(curr, boxR)
	}
	return true
}

func moveRobot2(r *robot, m aoc.Map[byte], moves []byte) {
	for _, move := range moves {
		switch move {
		case up:
			dir := aoc.Vec2{X: 0, Y: -1}
			canMoves := make(map[aoc.Vec2]bool)
			if moveVert(r.pos, m, dir, canMoves) {
				moveAllBlocks(m, canMoves, dir)
				r.pos = r.pos.Add(dir)
			}
		case down:
			dir := aoc.Vec2{X: 0, Y: 1}
			canMoves := make(map[aoc.Vec2]bool)
			if moveVert(r.pos, m, dir, canMoves) {
				moveAllBlocks(m, canMoves, dir)
				r.pos = r.pos.Add(dir)
			}
		case left:
			if moveL(r.pos, m) {
				r.pos = r.pos.Add(aoc.Vec2{X: -1, Y: 0})
			}
		case right:
			if moveR(r.pos, m) {
				r.pos = r.pos.Add(aoc.Vec2{X: 1, Y: 0})
			}
		}
	}
}

func calcBoxesGPSSum(m aoc.Map[byte], box byte) int {
	sum := 0
	for pos, v := range m.Positions() {
		if v == box {
			sum += pos.X + pos.Y*100
		}
	}
	return sum
}
