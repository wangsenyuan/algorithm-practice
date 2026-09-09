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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	if n <= 2 {
		// 只有[1, 1], [2, 2]
		return n
	}
	// n >= 3
	play := func(start int) int {
		f := make([]int, n+1)
		// f[i] = f[i-2] + a[i] - a[i+1]
		f[start] = a[start] - a[start+1]
		var arr []int
		arr = append(arr, 0, f[start])
		for i := start + 2; i < n; i++ {
			if i+1 < n {
				f[i] = f[i-2] + a[i] - a[i+1]
				arr = append(arr, f[i])
			}
			arr = append(arr, f[i-2]+a[i])
		}
		slices.Sort(arr)
		arr = slices.Compact(arr)
		cnt := make(BIT, len(arr)+2)
		cnt.add(sort.SearchInts(arr, 0), 1)

		var res int
		for i := start; i < n; i += 2 {
			// f[i-2] + a[i] > f[i]
			if i == start {
				res++
			} else {
				w := f[i-2] + a[i]
				j := sort.SearchInts(arr, w)
				if j > 0 {
					res += cnt.get(j - 1)
				}
			}
			cnt.add(sort.SearchInts(arr, f[i]), 1)
		}

		return res
	}

	ans := play(0) + play(1)

	return ans
}

type BIT []int

func (bit BIT) add(p int, v int) {
	p++
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) get(p int) int {
	p++
	var res int
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}
