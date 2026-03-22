package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) string {
	readString(reader)
	s := readString(reader)
	t := readString(reader)
	return solve(s, t)
}

func solve(s string, t string) string {
	n := len(s)

	freq := make([]int, 4)
	for i := range n {
		x := int(s[i] - '0')
		y := int(t[i] - '0')
		freq[x*2+y]++
	}

	var cnt int
	var player int

	// 如果是偶数，那么正好抵消掉
	if freq[3]%2 == 1 {
		cnt++
		player = 1
	}

	for freq[1] > 0 && freq[2] > 0 {
		if player == 0 {
			// 处理(1, 0)
			cnt++
			freq[2]--
		} else {
			cnt--
			freq[1]--
		}
		player ^= 1
	}
	if freq[2] > 0 {
		if player == 0 {
			cnt += (freq[2] + 1) / 2
		} else {
			cnt += freq[2] / 2
		}
	}
	if freq[1] > 0 {
		if player == 1 {
			cnt -= (freq[1] + 1) / 2
		} else {
			cnt -= freq[1] / 2
		}
	}
	// 00 不需要计算‘
	if cnt > 0 {
		return "First"
	}
	if cnt < 0 {
		return "Second"
	}
	return "Draw"
}
