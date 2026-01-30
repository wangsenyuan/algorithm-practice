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
		var l, r int
		fmt.Fscan(reader, &l, &r)
		res := solve(l, r)
		fmt.Fprintln(writer, res)
	}
}

func getSquareRoot(x int) int {
	if x <= 1 {
		return x
	}

	r := sort.Search(x, func(i int) bool {
		return i > 0 && (i > x/i || i*i > x)
	})
	// r * r > x
	return r - 1
}

var cand []int

func init() {
	var arr []int
	for i := 2; i <= 1e6; i++ {
		for j := i * i * i; j <= 1e18; j *= i {
			arr = append(arr, j)
			if j > 1e18/i {
				break
			}
		}
	}
	slices.Sort(arr)
	arr = slices.Compact(arr)

	for _, u := range arr {
		x := getSquareRoot(u)
		if x*x != u {
			cand = append(cand, u)
		}
	}

}

func solve(l int, r int) int {
	if r < 4 {
		if l == 1 {
			return 1
		}
		return 0
	}

	res := getSquareRoot(r) - getSquareRoot(l-1)

	i := sort.SearchInts(cand, r+1)
	j := sort.SearchInts(cand, l)

	return res + i - j
}
