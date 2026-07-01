package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var s string
	fmt.Fscan(reader, &s)
	return solve(s)
}

const inf = 1 << 60

func solve(s string) int {
	n := len(s)
	if n == 1 {
		return -1
	}
	if n == 2 {
		num, _ := strconv.Atoi(s)
		if num%25 == 0 {
			return 0
		}
		if s[1] != '0' {
			num, _ := strconv.Atoi(reverse(s))
			if num%25 == 0 {
				return 1
			}
		}
		return -1
	}

	play := func(last string) int {
		buf := []byte(s)

		var res int
		for i := n - 1; i >= 0; i-- {
			if buf[i] == last[1] {
				res += n - 1 - i
				for j := i; j+1 < n; j++ {
					buf[j], buf[j+1] = buf[j+1], buf[j]
				}
				break
			}
		}

		if buf[n-1] != last[1] {
			return inf
		}

		for i := n - 2; i >= 0; i-- {
			if buf[i] == last[0] {
				res += n - 2 - i
				// 避免0出现在第一个位置
				for j := i; j+1 < n-1; j++ {
					buf[j], buf[j+1] = buf[j+1], buf[j]
				}
				if buf[0] != '0' {
					// 必须找到第一个不为0的部分
					return res
				}

				for j := 1; j < n-2; j++ {
					if buf[j] != '0' {
						res += j
						return res
					}
				}

				return inf
			}
		}

		return inf
	}

	res := inf

	for _, w := range []int{0, 25, 50, 75} {
		last := fmt.Sprintf("%02d", w)
		res = min(res, play(last))
	}
	if res == inf {
		return -1
	}
	return res
}

func reverse(s string) string {
	buf := []byte(s)
	slices.Reverse(buf)
	return string(buf)
}
