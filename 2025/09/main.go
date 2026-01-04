package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/5aradise/aoc"
)

func main() {
	data, err := os.ReadFile(aoc.InputFile)
	if err != nil {
		panic(err)
	}

	redTiles, err := getInput(data)
	if err != nil {
		panic(err)
	}

	la := largestRectangleArea(redTiles)

	fmt.Println(aoc.FormatAnswers(la, nil))
}

func getInput(data []byte) ([]aoc.Vec2, error) {
	lines := bytes.Split(data, []byte("\n"))
	redTiles := make([]aoc.Vec2, 0, len(lines))
	for _, line := range lines {
		parts := bytes.SplitN(line, []byte(","), 2)
		x, err := strconv.Atoi(string(parts[0]))
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(string(parts[1]))
		if err != nil {
			return nil, err
		}
		redTiles = append(redTiles, aoc.Vec2{X: x, Y: y})
	}
	return redTiles, nil
}

func largestRectangleArea(tiles []aoc.Vec2) int {
	var la int
	for i, t1 := range tiles {
		for _, t2 := range tiles[i+1:] {
			a := rectangleArea(t1, t2)
			if a > la {
				la = a
			}
		}
	}
	return la
}

func rectangleArea(t1, t2 aoc.Vec2) int {
	dx := t1.X - t2.X
	if dx < 0 {
		dx = -dx
	}
	dx++
	dy := t1.Y - t2.Y
	if dy < 0 {
		dy = -dy
	}
	dy++
	return dx * dy
}
