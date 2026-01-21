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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	s := readString(reader)
	tc, _ := strconv.Atoi(s)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	readString(reader)
	s := readString(reader)
	return solve(s)
}

func solve(s string) int {
	// s中如果出现1，那么这些1是连续的
	// 假设固定r, [l...r] = 1, 那么就是找到l，使的l前面的都变成0，l后面的都变成1，
	// r - (l-1)- (pref[r] - pref[l-1]) +  pref[l-1] + suf[r+1] 最小
	// r - (l-1) - pref[r] + pref[l-1] + pref[l-1] + suf[r+1]
	var suf int
	n := len(s)
	for i := range n {
		if s[i] == '1' {
			suf++
		}
	}
	// 全部变成0，或者全部变成1
	ans := min(suf, n-suf)
	var best int
	var pref int
	for r := range n {
		if s[r] == '1' {
			suf--
			pref++
		}
		ans = min(ans, r+1-pref+suf+best)
		best = min(best, -(r+1)+2*pref)
	}

	return ans
}
