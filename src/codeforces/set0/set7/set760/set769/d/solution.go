package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

func solve(a []int, k int) int {
	// a[i] <= 1e4 ~ 2^14

	mx := slices.Max(a)

	freq := make([]int, mx+1)
	// C(14, k)
	var ans int

	for _, v := range a {
		if k == 0 {
			ans += freq[v]
		}
		freq[v]++
	}

	if k == 0 {
		return ans
	}

	next := func(x int) int {
		i := 1
		for (x>>i)&1 >= (x>>(i-1))&1 {
			i++
		}
		// x[i] = 0
		j := 0
		for (x>>j)&1 == 0 {
			j++
		}
		// x[j] = 1
		x ^= (1 << i)
		x ^= (1 << j)

		// reverse the suffix from 0 to i - 1
		for l, r := i-1, 0; l > r; l, r = l-1, r+1 {
			w := ((x >> l) & 1) ^ ((x >> r) & 1)
			if w == 1 {
				x ^= (1 << l)
				x ^= (1 << r)
			}
		}

		return x
	}

	var arr []int
	T := (1<<k - 1) << (14 - k)

	for w := 1<<k - 1; w <= T; w = next(w) {
		arr = append(arr, w)
	}

	for v := range mx + 1 {
		for _, w := range arr {
			y := v ^ w
			if y <= mx {
				ans += freq[y] * freq[v]
			}
		}
	}

	return ans / 2
}
