package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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
	return solve(k, a)
}

func solve(k int, a []int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	if sum < k {
		return -1
	}
	// sum >= k 肯定有答案

	mx := slices.Max(a)

	freq := make([]int, mx+2)
	for _, v := range a {
		freq[v]++
	}

	f2 := make([]int, mx+2)

	check := func(x int) bool {
		if x <= 1 {
			return false
		}

		var cnt int

		copy(f2, freq)

		for i := mx; i >= x; i-- {
			if i/2 >= x {
				if i%2 == 0 {
					f2[i/2] += 2 * f2[i]
				} else {
					f2[i/2] += f2[i]
					f2[i/2+1] += f2[i]
				}
			} else {
				cnt += f2[i]
			}
		}

		return cnt < k
	}

	w := sort.Search(mx+1, check)

	return w - 1
}
