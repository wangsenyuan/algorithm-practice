package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cnt, res := drive(reader)
	fmt.Println(cnt)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (int, []int) {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) (int, []int) {
	n := len(a)
	var res int

	b := slices.Clone(a)

	for i := 1; i+1 < n; i++ {
		if a[i] == a[i-1] || a[i] == a[i+1] {
			continue
		}
		// a[i] != a[i-1] and a[i] != a[i+1]
		// a[i] 需要翻转
		// 0 1 0 1 1
		var j int
		for ; i+j < n; j++ {
			if j&1 == 0 && a[i+j] != a[i] || j&1 == 1 && a[i+j] != a[i+1] {
				break
			}
		}

		res = max(res, j/2)
		j--

		for k := i; k < i+j; k++ {
			if k-(i-1) < i+j-k {
				b[k] = a[i-1]
			} else {
				b[k] = a[i+j]
			}
		}

		i += j - 1
	}

	return res, b
}
