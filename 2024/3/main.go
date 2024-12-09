package main

import (
	"fmt"
	"os"

	"github.com/5aradise/aoc"
)

func main() {
	input := getInput(aoc.InputFile)
	fmt.Println("Muls:")
	fmt.Println(aoc.Sum(parseAndCalcMuls(input)))
	fmt.Println("Muls with conds:")
	fmt.Println(aoc.Sum(parseAndCalcMulsWithConds(input)))
}

func getInput(inputFile string) string {
	data, _ := os.ReadFile(inputFile)
	return string(data)
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

func parseAndCalcMuls(s string) []int {
	var muls []int
	for i := 0; i < len(s); i++ {
		if s[i] == 'm' {
			if s[i+1:i+4] == "ul(" {
				num1 := 0
				for j := i + 4; ; j++ {
					char := s[j]
					figure, ok := figures[char]
					if ok {
						num1 = num1*10 + figure
					} else if char == ',' {
						num2 := 0
						for k := j + 1; ; k++ {
							char := s[k]
							figure, ok := figures[char]
							if ok {
								num2 = num2*10 + figure
							} else if char == ')' {
								mul := num1 * num2
								muls = append(muls, mul)
								i = k
								break
							} else {
								i = k - 1
								break
							}
						}
						break
					} else {
						i = j - 1
						break
					}
				}
			}
		}
	}
	return muls
}

func parseAndCalcMulsWithConds(s string) []int {
	var muls []int
	doCalc := true
	for i := 0; i < len(s); i++ {
		if s[i] == 'm' {
			if doCalc {
				if s[i+1:i+4] == "ul(" {
					num1 := 0
					for j := i + 4; ; j++ {
						char := s[j]
						figure, ok := figures[char]
						if ok {
							num1 = num1*10 + figure
						} else if char == ',' {
							num2 := 0
							for k := j + 1; ; k++ {
								char := s[k]
								figure, ok := figures[char]
								if ok {
									num2 = num2*10 + figure
								} else if char == ')' {
									mul := num1 * num2
									muls = append(muls, mul)
									i = k
									break
								} else {
									i = k - 1
									break
								}
							}
							break
						} else {
							i = j - 1
							break
						}
					}
				}
			}
		} else if s[i] == 'd' {
			if s[i+1:i+4] == "o()" {
				doCalc = true
				i += 3
			} else if s[i+1:i+7] == "on't()" {
				doCalc = false
				i += 6
			}
		}
	}
	return muls
}
