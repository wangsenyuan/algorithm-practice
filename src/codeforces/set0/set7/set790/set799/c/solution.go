package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, c, d int
	fmt.Fscan(reader, &n, &c, &d)
	fountains := make([][]int, n)
	for i := range n {
		var b, p int
		var t string
		fmt.Fscan(reader, &b, &p, &t)
		t = strings.TrimSpace(t)
		if t == "C" {
			fountains[i] = []int{b, p, 0}
		} else {
			fountains[i] = []int{b, p, 1}
		}
	}
	return solve(c, d, fountains)
}

func solve(c int, d int, fountains [][]int) int {
	var coins [][]int
	var diamonds [][]int
	for _, cur := range fountains {
		if cur[2] == 0 {
			coins = append(coins, cur)
		} else {
			diamonds = append(diamonds, cur)
		}
	}

	best := max(0, solve1(c, coins), solve1(d, diamonds))

	pick := []int{-1, -1}

	for _, cur := range fountains {
		if cur[2] == 0 {
			if cur[1] <= c {
				pick[0] = max(pick[0], cur[0])
			}
		} else {
			if cur[1] <= d {
				pick[1] = max(pick[1], cur[0])
			}
		}
	}

	if pick[0] >= 0 && pick[1] >= 0 {
		best = max(best, pick[0]+pick[1])
	}

	return best
}

const inf = 1 << 60

func solve1(money int, arr [][]int) int {
	if money == 0 {
		return -inf
	}
	// 选择两个，使的它们的price sum <= money, 最大化 beauty sum
	// 对于当前的item， 价格知道的情况下，就是要找到一个范围内的最大值
	s := NewSegTree(money + 1)

	best := -inf

	for _, cur := range arr {
		b, p := cur[0], cur[1]
		if p >= money {
			continue
		}
		// p <= money
		tmp := s.Get(0, money-p+1)
		if tmp > 0 {
			best = max(best, tmp+b)
		}
		tmp = s.Get(p, p+1)
		if b > tmp {
			s.Update(p, b)
		}
	}

	return best
}

type SegTree []int

func NewSegTree(n int) SegTree {
	return make(SegTree, 2*n)
}

func (s SegTree) Update(i int, v int) {
	n := len(s) >> 1
	i += n
	s[i] = v
	for i > 1 {
		s[i>>1] = max(s[i], s[i^1])
		i >>= 1
	}
}

func (s SegTree) Get(l int, r int) int {
	n := len(s) >> 1
	l += n
	r += n
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, s[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, s[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
