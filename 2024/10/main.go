package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/5aradise/aoc"
)

func main() {
	mp := getInput(aoc.InputFile)
	fmt.Println("Sum of the scores of trailheads:")
	fmt.Println(aoc.Sum(findPathsScores(mp)))
	fmt.Println("Sum of the ratings of trailheads:")
	fmt.Println(aoc.Sum(findPathsRating(mp)))
}

var figures = map[byte]uint8{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func getInput(inputFile string) [][]uint8 {
	data, _ := os.ReadFile(inputFile)
	bmp := bytes.Split(data, []byte{'\n'})
	mp := make([][]uint8, len(bmp))
	for i, brow := range bmp {
		row := make([]uint8, len(brow))
		for j, bv := range brow {
			row[j] = figures[bv]
		}
		mp[i] = row
	}
	return mp
}

func findPathsScores(mp [][]uint8) []int {
	var paths []int
	for i, row := range mp {
		for j, v := range row {
			if v == 0 {
				visited := aoc.Set[aoc.Vec2]{}
				paths = append(paths, findPathScoreFrom(aoc.Vec2{X: j, Y: i}, 1, visited, mp))
			}
		}
	}
	return paths
}

func findPathScoreFrom(curr aoc.Vec2, toFind uint8, visited aoc.Set[aoc.Vec2], mp [][]uint8) int {
	if visited.Has(curr) {
		return 0
	}
	visited.Add(curr)
	count := 0
	if toFind == 9 {
		curr1 := curr.Add(aoc.Vec2{X: 1, Y: 0})
		if curr1.X < len(mp[0]) && getByVec(mp, curr1) == toFind && !visited.Has(curr1) {
			visited.Add(curr1)
			count += 1
		}
		curr2 := curr.Add(aoc.Vec2{X: -1, Y: 0})
		if curr2.X >= 0 && getByVec(mp, curr2) == toFind && !visited.Has(curr2) {
			visited.Add(curr2)
			count += 1
		}
		curr3 := curr.Add(aoc.Vec2{X: 0, Y: 1})
		if curr3.Y < len(mp) && getByVec(mp, curr3) == toFind && !visited.Has(curr3) {
			visited.Add(curr3)
			count += 1
		}
		curr4 := curr.Add(aoc.Vec2{X: 0, Y: -1})
		if curr4.Y >= 0 && getByVec(mp, curr4) == toFind && !visited.Has(curr4) {
			visited.Add(curr4)
			count += 1
		}
	} else {
		curr1 := curr.Add(aoc.Vec2{X: 1, Y: 0})
		if curr1.X < len(mp[0]) && getByVec(mp, curr1) == toFind {
			count += findPathScoreFrom(curr1, toFind+1, visited, mp)
		}
		curr2 := curr.Add(aoc.Vec2{X: -1, Y: 0})
		if curr2.X >= 0 && getByVec(mp, curr2) == toFind {
			count += findPathScoreFrom(curr2, toFind+1, visited, mp)
		}
		curr3 := curr.Add(aoc.Vec2{X: 0, Y: 1})
		if curr3.Y < len(mp) && getByVec(mp, curr3) == toFind {
			count += findPathScoreFrom(curr3, toFind+1, visited, mp)
		}
		curr4 := curr.Add(aoc.Vec2{X: 0, Y: -1})
		if curr4.Y >= 0 && getByVec(mp, curr4) == toFind {
			count += findPathScoreFrom(curr4, toFind+1, visited, mp)
		}
	}
	return count
}

func findPathsRating(mp [][]uint8) []int {
	var paths []int
	for i, row := range mp {
		for j, v := range row {
			if v == 0 {
				paths = append(paths, findPathRatingFrom(aoc.Vec2{X: j, Y: i}, 1, mp))
			}
		}
	}
	return paths
}

func findPathRatingFrom(curr aoc.Vec2, toFind uint8, mp [][]uint8) int {
	count := 0
	if toFind == 9 {
		curr1 := curr.Add(aoc.Vec2{X: 1, Y: 0})
		if curr1.X < len(mp[0]) && getByVec(mp, curr1) == toFind {
			count += 1
		}
		curr2 := curr.Add(aoc.Vec2{X: -1, Y: 0})
		if curr2.X >= 0 && getByVec(mp, curr2) == toFind {
			count += 1
		}
		curr3 := curr.Add(aoc.Vec2{X: 0, Y: 1})
		if curr3.Y < len(mp) && getByVec(mp, curr3) == toFind {
			count += 1
		}
		curr4 := curr.Add(aoc.Vec2{X: 0, Y: -1})
		if curr4.Y >= 0 && getByVec(mp, curr4) == toFind {
			count += 1
		}
	} else {
		curr1 := curr.Add(aoc.Vec2{X: 1, Y: 0})
		if curr1.X < len(mp[0]) && getByVec(mp, curr1) == toFind {
			count += findPathRatingFrom(curr1, toFind+1, mp)
		}
		curr2 := curr.Add(aoc.Vec2{X: -1, Y: 0})
		if curr2.X >= 0 && getByVec(mp, curr2) == toFind {
			count += findPathRatingFrom(curr2, toFind+1, mp)
		}
		curr3 := curr.Add(aoc.Vec2{X: 0, Y: 1})
		if curr3.Y < len(mp) && getByVec(mp, curr3) == toFind {
			count += findPathRatingFrom(curr3, toFind+1, mp)
		}
		curr4 := curr.Add(aoc.Vec2{X: 0, Y: -1})
		if curr4.Y >= 0 && getByVec(mp, curr4) == toFind {
			count += findPathRatingFrom(curr4, toFind+1, mp)
		}
	}
	return count
}

func getByVec(mp [][]uint8, vec aoc.Vec2) uint8 {
	return mp[vec.Y][vec.X]
}
