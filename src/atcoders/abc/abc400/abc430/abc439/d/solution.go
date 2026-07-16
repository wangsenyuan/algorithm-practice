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

func drive(reader *bufio.Reader) int64 {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int64 {
	freq := make(map[int]int)

	var res int

	for _, w := range a {
		// a[i] / w = 7 / 5
		// a[i] = 7 * w / 5
		// w / a[k] = 5 / 3
		if 7*w%5 == 0 && 3*w%5 == 0 {
			res += freq[7*w/5] * freq[3*w/5]
		}
		freq[w]++
	}

	clear(freq)

	for i := len(a) - 1; i >= 0; i-- {
		w := a[i]
		if 7*w%5 == 0 && 3*w%5 == 0 {
			res += freq[7*w/5] * freq[3*w/5]
		}
		freq[w]++
	}

	return int64(res)
}
