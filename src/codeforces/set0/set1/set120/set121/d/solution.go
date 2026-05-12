package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	segs := make([][]int, n)
	for i := range n {
		segs[i] = make([]int, 2)
		fmt.Fscan(reader, &segs[i][0], &segs[i][1])
	}
	return solve(n, k, segs)
}

const inf = 1 << 60

type uint128 struct {
	hi uint64
	lo uint64
}

func (a uint128) add(b uint128) uint128 {
	lo, carry := bits.Add64(a.lo, b.lo, 0)
	hi, _ := bits.Add64(a.hi, b.hi, carry)
	return uint128{hi, lo}
}

func (a uint128) sub(b uint128) uint128 {
	lo, borrow := bits.Sub64(a.lo, b.lo, 0)
	hi, _ := bits.Sub64(a.hi, b.hi, borrow)
	return uint128{hi, lo}
}

func fromInt(x int) uint128 {
	return uint128{0, uint64(x)}
}

func mulInt(a int, b int) uint128 {
	hi, lo := bits.Mul64(uint64(a), uint64(b))
	return uint128{hi, lo}
}

func (a uint128) leInt(b int) bool {
	return a.hi == 0 && a.lo <= uint64(b)
}

type stat struct {
	sum uint128
	cnt int
}

func solve(n int, k int, segs [][]int) int {
	var limit int
	wn := inf
	for _, cur := range segs {
		limit = max(limit, cur[1])
		wn = min(wn, cur[1]-cur[0]+1)
	}
	limit += k

	var lucky []int

	var gen func(x int)
	gen = func(x int) {
		if x > 0 {
			lucky = append(lucky, x)
		}
		if limit >= 4 && x <= (limit-4)/10 {
			gen(x*10 + 4)
		}
		if limit >= 7 && x <= (limit-7)/10 {
			gen(x*10 + 7)
		}
	}

	gen(0)

	slices.Sort(lucky)

	slices.SortFunc(segs, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	m := len(lucky)
	// dp[i]表示在luck[i]的后面有多少x[?] >= lucky[i]
	dp := make([]stat, m+1)
	for i, j := m-1, n-1; i >= 0; i-- {
		dp[i] = dp[i+1]
		for j >= 0 && segs[j][0] >= lucky[i] {
			dp[i].sum = dp[i].sum.add(fromInt(segs[j][0]))
			dp[i].cnt++
			j--
		}
	}

	slices.SortFunc(segs, func(a, b []int) int {
		return cmp.Or(a[1]-b[1], a[0]-b[0])
	})

	fp := make([]stat, m+1)

	for i, j := 0, 0; i < m; i++ {
		fp[i+1] = fp[i]
		for j < n && segs[j][1] <= lucky[i] {
			fp[i+1].sum = fp[i+1].sum.add(fromInt(segs[j][1]))
			fp[i+1].cnt++
			j++
		}
	}

	check := func(w int) bool {
		if w == 0 {
			return true
		}

		for i := 0; i+w <= m; i++ {
			j := i + w - 1
			if lucky[j]-lucky[i]+1 <= wn {
				// 但是有一个比较麻烦，就是那些刚好落在 [luck[i], luck[j]]两端的，这些是不用移动的
				s1 := dp[i].sum.sub(mulInt(dp[i].cnt, lucky[i]))
				s2 := mulInt(fp[j+1].cnt, lucky[j]).sub(fp[j+1].sum)
				if s1.add(s2).leInt(k) {
					return true
				}
			}
		}
		return false
	}

	l, r := 0, m+1
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}
