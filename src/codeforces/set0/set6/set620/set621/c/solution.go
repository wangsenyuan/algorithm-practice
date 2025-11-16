package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, p int
	fmt.Fscan(reader, &n, &p)
	l := make([]int, n)
	r := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &l[i], &r[i])
	}
	return solve(p, l, r)
}

func solve(p int, l []int, r []int) float64 {
	n := len(l)
	// 换个角度，每次选对的贡献, 给前面1000， 给自己1000， 给后面1000

	var res float64

	get := func(i int) float64 {
		u := r[i] - l[i] + 1
		v := r[i]/p - (l[i]-1)/p
		v = u - v
		return float64(v) / float64(u)
	}

	for i := range n {
		x := get(i)
		y := get((i + 1) % n)
		res += 2000 * (1.0 - x*y)
	}
	return res
}
