package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/5aradise/aoc"
)

func main() {
	order, seqs := getInput(aoc.InputFile)
	fmt.Println("Valid sequences:")
	fmt.Println(sumSlMid(getValidSeqs(order, seqs)))
	fmt.Println("Fixed sequences:")
	fmt.Println(sumSlMid(fixedInvalidSeqs(order, seqs)))
}

func getInput(inputFile string) (map[int][]int, [][]int) {
	order := make(map[int][]int)
	seqs := make([][]int, 0, 128)

	data, _ := os.ReadFile(inputFile)
	rows := strings.Split(string(data), "\n")
	j := 0
	for i, row := range rows {
		if row == "" {
			j = i
			break
		}
		nums := strings.Split(row, "|")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		order[n1] = append(order[n1], n2)
	}
	for _, row := range rows[j+1:] {
		strs := strings.Split(row, ",")
		nums := make([]int, len(strs))
		for i, str := range strs {
			num, _ := strconv.Atoi(str)
			nums[i] = num
		}
		seqs = append(seqs, nums)
	}
	return order, seqs
}

func sumSlMid(sls [][]int) int {
	sum := 0
	for _, sl := range sls {
		sum += sl[len(sl)/2]
	}
	return sum
}

func getValidSeqs(order map[int][]int, seqs [][]int) [][]int {
	var valid [][]int
	for _, seq := range seqs {
		if isValidSeq(order, seq) {
			valid = append(valid, seq)
		}
	}
	return valid
}

func fixedInvalidSeqs(order map[int][]int, seqs [][]int) [][]int {
	var fixeds [][]int
	for _, seq := range seqs {
		if !isValidSeq(order, seq) {
			slices.SortFunc(seq, comparaByOrder(order))
			fixeds = append(fixeds, seq)
		}
	}
	return fixeds
}

func isValidSeq(order map[int][]int, seq []int) bool {
	return slices.IsSortedFunc(seq, comparaByOrder(order))
}

func comparaByOrder(order map[int][]int) func(a, b int) int {
	return func(a, b int) int {
		if slices.Contains(order[a], b) {
			return -1
		} else if slices.Contains(order[b], a) {
			return 1
		}
		return 0
	}
}
