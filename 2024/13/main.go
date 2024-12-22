package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/5aradise/aoc"
)

func main() {
	games := getInput(aoc.InputFile)
	fmt.Println("Tokens to be spent to win all possible prizes:")
	fmt.Println(calcTotal(games, 0))
	fmt.Println("Tokens to be spent to win all possible prizes(prize position increase = 10000000000000):")
	fmt.Println(calcTotal(games, 10000000000000))
}

type game struct {
	a, b, p aoc.Vec2
}

func getInput(inputFile string) []game {
	data, _ := os.ReadFile(inputFile)
	bGames := bytes.Split(data, []byte("\n\nButton A: X+"))
	bGames[0] = bGames[0][12:]
	games := make([]game, len(bGames))
	for i, bGame := range bGames {
		rows := bytes.Split(bGame, []byte("\n"))
		aXY := bytes.Split(rows[0], []byte(", Y+"))
		x, _ := strconv.Atoi(string(aXY[0]))
		y, _ := strconv.Atoi(string(aXY[1]))
		a := aoc.Vec2{X: x, Y: y}
		bXY := bytes.Split(rows[1][12:], []byte(", Y+"))
		x, _ = strconv.Atoi(string(bXY[0]))
		y, _ = strconv.Atoi(string(bXY[1]))
		b := aoc.Vec2{X: x, Y: y}
		pXY := bytes.Split(rows[2][9:], []byte(", Y="))
		x, _ = strconv.Atoi(string(pXY[0]))
		y, _ = strconv.Atoi(string(pXY[1]))
		prize := aoc.Vec2{X: x, Y: y}
		games[i] = game{a, b, prize}
	}
	return games
}

func calcTotal(gs []game, prizeInc int) int {
	total := 0
	for _, g := range gs {
		g.p = g.p.Add(aoc.Vec2{X: prizeInc, Y: prizeInc})
		total += calcTokens(g)
	}
	return total
}

func calcTokens(g game) int {
	axy := float64(g.a.X) / float64(g.a.Y)
	b := (float64(g.p.X) - axy*float64(g.p.Y)) /
		(float64(g.b.X) - axy*float64(g.b.Y))
	if b < 0 || !aoc.IsInteger(b) {
		return 0
	}
	a := (float64(g.p.X) - b*float64(g.b.X)) / float64(g.a.X)
	if a < 0 || !aoc.IsInteger(a) {
		return 0
	}
	return 3*int(math.Round(a)) + int(math.Round(b))
}
