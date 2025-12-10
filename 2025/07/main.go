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

	diagram := getInput(data)

	cs := countSplits(diagram)
	ct := countTimelines(diagram)

	fmt.Println(aoc.FormatAnswers(cs, ct))
}

func getInput(data []byte) [][]byte {
	return bytes.Split(data, []byte("\n"))
}

func countSplits(diagram [][]byte) int {
	var splits int

	curr := make([]bool, len(diagram[0]))
	curr[len(curr)/2] = true

	for _, row := range diagram[1:] {
		for i := range row {
			if row[i] == '^' && curr[i] {
				curr[i-1] = true
				curr[i] = false
				curr[i+1] = true
				splits++
			}
		}
	}
	return splits
}

func countTimelines(diagram [][]byte) int {
	curr := make([]int, len(diagram[0]))
	curr[len(curr)/2] = 1

	for _, row := range diagram[1:] {
		for i := range row {
			if cell := curr[i]; row[i] == '^' && cell != 0 {
				curr[i-1] += cell
				curr[i] = 0
				curr[i+1] += cell
			}
		}
	}

	var res int
	for _, v := range curr {
		res += v
	}
	return res
}
