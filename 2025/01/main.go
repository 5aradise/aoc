package main

import (
	"bytes"
	"errors"
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

	rots, err := getInput(data)
	if err != nil {
		panic(err)
	}

	zeros := zeroPointings(rots)
	allZeros := allZeroPointings(rots)

	fmt.Println(aoc.FormatAnswers(zeros, allZeros))
}

func getInput(data []byte) (rotations []int, err error) {
	for row := range bytes.SplitSeq(data, []byte("\n")) {
		rotation, err := parseRotation(string(row))
		if err != nil {
			return nil, err
		}
		rotations = append(rotations, rotation)
	}
	return rotations, nil
}

// R +
// L -
func parseRotation(s string) (int, error) {
	dir, val := s[0], s[1:]
	v, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	switch dir {
	case 'R':
	case 'L':
		v = -v
	default:
		return 0, errors.New("unknown direction")
	}
	return v, nil
}

func zeroPointings(rotations []int) int {
	var (
		pos   = 50
		zeros int
	)
	for _, rot := range rotations {
		pos = (pos + rot) % 100
		if pos == 0 {
			zeros++
		}
	}
	return zeros
}

func allZeroPointings(rotations []int) int {
	var (
		pos     = 50
		zeros   int
		circles int
	)
	for _, rot := range rotations {
		pos += rot
		if pos == 0 {
			zeros++
		} else if pos < 0 {
			if pos != rot {
				zeros++
			}
		}
		circles, pos = pos/100, pos%100
		if pos < 0 {
			pos += 100
		}
		if circles < 0 {
			circles = -circles
		}
		zeros += circles
	}
	return zeros
}
