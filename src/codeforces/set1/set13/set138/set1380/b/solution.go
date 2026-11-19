package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var s string
	fmt.Fscan(reader, &s)
	return solve(s)
}

func solve(s string) string {
	// c[1] 需要和s[1], s[2], ... s[n] 比较
	// 那么如果C[1]是R, 那么它的贡献 = S的个数 - P的个数. C[?] 似乎都是一样的逻辑
	freq := make([]int, 3)
	for i := range s {
		switch s[i] {
		case 'R':
			freq[0]++
		case 'S':
			freq[1]++
		default:
			freq[2]++
		}
	}

	score := make([]int, 3)
	score[0] = freq[1]
	score[1] = freq[2]
	score[2] = freq[0]

	w := slices.Max(score)
	index := slices.Index(score, w)

	letter := "RSP"[index]

	return strings.Repeat(string(letter), len(s))
}
