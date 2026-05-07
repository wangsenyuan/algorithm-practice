package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([]int, q)
	for i := range queries {
		fmt.Fscan(reader, &queries[i])
	}
	return solve(a, queries)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, queries []int) []int {
	n := len(a)
	a2 := make([]int, n)

	sum := or(a)
	bitsCount := bits.OnesCount(uint(sum))

	var arr []pair
	arr = append(arr, pair{0, bitsCount})

	for d := range 31 {
		if (sum>>d)&1 == 1 {
			continue
		}
		copy(a2, a)
		var tot int
		for d1 := d; d1 >= 0; d1-- {
			if or(a2)&(1<<d1) == 0 {
				var pos int
				w := 1 << d1
				for i, v := range a2 {
					if v&(w-1) > a2[pos]&(w-1) {
						pos = i
					}
				}
				ops := w - a2[pos]&(w-1)
				// 这里和变成0是一致的
				a2[pos] += ops
				tot += ops
			}
		}
		bitsCount++
		arr = append(arr, pair{tot, bitsCount})
	}

	ans := make([]int, len(queries))

	for i, q := range queries {
		j := sort.Search(len(arr), func(j int) bool {
			return arr[j].first > q
		})
		ans[i] = arr[j-1].second
	}
	return ans
}

func or(a []int) int {
	var res int
	for _, v := range a {
		res |= v
	}
	return res
}
