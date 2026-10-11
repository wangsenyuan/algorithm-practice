package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, _, res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) (s string, k int, ans string) {
	var n int
	fmt.Fscan(reader, &n, &k, &s)
	_ = n
	ans = solve(k, s)
	return
}

const inf = 1 << 30

func solve(k int, s string) string {
	n := len(s)
	// 直接brute force就可以了
	var left []int
	var right []int
	for i := range n {
		if s[i] == '(' {
			left = append(left, i)
		} else {
			right = append(right, i)
		}
	}
	slices.Reverse(right)

	k = min(k, len(left), len(right))

	left = left[:k]
	right = right[:k]
	slices.Reverse(right)

	buf := make([]byte, n)
	for i := range n {
		buf[i] = '0'
	}
	for _, i := range right[:k] {
		buf[i] = '1'
	}

	check := func() int {
		var ans int
		var level int
		for i := range n {
			if buf[i] == '0' {
				if s[i] == '(' {
					level++
				} else {
					// s[i] == ')'
					if level > 0 {
						ans += 2
					}
					level = max(0, level-1)
				}
			}
		}
		return ans
	}

	best := check()
	ans := string(buf)

	for _, i := range left {
		if len(right) > 0 {
			j := right[0]
			right = right[1:]
			buf[j] = '0'
		}
		buf[i] = '1'
		tmp := check()
		if tmp < best {
			best = tmp
			ans = string(buf)
		}
	}

	return ans
}
