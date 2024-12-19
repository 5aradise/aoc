package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/5aradise/aoc"
)

func main() {
	nums := getInput(aoc.InputFile)
	fmt.Println("Number of stones after 25 blinks (brute force):")
	fmt.Println(len(doBlinks(25, nums)))
	fmt.Println("Number of stones after 25 blinks (caching):")
	fmt.Println(doBlinks2(75, nums))
}

func getInput(inputFile string) []string {
	data, _ := os.ReadFile(inputFile)
	return strings.Split(string(data), " ")
}

func doBlinks(blinks int, stones []string) []string {
	for range blinks {
		stones = blink(stones)
	}
	return stones
}

func blink(stones []string) []string {
	new := make([]string, 0, len(stones))
	for _, stone := range stones {
		if stone == "0" {
			new = append(new, "1")
			continue
		}
		if len(stone)%2 == 0 {
			half := len(stone) / 2
			right := stone[half:]
			for i, r := range right {
				if r != '0' {
					right = right[i:]
					break
				}
			}
			if len(right) == len(stone)/2 && right[0] == '0' {
				right = "0"
			}
			new = append(new, stone[:half], right)
			continue
		}
		num, _ := strconv.Atoi(stone)
		num *= 2024
		new = append(new, strconv.Itoa(num))
	}
	return new
}

func Blink2(stones []int) []int {
	new := make([]int, 0, len(stones))
	for _, stone := range stones {
		if stone == 0 {
			new = append(new, 1)
			continue
		}
		stoneLen := len(strconv.Itoa(stone))
		if stoneLen%2 == 0 {
			del := 1
			for range stoneLen / 2 {
				del *= 10
			}
			new = append(new, stone/del, stone%del)
			continue
		}
		new = append(new, stone*2024)
	}
	return new
}

func doBlinks2(blinks int, stones []string) int {
	count := 0
	cache := make(map[[2]int]int)
	for _, stone := range stones {
		numStone, _ := strconv.Atoi(stone)
		count += getCount(cache, blinks, numStone)
	}
	return count
}

func getCount(cache map[[2]int]int, blinks, stone int) int {
	v, ok := cache[[2]int{blinks, stone}]
	if ok {
		return v
	}
	if blinks == 0 {
		return 1
	}
	if stone == 0 {
		if blinks < len(zeroStoneFuture) {
			return zeroStoneFuture[blinks]
		}
		c := getCount(cache, blinks-1, 1)
		cache[[2]int{blinks, stone}] = c
		return c
	}
	if stone == 2024 {
		if blinks == 1 {
			return 2
		}
		c := 2*getCount(cache, blinks-2, 2) + getCount(cache, blinks-2, 0) + getCount(cache, blinks-2, 4)
		cache[[2]int{blinks, stone}] = c
		return c
	}
	stoneLen := len(strconv.Itoa(stone))
	if stoneLen%2 == 0 {
		del := 1
		for range stoneLen / 2 {
			del *= 10
		}
		c := getCount(cache, blinks-1, stone/del) + getCount(cache, blinks-1, stone%del)
		cache[[2]int{blinks, stone}] = c
		return c
	}
	c := getCount(cache, blinks-1, stone*2024)
	cache[[2]int{blinks, stone}] = c
	return c
}

// trying to count the number of stones for each blink for 0
var zeroStoneFuture = []int{
	1, 1, 1, 2, 4, 4, 7,
	14, 16, 20, 39, 62, 81, 110, 200, 328, 418,
	667, 1059, 1546, 2377, 3572, 5602, 8268, 12343,
	19778, 29165, 43726, 67724, 102131, 156451, 234511,
	357632, 549949, 819967, 1258125, 1916299, 2886408, 4414216, 6669768,
	10174278, 15458147, 23333796, 35712308, 54046805, 81997335, 125001266,
	189148778, 288114305, 437102505, 663251546, 1010392024, 1529921658, 2327142660,
	3537156082,
}

func Calc0(blinks int) int {
	blinks -= 4
	count := 0
	res1 := []int{2}
	for i := range blinks {
		new := []int{}
		for _, stone := range res1 {
			if stone == 0 {
				count += 2 * zeroStoneFuture[blinks-i]
				continue
			}
			stoneLen := len(strconv.Itoa(stone))
			if stoneLen%2 == 0 {
				del := 1
				for range stoneLen / 2 {
					del *= 10
				}
				new = append(new, stone/del, stone%del)
				continue
			}
			new = append(new, stone*2024)
		}
		res1 = new
	}
	l := len(res1)
	runtime.GC()
	res3 := []int{4}
	for i := range blinks {
		new := []int{}
		for _, stone := range res3 {
			if stone == 0 {
				count += zeroStoneFuture[blinks-i]
				continue
			}
			stoneLen := len(strconv.Itoa(stone))
			if stoneLen%2 == 0 {
				del := 1
				for range stoneLen / 2 {
					del *= 10
				}
				new = append(new, stone/del, stone%del)
				continue
			}
			new = append(new, stone*2024)
		}
		res3 = new
	}
	return count + 2*l + len(res3) + zeroStoneFuture[blinks]
}
