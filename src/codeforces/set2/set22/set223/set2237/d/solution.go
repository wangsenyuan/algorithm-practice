package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int64 {
	var n int
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)
	return solve(s)
}

func solve(s string) int64 {
	cnt := make([]int, 3)
	var res int
	var y int
	var x int
	cnt[0]++
	for i := range len(s) {
		if s[i] == '0' {
			x += 2
		} else {
			x++
		}
		x %= 3
		res += cnt[0] + cnt[1] + cnt[2] - cnt[x]
		if i == 0 || s[i] != s[i-1] {
			y++
		} else {
			y = 1
		}
		res -= (y - 1) / 2
		cnt[x]++
	}
	return int64(res)
}
