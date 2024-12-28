package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/5aradise/aoc"
)

func main() {
	antennas, height, length := getInput(aoc.InputFile)
	fmt.Println("Impact of the signal:")
	fmt.Println(len(getAntinodes(antennas, height, length)))
	fmt.Println("Impact of the signal using updated model:")
	fmt.Println(len(getAllAntinodes(antennas, height, length)))
}

func getInput(inputFile string) (antennas map[byte][]aoc.Vec2, height, length int) {
	data, _ := os.ReadFile(inputFile)
	mp := bytes.Split(data, []byte{'\n'})
	antennas = make(map[byte][]aoc.Vec2)
	for i, row := range mp {
		for j, v := range row {
			if v != '.' {
				antennas[v] = append(antennas[v], aoc.Vec2{X: j, Y: i})
			}
		}
	}
	return antennas, len(mp), len(mp[0])
}

func getAntinodes(ass map[byte][]aoc.Vec2, h, l int) aoc.Set[aoc.Vec2] {
	antinodes := make(aoc.Set[aoc.Vec2])
	for _, as := range ass {
		for i, a1 := range as {
			for _, a2 := range as[i+1:] {
				an1, an2 := calcAntinodesPos(a1, a2)
				if isPosIn(h, l, an1) {
					antinodes.Add(an1)
				}
				if isPosIn(h, l, an2) {
					antinodes.Add(an2)
				}
			}
		}
	}
	return antinodes
}

func calcAntinodesPos(a1, a2 aoc.Vec2) (aoc.Vec2, aoc.Vec2) {
	dv := a1.Sub(a2)
	return a1.Add(dv), a2.Sub(dv)
}

func getAllAntinodes(ass map[byte][]aoc.Vec2, h, l int) aoc.Set[aoc.Vec2] {
	antinodes := make(aoc.Set[aoc.Vec2])
	for _, as := range ass {
		for i, a1 := range as {
			antinodes.Add(a1)
			for _, a2 := range as[i+1:] {
				poss := calcAllAntinodesPos(h, l, a1, a2)
				for _, pos := range poss {
					antinodes.Add(pos)
				}
			}
		}
	}
	return antinodes
}

func calcAllAntinodesPos(h, l int, a1, a2 aoc.Vec2) []aoc.Vec2 {
	var pos []aoc.Vec2
	dv := a1.Sub(a2)
	curr := a1.Add(dv)
	for isPosIn(h, l, curr) {
		pos = append(pos, curr)
		curr = curr.Add(dv)
	}
	curr = a2.Sub(dv)
	for isPosIn(h, l, curr) {
		pos = append(pos, curr)
		curr = curr.Sub(dv)
	}
	return pos
}

func isPosIn(height, length int, pos aoc.Vec2) bool {
	return 0 <= pos.X && pos.X < length && 0 <= pos.Y && pos.Y < height
}
