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

	ps, err := getInput(data)
	if err != nil {
		panic(err)
	}
	cmps, err := getCephalopodInput(data)
	if err != nil {
		panic(err)
	}

	gt := aoc.Reduce(ps, 0, func(sum int, p problem) int {
		return sum + p.solve()
	})
	cmgt := aoc.Reduce(cmps, 0, func(sum int, p problem) int {
		return sum + p.solve()
	})

	fmt.Println(aoc.FormatAnswers(gt, cmgt))
}

type operation int

const (
	_ operation = iota
	add
	mul
)

func parseOp(b byte) (operation, error) {
	switch b {
	case '+':
		return add, nil
	case '*':
		return mul, nil
	default:
		return 0, errors.New("invalid operation char")
	}
}

type problem struct {
	op   operation
	nums []int
}

func (p problem) solve() int {
	res := p.nums[0]
	switch p.op {
	case add:
		for _, num := range p.nums[1:] {
			res += num
		}
	case mul:
		for _, num := range p.nums[1:] {
			res *= num
		}
	default:
		panic("unknown operation")
	}
	return res
}

func getInput(data []byte) ([]problem, error) {
	lines := bytes.Split(data, []byte("\n"))
	numLines, opLine := lines[:len(lines)-1], lines[len(lines)-1]

	ops := bytes.Fields(opLine)
	ps := make([]problem, 0, len(ops))
	for _, op := range ops {
		op, err := parseOp(op[0])
		if err != nil {
			return nil, err
		}
		ps = append(ps, problem{
			op:   op,
			nums: make([]int, 0, len(numLines)),
		})
	}

	for _, line := range numLines {
		bnums := bytes.Fields(line)
		for i, bnum := range bnums {
			num, err := strconv.Atoi(string(bnum))
			if err != nil {
				return nil, err
			}
			ps[i].nums = append(ps[i].nums, num)
		}
	}

	return ps, nil
}

func getCephalopodInput(data []byte) ([]problem, error) {
	lines := bytes.Split(data, []byte("\n"))
	numLines, opLine := lines[:len(lines)-1], lines[len(lines)-1]

	var ps []problem
	for i, v := range opLine {
		op, err := parseOp(v)
		if err == nil {
			ps = append(ps, problem{
				op: op,
			})
		}

		var bnum []byte
		for _, line := range numLines {
			b := line[i]
			if b != ' ' {
				bnum = append(bnum, b)
			}
		}
		if len(bnum) == 0 {
			continue
		}
		num, err := strconv.Atoi(string(bnum))
		if err != nil {
			return nil, err
		}
		ps[len(ps)-1].nums = append(ps[len(ps)-1].nums, num)
	}
	return ps, nil
}
