package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	fmt.Println(solve(s))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) int {
	n := len(s)
	n++
	p := make([]int, n)
	np := make([]int, n)
	a := make([]int, n)
	na := make([]int, n)

	for i := range n {
		p[i] = i
		if i < n-1 {
			a[i] = int(s[i]-'a') + 1
		}
	}

	k := 27

	cnt := make([]int, max(k, n))

	sort := func(step int) {
		for i := range n {
			np[i] = p[i] - step
			if np[i] < 0 {
				np[i] += n
			}
		}
		clear(cnt)
		for i := range n {
			cnt[a[np[i]]]++
		}
		for i := 1; i < k; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			cnt[a[np[i]]]--
			p[cnt[a[np[i]]]] = np[i]
		}
		for i := range n {
			np[i] = p[i] + step
			if np[i] >= n {
				np[i] -= n
			}
		}

		na[p[0]] = 0
		nk := 1
		for i := 1; i < n; i++ {
			if a[p[i]] != a[p[i-1]] || a[np[i]] != a[np[i-1]] {
				nk++
			}
			na[p[i]] = nk - 1
		}
		copy(a, na)
		k = nk
	}

	sort(0)

	aa := make([][]int, 30)
	var d int
	for ; 1<<d < n; d++ {
		step := 1 << d
		aa[d] = make([]int, n)
		copy(aa[d], a)
		sort(step)
	}
	aa[d] = a
	d++

	lcp := func(x int, y int) int {
		var res int
		for i := d - 1; i >= 0; i-- {
			if aa[i][x] == aa[i][y] {
				res += 1 << i
				x += 1 << i
				y += 1 << i
			}
		}
		return res
	}

	h := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		h[i] = lcp(p[i], p[i+1])
	}

	L := make([]int, n-1)
	stack := make([]int, n)
	var top int
	for i := 0; i < n-1; i++ {
		for top > 0 && h[stack[top-1]] >= h[i] {
			top--
		}
		if top == 0 {
			L[i] = -1
		} else {
			L[i] = stack[top-1]
		}
		stack[top] = i
		top++
	}
	R := make([]int, n-1)
	top = 0
	for i := n - 2; i >= 0; i-- {
		for top > 0 && h[stack[top-1]] > h[i] {
			top--
		}
		if top == 0 {
			R[i] = n - 1
		} else {
			R[i] = stack[top-1]
		}
		stack[top] = i
		top++
	}

	ans := n * (n - 1) / 2
	for i := 1; i < n-1; i++ {
		ans += (i - L[i]) * (R[i] - i) * h[i]
	}

	return ans
}
