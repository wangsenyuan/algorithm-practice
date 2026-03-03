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
	res := solve(s)
	fmt.Println(res)
}

func solve(s string) string {
	// n := len(s)
	// 那么对前面的数字随便排列，然后看需要多少？
	// 只保留一个 1689
	freq := make([]int, 10)
	for _, c := range s {
		freq[int(c-'0')]++
	}
	freq[1]--
	freq[6]--
	freq[8]--
	freq[9]--

	leadingZero := true
	var res []byte
	var sum int
	for d := 9; d >= 0; d-- {
		if freq[d] == 0 {
			continue
		}
		if d > 0 {
			leadingZero = false
		}

		for freq[d] > 0 {
			sum = sum*10 + d
			sum %= 7
			res = append(res, byte(d+'0'))
			freq[d]--
		}
	}
	if leadingZero {
		// 把0都放置到后面去
		return "1869" + string(res)
	}
	sum = (sum * 10000) % 7
	if sum != 0 {
		sum = 7 - sum
	}
	suf := []string{
		"1869",
		"6189",
		"1689",
		"6891",
		"8691",
		"9861",
		"1896",
	}
	return string(res) + suf[sum]
}
