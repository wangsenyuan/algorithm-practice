package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, 1<<n)
	for i := range 1 << n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(arr []int) []int {
	n := len(arr)

	res := make([]int, n)

	var f func(l int, r int, l1 int, r1 int)
	f = func(l int, r int, l1 int, r1 int) {
		// 处理区间l...r
		if l+1 == r {
			res[l1] = arr[l]
			return
		}
		// 找这个区间内的最小值
		mn := l
		for i := l; i < r; i++ {
			if arr[i] < arr[mn] {
				mn = i
			}
		}

		mid := (l + r) >> 1
		if mn < mid {
			f(l, mid, l1, (l1+r1)>>1)
			f(mid, r, (l1+r1)>>1, r1)
		} else {
			f(mid, r, l1, (l1+r1)>>1)
			f(l, mid, (l1+r1)>>1, r1)
		}
	}

	f(0, n, 0, n)

	return res
}
