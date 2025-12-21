package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)

	if n < 3 {
		return 0
	}

	var sum int
	for _, v := range a {
		sum += v
	}
	if sum%3 != 0 {
		return 0
	}

	s1 := sum / 3

	freq := make(map[int]int)
	// freq[0]++

	var res int

	var s2 int
	for i, v := range a {
		s2 += v
		if s2 == 2*s1 && i < n-1 {
			res += freq[s1]
		}
		freq[s2]++
	}

	return res
}
