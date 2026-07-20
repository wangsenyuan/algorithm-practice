package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var n int
	var s int64
	fmt.Fscan(reader, &n, &s)

	a := make([]int64, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	return solve(s, a)
}

func solve(s int64, a []int64) string {
	var sum int64
	for _, v := range a {
		sum += v
	}

	s1 := s % sum
	// n := len(a)
	if s1 == 0 {
		return "Yes"
	}

	a = append(a, a...)

	var tmp int64
	pref := make(map[int64]int)
	pref[0] = 0
	for i, v := range a {
		tmp += v
		if _, ok := pref[tmp-s1]; ok {
			return "Yes"
		}

		pref[tmp] = i + 1
	}

	return "No"
}
