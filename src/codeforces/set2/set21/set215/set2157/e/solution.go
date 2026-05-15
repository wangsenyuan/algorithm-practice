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

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) int {
	n := len(a)
	freq := make([]int, 4*n+1)

	ok := true
	for _, v := range a {
		freq[v]++
		if freq[v] > k {
			ok = false
			break
		}
	}
	if ok {
		return 0
	}

	play := func(m int) bool {
		clear(freq)
		for _, v := range a {
			freq[v]++
		}
		var zeros []int
		for i := 4 * n; i >= 0; i-- {
			if freq[i] == 0 {
				zeros = append(zeros, i)
			}
			for freq[i] > 1 && len(zeros) > 0 && last(zeros) <= i+m {
				freq[i]--
				freq[last(zeros)]++
				zeros = zeros[:len(zeros)-1]
			}

			if freq[i] > 1 {
				freq[i+m] += freq[i] - 1
				freq[i] = 1
			}
		}
		for _, x := range freq {
			if x > k {
				return false
			}
		}
		return true
	}
	var res int

	for i := 20; i >= 0; i-- {
		if !play(res + 1<<i) {
			res += 1 << i
		}
	}

	return res + 1
}

func last(a []int) int {
	return a[len(a)-1]
}
