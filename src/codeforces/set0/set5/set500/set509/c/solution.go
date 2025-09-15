package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}

	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) (b []int, res []string) {
	var n int
	fmt.Fscan(reader, &n)
	b = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	a := solve(b)
	res = make([]string, n)
	for i, cur := range a {
		var buf []byte
		for j := len(cur) - 1; j >= 0; j-- {
			buf = append(buf, byte(cur[j]+'0'))
		}
		res[i] = string(buf)
	}
	return
}

func solve(b []int) [][]int {
	n := len(b)

	calc2 := func(ds int, w int) []int {
		res := make([]int, w)
		for i := w - 1; i >= 0; i-- {
			var s int
			if i == w-1 {
				s++
			}
			for x := s; x <= 9; x++ {
				if x+i*9 >= ds {
					res[i] = x
					ds -= x
					break
				}
			}
		}
		return res
	}

	calc1 := func(prev []int, ds int, w int) []int {
		// w = len(prev), 尽量和prev一致，且满足后面的相加 == ds
		res := make([]int, w)

		var sum int
		for _, v := range prev {
			sum += v
		}

		for i := 0; i < w; i++ {
			// 如果在这里比 prev[i]大，能够满足条件
			x := prev[i]
			sum -= x
			for y := x + 1; y <= 9; y++ {
				diff := ds - (sum + y)
				if diff >= 0 && diff <= 9*i {
					// 后面的尽可能大，前面的才能尽可能小
					for j := 0; j < i && diff > 0; j++ {
						res[j] = min(diff, 9)
						diff -= res[j]
					}
					res[i] = y

					for j := i + 1; j < w; j++ {
						res[j] = prev[j]
					}
					return res
				}
			}
		}

		return nil
	}

	calc := func(prev []int, ds int) []int {
		w0 := len(prev)
		w1 := max((ds+8)/9, w0)
		if w0 == w1 {
			tmp := calc1(prev, ds, w1)
			if len(tmp) > 0 {
				return tmp
			}
			w1++
		}
		return calc2(ds, w1)
	}
	a := make([][]int, n)

	a[0] = calc2(b[0], (b[0]+8)/9)
	for i := 1; i < n; i++ {
		a[i] = calc(a[i-1], b[i])
	}

	return a
}
