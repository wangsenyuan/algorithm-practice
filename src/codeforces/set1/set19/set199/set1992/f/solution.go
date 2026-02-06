package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	drive(reader, writer)
}

func drive(in *bufio.Reader, out *bufio.Writer) {
	buf := make([]byte, 4096)
	var _i, _n int
	rc := func() byte {
		if _i == _n {
			_n, _ = in.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	rd := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}
	tc := rd()
	for range tc {
		n := rd()
		x := rd()
		a := make([]int, n)
		for i := range n {
			a[i] = rd()
		}
		res := solve(a, x)
		fmt.Fprintln(out, res)
	}
}

func solve(a []int, x int) int {

	var res int
	divs := make(map[int]bool)
	divs[1] = true

	for _, num := range a {
		if num == 1 || x%num != 0 {
			continue
		}

		if _, found := divs[x/num]; found {
			res++
			clear(divs)
			divs[num] = true
			divs[1] = true
			continue
		}

		var arr []int
		for k := range divs {
			if _, ok := divs[k*num]; ok || x%(k*num) != 0 {
				continue
			}
			arr = append(arr, k*num)
		}

		for _, v := range arr {
			divs[v] = true
		}
	}
	return res + 1
}
