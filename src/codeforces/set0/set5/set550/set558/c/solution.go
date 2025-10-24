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

const inf = 1 << 60

func solve(a []int) int {

	op := func(v int, x int) int {
		if x == 0 {
			return bits.Len(uint(v))
		}
		// 如何操作v得到x？
		r1 := bits.TrailingZeros(uint(x))
		x1 := x >> r1
		var r2 int
		for v1 := v; v1 != x1; v1 >>= 1 {
			if v1 < x1 {
				return -1
			}
			r2++
		}

		var res int
		if r2 > 0 {
			res = bits.Len(uint(v & ((1 << r2) - 1)))
			// 先把右边变成0，
			v >>= res
			// 再移动r2次，就可以和x对其了
			r2 -= res
		}
		res += abs(r1 - r2)
		return res
	}

	play := func(x int) int {
		var res int
		for _, v := range a {
			cur := op(v, x)
			if cur < 0 {
				return inf
			}
			res += cur
		}
		return res
	}

	best := play(0)

	first := slices.Min(a)
	second := slices.Max(a)

	for r := range 30 {
		if first>>r == 0 {
			break
		}
		for l := range 30 {
			tmp := first >> r << l
			if tmp > second {
				// 超过最大值的2倍的时候
				break
			}
			best = min(best, play(tmp))
		}
	}

	return best
}

func abs(num int) int {
	return max(num, -num)
}
