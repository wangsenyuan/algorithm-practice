package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Println(res)
}

func solve(s string) int {
	// 3个字符，最多时 9 + 9 + 9 = 27
	// 最小是0
	// 那就迭代两边期望的结果，然后算一下
	best := inf
	for x := range 28 {
		best = min(best, calc(s[:3], x)+calc(s[3:], x))
	}
	return best
}

func calc(s string, expect int) int {
	var sum int
	buf := make([]int, len(s))
	for i := range len(s) {
		buf[i] = int(s[i] - '0')
		sum += buf[i]
	}

	if sum == expect {
		return 0
	}

	sort.Ints(buf)

	var res int

	if sum < expect {
		for i := range len(s) {
			x := min(expect-sum, 9-buf[i])
			sum += x
			res++
			if sum == expect {
				break
			}
		}
	} else {
		for i := len(s) - 1; i >= 0; i-- {
			x := min(sum-expect, buf[i])
			sum -= x
			res++
			if sum == expect {
				break
			}
		}
	}

	if sum != expect {
		return inf
	}

	return res
}

const inf = 1 << 60
