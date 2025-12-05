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

	banks, err := getInput(data)
	if err != nil {
		panic(err)
	}

	sum2 := aoc.Reduce(banks, 0, func(sum int, bank []int) int {
		return sum + largestJoltage2(bank)
	})
	sum12 := aoc.Reduce(banks, 0, func(sum int, bank []int) int {
		return sum + largestJoltage12(bank)
	})

	fmt.Println(aoc.FormatAnswers(sum2, sum12))
}

func getInput(data []byte) (banks [][]int, err error) {
	for line := range bytes.Lines(data) {
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		bank := make([]int, 0, len(line))
		for _, num := range line {
			bat, err := strconv.Atoi(string(num))
			if err != nil {
				return nil, err
			}
			bank = append(bank, bat)
		}
		banks = append(banks, bank)
	}
	return banks, nil
}

func largestJoltage2(bank []int) int {
	var (
		fbat = bank[len(bank)-2]
		sbat = bank[len(bank)-1]
	)
	for i := len(bank) - 3; i >= 0; i-- {
		bat := bank[i]
		if bat >= fbat {
			if fbat > sbat {
				sbat = fbat
			}
			fbat = bat
		}
	}
	return fbat*10 + sbat
}

func largestJoltage12(bank []int) int {
	bats := bank[len(bank)-12:]
	for i := len(bank) - 13; i >= 0; i-- {
		insertDigit(bats, bank[i])
	}
	return digitsToNum(bats)
}

func insertDigit(num []int, dig int) {
	if len(num) > 0 && dig >= num[0] {
		insertDigit(num[1:], num[0])
		num[0] = dig
	}
}

func digitsToNum(digits []int) int {
	var (
		num   = 0
		pow10 = 1
	)

	for i := len(digits) - 1; i >= 0; i-- {
		dig := digits[i]
		num += dig * pow10
		pow10 *= 10
	}
	return num
}
