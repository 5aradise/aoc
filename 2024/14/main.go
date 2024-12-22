package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/5aradise/aoc"
)

const MAP_LENGTH, MAP_HEIGHT = 101, 103

func main() {
	bots := getInput(aoc.InputFile)
	fmt.Println("Safety factor after 100 seconds:")
	fmt.Println(calcSafetyFactor(bots, 100))
	fmt.Println("Fewest number of seconds that must elapse for the robots to display the Easter egg:")
	secs, poss := secsToEE(bots)
	fmt.Println(secs)
	showBots(poss)
}

type bot struct {
	pos, v aoc.Vec2
}

func getInput(inputFile string) []bot {
	data, _ := os.ReadFile(inputFile)
	rows := bytes.Split(data, []byte{'\n'})
	bots := make([]bot, len(rows))
	for i, row := range rows {
		pv := bytes.Split(row[2:], []byte(" v="))
		p := bytes.Split(pv[0], []byte(","))
		px, _ := strconv.Atoi(string(p[0]))
		py, _ := strconv.Atoi(string(p[1]))
		v := bytes.Split(pv[1], []byte(","))
		vx, _ := strconv.Atoi(string(v[0]))
		vy, _ := strconv.Atoi(string(v[1]))
		bots[i] = bot{aoc.Vec2{X: px, Y: py}, aoc.Vec2{X: vx, Y: vy}}
	}
	return bots
}

func calcSafetyFactor(bs []bot, secs int) int {
	qs := quadrantsCount(bs, secs)
	return qs[0] * qs[1] * qs[2] * qs[3]
}

func quadrantsCount(bs []bot, secs int) [4]int {
	qs := [4]int{}
	for _, b := range bs {
		pos := posAfter(b, MAP_LENGTH, MAP_HEIGHT, secs)
		if pos.X < MAP_LENGTH/2 {
			if pos.Y < MAP_HEIGHT/2 {
				qs[0]++
			} else if pos.Y > MAP_HEIGHT/2 {
				qs[1]++
			}
		} else if pos.X > MAP_LENGTH/2 {
			if pos.Y < MAP_HEIGHT/2 {
				qs[2]++
			} else if pos.Y > MAP_HEIGHT/2 {
				qs[3]++
			}
		}
	}
	return qs
}

func secsToEE(bs []bot) (int, []aoc.Vec2) {
main:
	for k := range 10000 {
		uniq := make(aoc.Set[aoc.Vec2])
		for _, b := range bs {
			newPos := posAfter(b, MAP_LENGTH, MAP_HEIGHT, k)
			if uniq.Has(newPos) {
				continue main
			}
			uniq.Add(newPos)
		}
		return k, uniq.ToSlice()
	}
	return -1, nil
}

func posAfter(b bot, length, height, secs int) aoc.Vec2 {
	xAfter := (b.pos.X + secs*b.v.X) % length
	if xAfter < 0 {
		xAfter += length
	}
	yAfter := (b.pos.Y + secs*b.v.Y) % height
	if yAfter < 0 {
		yAfter += height
	}
	return aoc.Vec2{X: xAfter, Y: yAfter}
}

func showBots(bs []aoc.Vec2) {
	field := make([]string, MAP_LENGTH*MAP_HEIGHT)
	for i := range field {
		field[i] = "."
	}
	m := aoc.NewMapFromSlice(field, MAP_LENGTH)
	for _, b := range bs {
		m.Set(b, "@")
	}
	m.Show()
}
