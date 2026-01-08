package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	readString(reader)
	s := readString(reader)
	return solve(s)
}

func solve(s string) int {
	best := len(s)
	var flag int
	id := func(x byte) int {
		if x >= 'a' && x <= 'z' {
			return int(x - 'a')
		}
		return int(x-'A') + 26
	}
	n := len(s)
	for i := range n {
		flag |= 1 << id(s[i])
	}

	cnt := make([]int, 52)
	var cur int

	add := func(x byte) {
		i := id(x)
		cnt[i]++
		if cnt[i] == 1 {
			cur ^= 1 << i
		}
	}
	rem := func(x byte) {
		i := id(x)
		cnt[i]--
		if cnt[i] == 0 {
			cur ^= 1 << i
		}
	}

	for l, r := 0, 0; r < n; r++ {
		add(s[r])
		for cur == flag {
			rem(s[l])
			if cur != flag {
				add(s[l])
				break
			}
			l++
		}
		if cur == flag {
			best = min(best, r-l+1)
		}
	}

	return best
}
