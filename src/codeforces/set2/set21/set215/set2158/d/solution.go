package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var n, s, t string
		fmt.Fscan(reader, &n, &s, &t)
		ok, res := solve(s, t)
		if !ok {
			fmt.Fprintln(writer, "-1")
			continue
		}
		fmt.Fprintln(writer, len(res))
		for _, cur := range res {
			fmt.Fprintln(writer, cur[0], cur[1])
		}
	}
}

func solve(s string, t string) (bool, [][]int) {
	n := len(s)

	a := convert(s)
	b := convert(t)

	play := func(a []int) [][]int {
		// 把a变成全0
		var res [][]int

		flip := func(l int, r int) {
			res = append(res, []int{l + 1, r + 1})
		}

		p := -1
		for i := 0; i+1 < n; i++ {
			if a[i] == a[i+1] {
				p = i
				break
			}
		}
		if p == -1 {
			// 0101, or 1010
			if a[0] == 0 {
				// 0101 => 0010
				flip(1, 3)
				for i := 1; i <= 3; i++ {
					a[i] ^= 1
				}
				p = 0
			} else {
				// 1010 => 0100
				flip(0, 2)
				for i := 0; i <= 2; i++ {
					a[i] ^= 1
				}
				p = 2
			}
		} else if a[p] == 1 {
			flip(p, p+1)
		}

		l, r := p, p+1
		d := 0
		for r+1 < n {
			if d != a[r+1] {
				flip(l, r)
				d ^= 1
			}
			r++
		}
		for l > 0 {
			if d != a[l-1] {
				flip(l, r)
				d ^= 1
			}
			l--
		}
		if d != 0 {
			flip(0, n-1)
		}
		return res
	}

	res1 := play(a)
	res2 := play(b)
	slices.Reverse(res2)
	return true, append(res1, res2...)
}

func convert(s string) []int {
	n := len(s)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = int(s[i] - '0')
	}
	return a
}
