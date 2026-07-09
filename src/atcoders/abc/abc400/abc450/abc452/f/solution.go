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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(k, p)
}

func solve(k int, p []int) int {
	n := len(p)
	s1 := make(BIT, n+3)
	s2 := make(BIT, n+3)
	var sum1 int
	var sum2 int

	var ans int
	var l1, l2 int
	for r, v := range p {
		sum1 += (r - l1) - s1.Get(v)
		s1.Update(v, 1)
		for l1 <= r && sum1 >= k {
			s1.Update(p[l1], -1)
			sum1 -= s1.Get(p[l1])
			l1++
		}
		// sum1 < k
		sum2 += (r - l2) - s2.Get(v)
		s2.Update(v, 1)
		// sum1 <= k
		for l2 <= r && sum2 > k {
			s2.Update(p[l2], -1)
			sum2 -= s2.Get(p[l2])
			l2++
		}
		// sum2 <= k
		if sum2 == k {
			ans += l1 - l2
		}
	}

	return ans
}

type BIT []int

func (t BIT) Update(i int, v int) {
	for i < len(t) {
		t[i] += v
		i += i & -i
	}
}

func (t BIT) Get(i int) int {
	var res int
	for i > 0 {
		res += t[i]
		i -= i & -i
	}
	return res
}
