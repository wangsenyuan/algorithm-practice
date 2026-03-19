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

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	nums := make([]int, len(ss))
	for i, s := range ss {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}

func drive(reader *bufio.Reader) int {
	readString(reader)
	a := readNums(reader)
	m := readString(reader)
	return solve(a, m)
}

func solve(a []int, m string) int {
	n := len(a)
	var pref int
	for _, v := range a {
		pref += v
	}

	var best int
	var suf int
	for i := n - 1; i >= 0; i-- {
		pref -= a[i]
		if m[i] == '1' {
			// 这一位如果为0, 就能满足 x < m
			best = max(best, pref+suf)
			suf += a[i]
		}
	}
	best = max(best, suf)
	return best
}
