package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
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
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}

	return solve(a, b)
}

const inf = 1 << 60

func solve(a []int, b []int) int {
	n := len(a)
	dp := make([]int, n)
	for i := range n {
		dp[i] = inf
	}
	dp[0] = 0
	var q []vec

	q = append(q, vec{b[0], dp[0]})

	for i := 1; i < n; i++ {
		// dp[i] = min (b[j] * a[i] + dp[j] * 1)
		p := vec{a[i], 1}
		for len(q) > 1 && q[0].dot(p) >= q[1].dot(p) {
			q = q[1:]
		}
		dp[i] = q[0].dot(p)

		s := vec{b[i], dp[i]}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).detCmp(s.sub(q[len(q)-1])) >= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, s)
	}

	return dp[n-1]
}

type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }
func (a vec) detCmp(b vec) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}
