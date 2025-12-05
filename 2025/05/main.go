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

	rs, ids, err := getInput(data)
	if err != nil {
		panic(err)
	}

	rs = compact(rs)

	fresh := aoc.Reduce(ids, 0, func(acc int, id int) int {
		if fresh(rs, id) {
			return acc + 1
		}
		return acc
	})
	allFresh := aoc.Reduce(rs, 0, func(acc int, r rng) int {
		return acc + (r[1] - r[0] + 1)
	})

	fmt.Println(aoc.FormatAnswers(fresh, allFresh))
}

type rng [2]int

var zeroRng rng

func (r rng) in(i int) bool {
	return r[0] <= i && i <= r[1]
}

func (r1 rng) intersect(r2 rng) (rng, bool) {
	if r1[0] <= r2[1] && r1[1] >= r2[0] {
		return rng{min(r1[0], r2[0]), max(r1[1], r2[1])}, true
	}
	return zeroRng, false
}

func getInput(data []byte) (rs []rng, ids []int, err error) {
	lines := bytes.Split(data, []byte("\n"))
	var idsStart int
	for i, line := range lines {
		if len(line) == 0 {
			idsStart = i
			break
		}
		brng := bytes.Split(line, []byte("-"))
		s, err := strconv.Atoi(string(brng[0]))
		if err != nil {
			return nil, nil, err
		}
		e, err := strconv.Atoi(string(brng[1]))
		if err != nil {
			return nil, nil, err
		}
		rs = append(rs, rng{s, e})
	}
	for _, line := range lines[idsStart+1:] {
		id, err := strconv.Atoi(string(line))
		if err != nil {
			return nil, nil, err
		}
		ids = append(ids, id)
	}
	return rs, ids, nil
}

func compact(rs []rng) []rng {
	var res []rng

	for i, r1 := range rs {
		if r1 == zeroRng {
			continue
		}
		for j, r2 := range rs[i+1:] {
			if r2 == zeroRng {
				continue
			}
			r, ok := r1.intersect(r2)
			if ok {
				r1 = r
				rs[i+j+1] = zeroRng
			}
		}
		res = append(res, r1)
	}

	if len(rs) != len(res) {
		return compact(res)
	}
	return res
}

func fresh(rs []rng, id int) bool {
	for _, r := range rs {
		if r.in(id) {
			return true
		}
	}
	return false
}
