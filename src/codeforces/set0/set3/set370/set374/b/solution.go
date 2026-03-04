package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(solve(s))
}

func solve(s string) int {
	// 理解错了，只能把sum = 9 的合并在一起
	n := len(s)

	ways := 1

	// 3636

	get := func(i int) int {
		return int(s[i] - '0')
	}

	for i := 0; i < n; i++ {
		if s[i] == '9' {
			continue
		}
		// 1818181
		j := i
		for i+1 < n && get(i)+get(i+1) == 9 {
			i++
		}

		if i == j {
			continue
		}
		if (i-j+1)%2 == 1 {
			ways *= (i - j + 2) / 2
		}
	}

	return ways
}
