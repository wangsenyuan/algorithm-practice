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

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) string {
	freq := make(map[int]int)

	for _, v := range a {
		freq[v]++
	}
	n := len(a)

	swapAndCheck := func(l int, r int) bool {
		a[l], a[r] = a[r], a[l]

		ok := true

		for i := 0; i < n; {
			j := i
			for i < n && a[i] == a[j] {
				i++
			}
			if i-j != freq[a[j]] {
				ok = false
				break
			}
		}

		a[l], a[r] = a[r], a[l]
		return ok
	}

	play := func(num int) string {
		// 需要把num连接到一起
		L, R := -1, -1
		p, q := -1, -1
		for i := range n {
			if a[i] == num {
				if L == -1 {
					L = i
				}
				R = i
			}
		}
		for i := L; i <= R; i++ {
			if a[i] != num {
				if p == -1 {
					p = i
				}
				q = i
			}
		}
		if L > 0 && swapAndCheck(L-1, R) {
			return "YES"
		}
		if R < n-1 && swapAndCheck(L, R+1) {
			return "YES"
		}
		if swapAndCheck(p, R) {
			return "YES"
		}
		if swapAndCheck(L, q) {
			return "YES"
		}
		return "NO"
	}

	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		if i-j == freq[a[j]] {
			continue
		}
		// i - j < freq[a[j]]
		return play(a[j])
	}

	return "YES"
}

func last(a []int) int {
	return a[len(a)-1]
}
