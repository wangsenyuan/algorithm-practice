package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)

	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}

	return solve(p)
}

func solve(p []int) int {
	n := len(p)
	t0 := make(BIT, n+3)
	t1 := make(BIT, n+3)
	// 需要知道前面有多少个比v小
	t2 := make(BIT, n+3)

	// 0 for up, 1 for down
	var res int
	for _, v := range p {
		res = add(res, t0.query(v, n), t1.query(v, n))
		// 如果要让v是一个新的half peak
		tmp0 := add(t2.query(0, v), t0.query(0, v), t1.query(0, v))
		// 如果要让v是一个新的half valley
		tmp1 := add(t0.query(v, n), t1.query(v, n))

		t0.add(v, tmp0)
		t1.add(v, tmp1)

		t2.add(v, 1)
	}

	return res
}

type BIT []int

func (bit BIT) add(i int, v int) {
	for i < len(bit) {
		bit[i] = add(bit[i], v)
		i += i & -i
	}
}

func (bit BIT) get(i int) int {
	var res int
	for i > 0 {
		res = add(res, bit[i])
		i -= i & -i
	}
	return res
}

func (bit BIT) query(l int, r int) int {
	return add(bit.get(r), MOD-bit.get(l-1))
}

const MOD = 998244353

func add(nums ...int) int {
	var res int
	for _, v := range nums {
		res += v
		if res >= MOD {
			res -= MOD
		}
	}
	return res
}

func bruteForce(p []int) int {
	n := len(p)
	dp := make([][][]int, n)

	// n 是offset
	var res int
	for i := range n {
		dp[i] = make([][]int, 2)
		for j := range 2 {
			dp[i][j] = make([]int, 3)
		}
		for i1 := range i {
			// 0 for p[?] > p[i1]
			d1 := 0
			if p[i1] < p[i] {
				d1 = 1
			}
			// diff can only be -1, 0, 1
			for d := range 2 {
				for diff := range 3 {
					if dp[i1][d][diff] > 0 {
						diff1 := diff
						if d1 != d {
							if d == 0 {
								diff1--
							} else {
								diff1++
							}
						}
						if diff1 >= 0 && diff1 < 3 {
							dp[i][d1][diff1] = add(dp[i][d1][diff1], dp[i1][d][diff])
						}
					}
				}
			}
		}

		for d := range 2 {
			res = add(res, dp[i][d][2])
		}

		for i1 := range i {
			if p[i1] < p[i] {
				dp[i][1][1] = add(dp[i][1][1], 1)
			}
		}
	}

	return res
}
