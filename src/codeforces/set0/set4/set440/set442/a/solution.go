package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	cards := strings.Split(s, " ")
	res := solve(cards)
	fmt.Println(res)
}

func solve(cards []string) int {
	colors := map[byte]int{
		'R': 0,
		'G': 1,
		'B': 2,
		'Y': 3,
		'W': 4,
	}

	nums := map[byte]int{
		'1': 0,
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
	}

	slices.Sort(cards)

	cards = slices.Compact(cards)

	n := len(cards)

	check := func(s1 int, s2 int) bool {
		for i := range n {
			for j := range i {
				ok := false
				if cards[i][0] != cards[j][0] {
					if (s1>>colors[cards[i][0]])&1 > 0 || (s1>>colors[cards[j][0]])&1 > 0 {
						ok = true
					}
				}
				if cards[i][1] != cards[j][1] {
					if (s2>>nums[cards[i][1]])&1 > 0 || (s2>>nums[cards[j][1]])&1 > 0 {
						ok = true
					}
				}
				if !ok {
					return false
				}
			}
		}
		return true
	}

	res := 1 << 30
	for s1 := 0; s1 < 1<<5; s1++ {
		for s2 := 0; s2 < 1<<5; s2++ {
			if check(s1, s2) {
				res = min(res, bits.OnesCount(uint(s1))+bits.OnesCount(uint(s2)))
			}
		}
	}
	return res
}
