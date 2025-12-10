package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"

	"github.com/5aradise/aoc"
)

type box struct {
	x, y, z int
}

func (b1 box) dist(b2 box) float64 {
	xDiff := b1.x - b2.x
	yDiff := b1.y - b2.y
	zDiff := b1.z - b2.z
	return math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff + zDiff*zDiff))
}

func main() {
	data, err := os.ReadFile(aoc.InputFile)
	if err != nil {
		panic(err)
	}

	boxes, err := getInput(data)
	if err != nil {
		panic(err)
	}

	cs := makeCircuits(boxes, 1000)
	lc := findLastConn(boxes)

	fmt.Println(aoc.FormatAnswers(sizeMulOf3Largest(cs), lc[0].x*lc[1].x))
}

func getInput(data []byte) ([]box, error) {
	rows := bytes.Split(data, []byte("\n"))
	boxes := make([]box, 0, len(rows))
	for _, row := range rows {
		coords := bytes.SplitN(row, []byte(","), 3)

		var (
			b   box
			err error
		)
		b.x, err = strconv.Atoi(string(coords[0]))
		if err != nil {
			return nil, err
		}
		b.y, err = strconv.Atoi(string(coords[1]))
		if err != nil {
			return nil, err
		}
		b.z, err = strconv.Atoi(string(coords[2]))
		if err != nil {
			return nil, err
		}
		boxes = append(boxes, b)
	}
	return boxes, nil
}

type conn struct {
	points [2]box
	dist   float64
}

func makeCircuits(boxes []box, connsCount int) [][]box {
	conns := make([]conn, 0, len(boxes)*(len(boxes)-1)/2)
	for i, box1 := range boxes {
		for _, box2 := range boxes[i+1:] {
			conns = append(conns, conn{
				points: [2]box{box1, box2},
				dist:   box1.dist(box2),
			})
		}
	}

	shortConns := make(map[box][]box)
	for range connsCount {
		var (
			shortI = -1
			short  = conn{dist: math.MaxFloat64}
		)
		for i, conn := range conns {
			if conn.dist < short.dist {
				shortI = i
				short = conn
			}
		}
		conns[shortI] = conn{dist: math.MaxFloat64}
		shortConns[short.points[0]] = append(shortConns[short.points[0]], short.points[1])
		shortConns[short.points[1]] = append(shortConns[short.points[1]], short.points[0])
	}

	var circuits [][]box
	for b := range shortConns {
		circuits = append(circuits, makeCircuit(nil, shortConns, b))
	}

	return circuits
}

func makeCircuit(circuit []box, conns map[box][]box, b box) []box {
	rbs, ok := conns[b]
	if !ok {
		return circuit
	}
	delete(conns, b)

	for _, rb := range rbs {
		circuit = makeCircuit(circuit, conns, rb)
	}
	return append(circuit, b)
}

func sizeMulOf3Largest(circuits [][]box) int {
	slices.SortFunc(circuits, func(a, b []box) int { return -(len(a) - len(b)) })
	res := 1
	for _, c := range circuits[:3] {
		res *= len(c)
	}
	return res
}

func findLastConn(boxes []box) [2]box {
	conns := make([]conn, 0, len(boxes)*(len(boxes)-1)/2)
	for i, box1 := range boxes {
		for _, box2 := range boxes[i+1:] {
			conns = append(conns, conn{
				points: [2]box{box1, box2},
				dist:   box1.dist(box2),
			})
		}
	}
	slices.SortFunc(conns, func(a, b conn) int { return int(a.dist - b.dist) })

	var (
		nextMembID int
		circuits   = make(map[int][]box)
		membership = make(map[box]int)
		lastConn   [2]box
	)
	for _, c := range conns {
		memb0, ok0 := membership[c.points[0]]
		memb1, ok1 := membership[c.points[1]]

		if ok0 {
			if ok1 {
				if memb0 == memb1 {
					continue
				}

				var (
					membBig   = memb0
					membSmall = memb1
					cirBig    = circuits[memb0]
					cirSmall  = circuits[memb1]
				)
				if len(cirSmall) > len(cirBig) {
					membSmall, membBig = membBig, membSmall
					cirSmall, cirBig = cirBig, cirSmall
				}
				circuits[membBig] = append(cirBig, cirSmall...)
				delete(circuits, membSmall)
				for _, b := range cirSmall {
					membership[b] = membBig
				}

			} else {
				circuits[memb0] = append(circuits[memb0], c.points[1])
				membership[c.points[1]] = memb0
			}
		} else if ok1 {
			circuits[memb1] = append(circuits[memb1], c.points[0])
			membership[c.points[0]] = memb1
		} else {

			memb := nextMembID
			circuits[memb] = c.points[:]
			membership[c.points[0]] = memb
			membership[c.points[1]] = memb
			nextMembID++
		}

		if len(membership) == len(boxes) && len(circuits) == 1 {
			lastConn = c.points
			break
		}
	}
	return lastConn
}
