package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/5aradise/aoc"
)

func main() {
	input := getInput(aoc.InputFile)
	fmt.Println("Total calibration result (+ *):")
	fmt.Println(totalCalib(input, isValidCalib2))
	fmt.Println("Total calibration result (+ * ||):")
	fmt.Println(totalCalib(input, isValidCalib3))
}

func getInput(inputFile string) map[int][]int {
	data, _ := os.ReadFile(inputFile)
	rows := strings.Split(string(data), "\n")
	input := make(map[int][]int, len(rows))
	for _, row := range rows {
		splt := strings.Split(row, ":")
		n, _ := strconv.Atoi(splt[0])
		numsStr := strings.Split(splt[1], " ")[1:]
		nums := make([]int, len(numsStr))
		for i, numStr := range numsStr {
			num, _ := strconv.Atoi(numStr)
			nums[i] = num
		}
		input[n] = nums
	}
	return input
}

func totalCalib(calibs map[int][]int, calibChecker func(res int, seq []int) bool) int {
	total := 0
	for num, seq := range calibs {
		if calibChecker(num, seq) {
			total += num
		}
	}
	return total
}

func isValidCalib2(res int, seq []int) bool {
	if len(seq) == 0 {
		return res == 0
	}
	if res <= 0 {
		return false
	}
	newSeq, last := seq[:len(seq)-1], seq[len(seq)-1]
	div := float64(res) / float64(last)
	intDiv := int(div)
	if div == float64(intDiv) && isValidCalib2(intDiv, newSeq) {
		return true
	}
	return isValidCalib2(res-last, newSeq)
}

func isValidCalib3(res int, seq []int) bool {
	if len(seq) == 0 {
		return res == 0
	}
	if res <= 0 {
		return false
	}
	newSeq, last := seq[:len(seq)-1], seq[len(seq)-1]
	n := getLongerSmallest(last)
	if res%n == last && isValidCalib3(res/n, newSeq) {
		return true
	}
	div := float64(res) / float64(last)
	intDiv := int(div)
	if div == float64(intDiv) && isValidCalib3(intDiv, newSeq) {
		return true
	}
	return isValidCalib3(res-last, newSeq)
}

func getLongerSmallest(num int) int {
	curr := 1
	for len := 1; ; len++ {
		curr *= 10
		if num < curr {
			return curr
		}
	}
}
