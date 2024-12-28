package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/5aradise/aoc"
)

func main() {
	ls, ks := getInput(aoc.InputFile)
	fmt.Println("Unique lock/key pairs fit together without overlapping:")
	fmt.Println(len(findFitPairs(ls, ks)))
}

type schema [5]uint8

type pair struct {
	lock, key schema
}

func getInput(inputFile string) (locks, keys aoc.Set[schema]) {
	locks, keys = make(aoc.Set[schema]), make(aoc.Set[schema])
	data, _ := os.ReadFile(inputFile)
	schemas := bytes.Split(data, []byte("\n\n"))
	for _, rawSchema := range schemas {
		var s schema
		rows := bytes.Split(rawSchema, []byte("\n"))
		for _, row := range rows[1:6] {
			for i, v := range row {
				if v == '#' {
					s[i]++
				}
			}
		}
		if rows[0][0] == '#' {
			locks.Add(s)
		} else {
			keys.Add(s)
		}
	}
	return
}

func findFitPairs(locks, keys aoc.Set[schema]) []pair {
	var pairs []pair
	for lock := range locks {
	nextKey:
		for key := range keys {
			for i := range lock {
				if lock[i]+key[i] > 5 {
					continue nextKey
				}
			}
			pairs = append(pairs, pair{lock, key})
		}
	}
	return pairs
}
