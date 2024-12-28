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
	makeSafetyRep(input)
}

func getInput(inputFile string) [][]int {
	data, _ := os.ReadFile(inputFile)
	rows := strings.Split(string(data), "\n")
	input := make([][]int, len(rows))
	for i, row := range rows {
		strlvls := strings.Split(row, " ")
		lvls := make([]int, len(strlvls))
		for i, strlvl := range strlvls {
			lvl, _ := strconv.Atoi(strlvl)
			lvls[i] = lvl
		}
		input[i] = lvls
	}
	return input
}

func makeSafetyRep(reps [][]int) {
	count := 0
	countT := 0
	for _, rep := range reps {
		isSafe := isRepSafe(rep)
		isSafeT := isRepSafeT(rep)
		if isSafe {
			count++
		}
		if isSafeT {
			countT++
		}
		if isSafe != isSafeT {
			if isSafe {
				fmt.Print("✔️ ")
			} else {
				fmt.Print("❌")
			}
			fmt.Print(" ")
			if isSafeT {
				fmt.Print("✔️ ")
			} else {
				fmt.Print("❌")
			}
			fmt.Printf(" %v\n", rep)
		}
	}
	fmt.Printf("Safe: %d\n", count)
	fmt.Printf("Tolerantly safe: %d\n", countT)
}

func isRepSafe(rep []int) bool {
	diff := rep[1] - rep[0]
	if diff == 0 || !(-3 <= diff && diff <= 3) {
		return false
	}
	if diff > 0 {
		for i := 1; i < len(rep)-1; i++ {
			diff = rep[i+1] - rep[i]
			if !(1 <= diff && diff <= 3) {
				return false
			}
		}
	} else {
		for i := 1; i < len(rep)-1; i++ {
			diff = rep[i+1] - rep[i]
			if !(-3 <= diff && diff <= -1) {
				return false
			}
		}
	}
	return true
}

func isRepSafeT(rep []int) bool {
	isRemoved := false
	popularDiffN := 0
	for i := range rep[:3] {
		diff := rep[i+1] - rep[i]
		if diff > 0 {
			popularDiffN += 1
		} else {
			popularDiffN += -1
		}
	}
	isInc := false
	if popularDiffN > 0 {
		isInc = true
	} else if popularDiffN == 0 {
		return false
	}
	if isInc {
		for i := 0; i < len(rep)-1; i++ {
			diff := rep[i+1] - rep[i]
			if !(1 <= diff && diff <= 3) {
				if isRemoved {
					return false
				}
				if i == len(rep)-2 {
					return true
				}
				if i != 0 {
					diff11 := rep[i+1] - rep[i-1]
					diff12 := rep[i+2] - rep[i+1]
					if !(1 <= diff11 && diff11 <= 3) || !(1 <= diff12 && diff12 <= 3) {
						diff21 := rep[i] - rep[i-1]
						diff22 := rep[i+2] - rep[i]
						if !(1 <= diff21 && diff21 <= 3) || !(1 <= diff22 && diff22 <= 3) {
							return false
						}
					}
				} else {
					diff1 := rep[2] - rep[0]
					diff2 := rep[2] - rep[1]
					if !(1 <= diff1 && diff1 <= 3) && !(1 <= diff2 && diff2 <= 3) {
						return false
					}
				}
				i++
				isRemoved = true
			}
		}
	} else {
		for i := 0; i < len(rep)-1; i++ {
			diff := rep[i+1] - rep[i]
			if !(-3 <= diff && diff <= -1) {
				if isRemoved {
					return false
				}
				if i == len(rep)-2 {
					return true
				}
				if i != 0 {
					diff11 := rep[i+1] - rep[i-1]
					diff12 := rep[i+2] - rep[i+1]
					if !(-3 <= diff11 && diff11 <= -1) || !(-3 <= diff12 && diff12 <= -1) {
						diff21 := rep[i] - rep[i-1]
						diff22 := rep[i+2] - rep[i]
						if !(-3 <= diff21 && diff21 <= -1) || !(-3 <= diff22 && diff22 <= -1) {
							return false
						}
					}

				} else {
					diff1 := rep[2] - rep[0]
					diff2 := rep[2] - rep[1]
					if !(-3 <= diff1 && diff1 <= -1) && !(-3 <= diff2 && diff2 <= -1) {
						return false
					}
				}
				i++
				isRemoved = true
			}
		}
	}
	return true
}
