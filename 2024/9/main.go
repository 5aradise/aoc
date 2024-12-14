package main

import (
	"fmt"
	"os"

	"github.com/5aradise/aoc"
)

func main() {
	mp1 := getInput(aoc.InputFile)
	mp2 := make([]int, len(mp1))
	copy(mp2, mp1)
	fmt.Println("Defragment and compact by blocks:")
	fmt.Println(checksum(defragmentAndCompactByBlocks(mp1)))
	fmt.Println("Defragment and compact by files:")
	fmt.Println(checksum(defragmentAndCompactByFiles(mp2)))
}

var figures = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func getInput(inputFile string) []int {
	data, _ := os.ReadFile(inputFile)
	nums := make([]int, len(data))
	for i, numStrB := range data {
		nums[i] = figures[numStrB]
	}
	return nums
}

func defragmentAndCompactByBlocks(mp []int) []int {
	var res []int
	i, j := 0, len(mp)-1
	for i < j {
		for range mp[i] {
			res = append(res, i/2)
		}
		i++
		for range mp[i] {
			for mp[j] == 0 {
				j -= 2
			}
			if j < i {
				j += 2
				break
			}
			res = append(res, j/2)
			mp[j]--
		}
		i++
	}
	for range mp[j] {
		res = append(res, j/2)
	}
	return res
}

func defragmentAndCompactByFiles(mp []int) []int {
	var res []int
	i := 0
	for i < len(mp)-1 {
		currToFill := mp[i]
		if currToFill >= 0 {
			for range currToFill {
				res = append(res, i/2)
			}
		} else {
			for range -currToFill {
				res = append(res, 0)
			}
		}
		i++
		currEmpty := mp[i]
		for j := len(mp) - 1; j > i && currEmpty > 0; j -= 2 {
			toFill := mp[j]
			if toFill > 0 && toFill <= currEmpty {
				mp[j] = -toFill
				currEmpty -= toFill
				for range toFill {
					res = append(res, j/2)
				}
			}
		}
		for range currEmpty {
			res = append(res, 0)
		}
		i++
	}
	currToFill := mp[i]
	if currToFill >= 0 {
		for range currToFill {
			res = append(res, i/2)
		}
	} else {
		for range -currToFill {
			res = append(res, 0)
		}
	}
	return res
}

func checksum(nums []int) int {
	sum := 0
	for i, num := range nums {
		sum += i * num
	}
	return sum
}
