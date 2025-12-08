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
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	t := readString(reader)
	x, _ := strconv.Atoi(t)
	s := readString(reader)
	return solve(x, s)
}

func solve(t int, s string) int {
	// 2和5算一组
	// 6 & 9
	cnt := make([]int, 10)
	for _, c := range s {
		cnt[c-'0']++
	}
	cnt[2] += cnt[5]
	cnt[5] = 0
	cnt[6] += cnt[9]
	cnt[9] = 0

	w := make([]int, 10)
	for i := t; i > 0; i /= 10 {
		w[i%10]++
	}
	w[2] += w[5]
	w[5] = 0
	w[6] += w[9]
	w[9] = 0

	ans := inf

	for i := range 10 {
		if w[i] == 0 {
			continue
		}
		ans = min(ans, cnt[i]/w[i])
	}

	return ans
}

const inf = 1 << 60
