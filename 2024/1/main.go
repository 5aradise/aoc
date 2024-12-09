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
	l1, l2 := getInput(aoc.InputFile)
	fmt.Println("Dist:")
	fmt.Println(calcDist(l1, l2))
	fmt.Println("Sim:")
	fmt.Println(calcSim(l1, l2))
}

func getInput(inputFile string) ([]int, []int) {
	data, _ := os.ReadFile(inputFile)
	rows := strings.Split(string(data), "\n")
	l1, l2 := make([]int, len(rows)), make([]int, len(rows))
	for i, row := range rows {
		nums := strings.Split(row, "   ")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		l1[i] = n1
		l2[i] = n2
	}
	return l1, l2
}

func calcDist(sl1 []int, sl2 []int) int {
	dist := 0
	slices.Sort(sl1)
	slices.Sort(sl2)
	for i := range sl1 {
		iDist := sl2[i] - sl1[i]
		if iDist < 0 {
			iDist = -iDist
		}
		dist += iDist
	}
	return dist
}

func calcSim(sl1 []int, sl2 []int) int {
	sim := 0
	appearsSl2 := make(map[int]int, len(sl2))
	for _, v := range sl2 {
		appearsSl2[v]++
	}
	for _, v := range sl1 {
		sim += v * appearsSl2[v]
	}
	return sim
}
