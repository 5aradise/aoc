package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/5aradise/aoc"
)

func main() {
	input := getInput(aoc.InputFile)
	fmt.Println("Count XMAS:")
	fmt.Println(countXMAS(input))
	fmt.Println("Count X-MAS:")
	fmt.Println(countXtMAS(input))
}

func getInput(inputFile string) []string {
	data, _ := os.ReadFile(inputFile)
	return strings.Split(string(data), "\n")
}

func countXMAS(rows []string) int {
	count := 0
	hight := len(rows)
	length := len(rows[0])
	doTryDiagonalL := false
	for i := range hight {
		for j := range length {
			r := rows[i][j]
			if r == 'X' {
				if length-j >= 4 {
					if rows[i][j+1:j+4] == "MAS" {
						count++
					}
					doTryDiagonalL = true
				}
				if hight-i >= 4 {
					if rsToStr(rows[i+1][j], rows[i+2][j], rows[i+3][j]) == "MAS" {
						count++
					}
					if j >= 3 {
						if rsToStr(rows[i+1][j-1], rows[i+2][j-2], rows[i+3][j-3]) == "MAS" {
							count++
						}
					}
					if doTryDiagonalL {
						if rsToStr(rows[i+1][j+1], rows[i+2][j+2], rows[i+3][j+3]) == "MAS" {
							count++
						}
					}
				}
			} else if r == 'S' {
				if length-j >= 4 {
					if rows[i][j+1:j+4] == "AMX" {
						count++
					}
					doTryDiagonalL = true
				}
				if hight-i >= 4 {
					if rsToStr(rows[i+1][j], rows[i+2][j], rows[i+3][j]) == "AMX" {
						count++
					}
					if j >= 3 {
						if rsToStr(rows[i+1][j-1], rows[i+2][j-2], rows[i+3][j-3]) == "AMX" {
							count++
						}
					}
					if doTryDiagonalL {
						if rsToStr(rows[i+1][j+1], rows[i+2][j+2], rows[i+3][j+3]) == "AMX" {
							count++
						}
					}
				}
			}
			doTryDiagonalL = false
		}
	}
	return count
}

var win = []byte{0, 0, 0}

func rsToStr(a, b, c byte) string {
	win[0] = a
	win[1] = b
	win[2] = c
	return string(win)
}

func countXtMAS(rows []string) int {
	count := 0
	hight := len(rows)
	length := len(rows[0])
	for i := 1; i < hight-1; i++ {
		for j := 1; j < length-1; j++ {
			if rows[i][j] == 'A' {
				tl := rows[i-1][j-1]
				tr := rows[i-1][j+1]
				bl := rows[i+1][j-1]
				br := rows[i+1][j+1]
				if (tl == 'M' && br == 'S' || tl == 'S' && br == 'M') && (tr == 'M' && bl == 'S' || tr == 'S' && bl == 'M') {
					count++
				}
			}
		}
	}
	return count
}
