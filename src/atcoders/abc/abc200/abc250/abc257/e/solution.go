package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	c := make([]int, 9)
	for i := range c {
		fmt.Fscan(reader, &c[i])
	}
	return solve(n, c)
}

const inf = 1 << 60

func solve(n int, c []int) string {
	var best int
	for i := range 9 {
		if c[i] <= c[best] {
			best = i
		}
	}

	cnt := n / c[best]
	buf := make([]int, cnt)

	for i := range cnt {
		buf[i] = best
	}

	suf := c[best] * cnt
	var pref int
	for i := range cnt {
		suf -= c[best]
		cur := 8
		for cur > best {
			if pref+c[cur]+suf <= n {
				break
			}
			cur--
		}

		if cur == best {
			break
		}
		buf[i] = cur
		pref += c[cur]
	}

	var res strings.Builder
	for i := range cnt {
		res.WriteString(strconv.Itoa(buf[i] + 1))
	}
	return res.String()
}
