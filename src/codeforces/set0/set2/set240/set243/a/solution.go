package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
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
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	x := slices.Max(a)
	freq := make([]int, 2*x+1)
	H := bits.Len(uint(2 * x))

	pos := make([]int, H)
	for i := range H {
		pos[i] = -2
	}
	n := len(a)

	ptr := make([][]int, n)

	for i := range n {
		ptr[i] = make([]int, H)
		for d := range H {
			ptr[i][d] = pos[d]
			if (a[i]>>d)&1 == 1 {
				pos[d] = i
			}
		}
		or := a[i]
		freq[or]++
		l := i
		for l >= 0 {
			l1 := -1
			for d := range H {
				if (or>>d)&1 == 0 {
					l1 = max(l1, ptr[l][d])
				}
			}
			for d := range H {
				if ptr[i][d] >= l1 {
					or |= 1 << d
				}
			}
			freq[or]++
			l = l1
		}
	}
	var res int
	for i := range 2*x + 1 {
		if freq[i] > 0 {
			res++
		}
	}
	return res
}
