package main

import (
	"bytes"
	"fmt"
	"iter"
	"os"
	"strconv"

	"github.com/5aradise/aoc"
)

func main() {
	data, err := os.ReadFile(aoc.InputFile)
	if err != nil {
		panic(err)
	}

	rngs, err := getInput(data)
	if err != nil {
		panic(err)
	}

	sum1 := invalidIDSum(rngs, invalidID1)
	sum2 := invalidIDSum(rngs, invalidID2)

	fmt.Println(aoc.FormatAnswers(sum1, sum2))
}

func getInput(data []byte) (ranges []rnge, err error) {
	for rng := range bytes.SplitSeq(data, []byte(",")) {
		nums := bytes.SplitN(rng, []byte("-"), 2)
		sf, sl := string(nums[0]), string(nums[1])

		f, err := strconv.Atoi(sf)
		if err != nil {
			return nil, err
		}
		l, err := strconv.Atoi(sl)
		if err != nil {
			return nil, err
		}

		ranges = append(ranges, rnge{f, l})
	}
	return ranges, nil
}

func invalidIDSum(rngs []rnge, invalidID func(string) bool) int {
	var sum int
	for _, rng := range rngs {
		for id := range rng.iter() {
			sid := strconv.Itoa(id)
			if invalidID(sid) {
				sum += id
			}
		}
	}
	return sum
}

func invalidID1(id string) bool {
	if len(id)%2 != 0 {
		return false
	}

	return id[:len(id)/2] == id[len(id)/2:]
}

func invalidID2(id string) bool {
nextSubs:
	for i := 1; i <= len(id)/2; i++ {
		if len(id)%i != 0 {
			continue
		}

		subs := id[:i]

		rest := id[i:]
		for range (len(id) / i) - 1 {
			if rest[:i] != subs {
				continue nextSubs
			}
			rest = rest[i:]
		}
		return true
	}
	return false
}

type rnge [2]int

func (r rnge) iter() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := r[0]; i <= r[1]; i++ {
			if !yield(i) {
				return
			}
		}
	}
}
