package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, ans := process(reader)
	if len(ans) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
}

func process(reader *bufio.Reader) (p []int, q []int, ans []int) {
	var n int
	fmt.Fscan(reader, &n)
	p = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}

	var k int
	fmt.Fscan(reader, &k)
	q = make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &q[i])
	}

	ans = solve(slices.Clone(p), q)
	return
}

func solve(p []int, q []int) []int {
	n := len(p)
	if n%2 == 1 {
		return nil
	}
	neg := make([]bool, n)
	for _, i := range q {
		neg[i-1] = true
	}

	s1 := make([]int, n)
	var t1 int
	s2 := make([]int, n)
	var t2 int
	ans := slices.Clone(p)

	doIt := func(x int) {
		for t1 > 0 && (t2 > 0 || p[s1[t1-1]] != x) {
			x := p[s1[t1-1]]
			if t2 > 0 && x == p[s2[t2-1]] {
				// ans[s1[t1-1]] = 1
				ans[s2[t2-1]] *= -1
				t2--
			} else {
				s2[t2] = s1[t1-1]
				t2++
			}

			t1--
		}
	}

	for i := 0; i < n; i++ {
		if neg[i] {
			doIt(p[i])
			if t1 == 0 || t2 > 0 {
				// no answer, 中间的也必须是correct的
				return nil
			}
			// p[s1[t1-1]] = p[i]
			// ans[s1[t1-1]] = 1
			t1--
			ans[i] *= -1
		} else {
			s1[t1] = i
			t1++
		}
	}

	if t1 > 0 {
		doIt(-1)
		if t1 > 0 || t2 > 0 {
			return nil
		}
	}

	return ans
}
