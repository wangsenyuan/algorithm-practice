package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int64 {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) int64 {
	freq := make([]int, k)
	freq[0]++
	var res int
	var s1, s2 int
	for i, v := range a {
		s2 = (s2 + v) % k
		if i >= k {
			res += freq[s2]
			s1 = (s1 + a[i-k]) % k
			freq[s1]++
		}
	}
	return int64(res)
}
